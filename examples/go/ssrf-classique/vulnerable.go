// Package main - service d'aperçu d'URL (URL preview / webhook fetcher).
//
// VULNÉRABLE : CWE-918 (Server-Side Request Forgery).
//
// Ce programme illustre une SSRF classique : l'application accepte une URL
// fournie par l'utilisateur et effectue elle-même la requête HTTP sortante,
// sans aucune validation de la destination (pas de whitelist, pas de
// filtrage des plages d'IP privées/loopback/link-local).
//
// Ceci n'est PAS un outil d'attaque : il n'y a aucun payload, aucun serveur
// d'exploitation. C'est un exemple pédagogique de code applicatif vulnérable
// (typique d'une fonctionnalité "aperçu de lien" ou "webhook fetcher").
package main

import (
	"fmt"
	"io"
	"net/http"
)

// fetchURLPreview récupère le contenu d'une URL fournie par l'utilisateur
// pour en générer un aperçu (ex: prévisualisation d'un lien partagé).
//
// POURQUOI C'EST VULNÉRABLE :
//   - `rawURL` provient directement d'une requête utilisateur (paramètre HTTP).
//   - Aucune whitelist de domaines/hôtes autorisés n'est appliquée.
//   - Aucune résolution DNS préalable ni filtrage des plages d'adresses
//     privées (10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16), loopback
//     (127.0.0.0/8) ou link-local (169.254.0.0/16, qui héberge notamment
//     les services de métadonnées cloud) n'est effectué avant l'appel.
//   - Le client HTTP par défaut suit les redirections automatiquement,
//     ce qui permet de contourner un filtrage même partiel en amont
//     (redirection 302 vers une IP interne après validation d'une URL
//     publique).
//
// Un attaquant peut donc fournir une URL pointant vers un service interne
// (ex: http://127.0.0.1:6379, http://169.254.169.254/...) et le serveur
// effectuera la requête à sa place, exposant potentiellement des données
// internes dans la réponse renvoyée au client (SSRF non aveugle : le
// contenu récupéré est directement affiché).
func fetchURLPreview(rawURL string) (string, error) {
	// Aucune validation de rawURL avant utilisation : le serveur fait
	// confiance à une entrée entièrement contrôlée par l'utilisateur.
	resp, err := http.Get(rawURL)
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

// previewHandler expose la fonctionnalité vulnérable via HTTP.
// GET /preview?url=<url fournie par l'utilisateur>
func previewHandler(w http.ResponseWriter, r *http.Request) {
	targetURL := r.URL.Query().Get("url")
	if targetURL == "" {
		http.Error(w, "paramètre 'url' requis", http.StatusBadRequest)
		return
	}

	content, err := fetchURLPreview(targetURL)
	if err != nil {
		http.Error(w, "erreur lors de la récupération", http.StatusBadGateway)
		return
	}

	// Le contenu récupéré (potentiellement interne) est renvoyé au client.
	fmt.Fprint(w, content)
}

func main() {
	http.HandleFunc("/preview", previewHandler)
	http.ListenAndServe(":8080", nil)
}
