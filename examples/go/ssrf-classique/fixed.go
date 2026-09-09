// Package main - service d'aperçu d'URL (URL preview / webhook fetcher).
//
// CORRIGÉ : mitige CWE-918 (Server-Side Request Forgery) via une whitelist
// de domaines, une résolution DNS explicite avec validation de l'IP
// résultante contre les plages privées/réservées, et la désactivation des
// redirections automatiques.
//
// Voir vulnerable.go pour la version non protégée et l'explication du
// problème d'origine.
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// domainAllowlist : seuls ces hôtes exacts peuvent être contactés.
// En production, cette liste doit venir d'une configuration métier validée,
// pas d'une entrée utilisateur.
var domainAllowlist = map[string]bool{
	"images.example.com": true,
	"cdn.example.com":    true,
}

// isBlockedIP refuse les plages privées, loopback, link-local et les
// adresses réservées (RFC 1918, RFC 3927, etc.). Le lien-local
// 169.254.0.0/16 est explicitement couvert : il héberge les services de
// métadonnées cloud (AWS/Azure/GCP) qui ne doivent jamais être atteignables
// depuis une requête utilisateur.
func isBlockedIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() ||
		ip.IsPrivate() || ip.IsUnspecified() || ip.IsMulticast() {
		return true
	}
	return false
}

// safeDialContext force la résolution DNS à se faire une seule fois ici
// (pas dans le client HTTP), valide l'IP obtenue, puis se connecte
// directement sur cette IP validée (DNS pinning). Cela empêche le
// contournement par DNS rebinding : sans ce pinning, le client HTTP
// pourrait résoudre le nom une seconde fois et obtenir une IP différente
// (interne) entre la validation et la requête réelle.
func safeDialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	ips, err := net.DefaultResolver.LookupIP(ctx, "ip", host)
	if err != nil {
		return nil, fmt.Errorf("résolution DNS impossible: %w", err)
	}
	if len(ips) == 0 {
		return nil, errors.New("aucune adresse résolue")
	}

	var validIP net.IP
	for _, ip := range ips {
		if !isBlockedIP(ip) {
			validIP = ip
			break
		}
	}
	if validIP == nil {
		return nil, errors.New("destination refusée: adresse IP interne/réservée")
	}

	dialer := &net.Dialer{Timeout: 5 * time.Second}
	return dialer.DialContext(ctx, network, net.JoinHostPort(validIP.String(), port))
}

// newSafeHTTPClient construit un client HTTP dont le Dialer applique la
// validation d'IP décrite ci-dessus, et qui refuse de suivre les
// redirections automatiquement (chaque redirection potentielle doit être
// revalidée explicitement par l'appelant, pas suivie en aveugle).
func newSafeHTTPClient() *http.Client {
	transport := &http.Transport{
		DialContext:           safeDialContext,
		ResponseHeaderTimeout: 5 * time.Second,
	}
	return &http.Client{
		Transport: transport,
		Timeout:   8 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Refuser toute redirection automatique : cela empêche
			// qu'une URL publique validée redirige (302) vers une
			// ressource interne après coup.
			return errors.New("redirections désactivées")
		},
	}
}

// validateTargetURL vérifie le schéma et l'hôte contre la whitelist avant
// toute tentative de connexion.
func validateTargetURL(rawURL string) (*url.URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("URL invalide: %w", err)
	}

	if u.Scheme != "https" && u.Scheme != "http" {
		return nil, errors.New("schéma non autorisé")
	}

	host := strings.ToLower(u.Hostname())
	if !domainAllowlist[host] {
		return nil, errors.New("hôte non autorisé (hors whitelist)")
	}

	return u, nil
}

// fetchURLPreview récupère le contenu d'une URL, uniquement si elle
// appartient à la whitelist de domaines et résout vers une IP publique
// non réservée.
func fetchURLPreview(rawURL string) (string, error) {
	u, err := validateTargetURL(rawURL)
	if err != nil {
		return "", err
	}

	client := newSafeHTTPClient()
	resp, err := client.Get(u.String())
	if err != nil {
		return "", fmt.Errorf("échec de la requête: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20)) // 1 Mo max
	if err != nil {
		return "", fmt.Errorf("échec de lecture: %w", err)
	}

	return string(body), nil
}

// previewHandler expose la fonctionnalité corrigée via HTTP.
// GET /preview?url=<url fournie par l'utilisateur>
func previewHandler(w http.ResponseWriter, r *http.Request) {
	targetURL := r.URL.Query().Get("url")
	if targetURL == "" {
		http.Error(w, "paramètre 'url' requis", http.StatusBadRequest)
		return
	}

	content, err := fetchURLPreview(targetURL)
	if err != nil {
		// Message générique : ne pas divulguer les détails internes
		// de la validation (évite de fournir un oracle à l'attaquant).
		http.Error(w, "impossible de récupérer cette URL", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, content)
}

func main() {
	http.HandleFunc("/preview", previewHandler)
	http.ListenAndServe(":8080", nil)
}
