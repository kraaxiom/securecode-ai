// Package main - validateur de webhook asynchrone.
//
// CORRIGÉ : mitige CWE-918 (SSRF aveugle) en appliquant les mêmes contrôles
// qu'une SSRF classique (whitelist, résolution/validation d'IP, pas de
// redirection automatique) même si le résultat n'est jamais exposé
// directement au client, ainsi qu'une journalisation des tentatives
// refusées pour permettre la détection de scans internes.
//
// Voir vulnerable.go pour la version non protégée.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// webhookDomainAllowlist : hôtes de webhook explicitement autorisés
// (typiquement fournis par un partenaire intégré et validés hors-bande,
// pas saisis librement par n'importe quel utilisateur final).
var webhookDomainAllowlist = map[string]bool{
	"partner-a.example.com": true,
	"partner-b.example.com": true,
}

func isBlockedIP(ip net.IP) bool {
	return ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() ||
		ip.IsPrivate() || ip.IsUnspecified() || ip.IsMulticast()
}

// safeDialContext : résolution DNS unique + validation de l'IP + pinning,
// identique au principe appliqué dans examples/go/ssrf-classique/fixed.go.
func safeDialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	ips, err := net.DefaultResolver.LookupIP(ctx, "ip", host)
	if err != nil {
		return nil, fmt.Errorf("résolution DNS impossible: %w", err)
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

func newSafeHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{DialContext: safeDialContext},
		Timeout:   5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("redirections désactivées")
		},
	}
}

func validateWebhookURL(rawURL string) (*url.URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("URL invalide: %w", err)
	}
	if u.Scheme != "https" {
		return nil, errors.New("seul https est autorisé pour les webhooks")
	}
	host := strings.ToLower(u.Hostname())
	if !webhookDomainAllowlist[host] {
		return nil, errors.New("hôte de webhook non autorisé")
	}
	return u, nil
}

// registerWebhook applique désormais la même politique de sécurité que la
// SSRF classique, même si le résultat de la requête n'est jamais renvoyé
// au client. Les échecs de validation sont journalisés côté serveur pour
// permettre de détecter des tentatives de scan interne répétées.
func registerWebhook(webhookURL string) {
	u, err := validateWebhookURL(webhookURL)
	if err != nil {
		log.Printf("webhook refusé (validation SSRF): %v", err)
		return
	}

	go func() {
		client := newSafeHTTPClient()
		resp, err := client.Head(u.String())
		if err != nil {
			log.Printf("échec de vérification du webhook %s: %v", u.Host, err)
			return
		}
		defer resp.Body.Close()
	}()
}

// registerHandler expose l'enregistrement de webhook via HTTP.
// POST /webhooks/register?url=<url fournie par l'utilisateur>
func registerHandler(w http.ResponseWriter, r *http.Request) {
	webhookURL := r.URL.Query().Get("url")
	if webhookURL == "" {
		http.Error(w, "paramètre 'url' requis", http.StatusBadRequest)
		return
	}

	registerWebhook(webhookURL)

	// Réponse générique, identique que la validation réussisse ou non,
	// pour ne pas transformer ce endpoint en oracle de reconnaissance
	// réseau interne.
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "webhook en cours de vérification")
}

func main() {
	http.HandleFunc("/webhooks/register", registerHandler)
	http.ListenAndServe(":8080", nil)
}
