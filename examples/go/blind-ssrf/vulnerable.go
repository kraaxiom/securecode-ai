// Package main - validateur de webhook asynchrone.
//
// VULNÉRABLE : CWE-918 (Server-Side Request Forgery), variante "blind SSRF".
//
// Ce programme illustre une SSRF aveugle : l'application enregistre une URL
// de webhook fournie par l'utilisateur, puis un worker en arrière-plan
// tente de la "vérifier" en émettant une requête HTTP, sans jamais renvoyer
// le contenu ni le statut détaillé au client. L'attaquant ne voit aucun
// résultat direct, mais peut tout de même déclencher des requêtes internes
// (scan de ports, exfiltration via callback DNS/HTTP hors bande, etc.).
//
// Exemple pédagogique uniquement : aucun payload, aucun serveur
// d'exploitation.
package main

import (
	"fmt"
	"net/http"
	"time"
)

// registerWebhook enregistre une URL de webhook fournie par l'utilisateur
// puis déclenche une vérification asynchrone en arrière-plan.
//
// POURQUOI C'EST VULNÉRABLE :
//   - `webhookURL` vient directement de l'utilisateur, sans whitelist.
//   - La vérification se fait dans une goroutine "fire and forget" : le
//     résultat (succès, erreur, contenu) n'est jamais renvoyé au client,
//     ce qui donne l'illusion trompeuse qu'il n'y a "rien à voir" alors que
//     la requête interne a bien été émise depuis le serveur.
//   - Aucun filtrage des plages d'adresses privées/loopback/link-local
//     n'est appliqué avant l'appel HTTP.
//   - Même sans lire la réponse, l'attaquant peut exploiter le délai de
//     réponse (time-based) ou provoquer un effet de bord côté service
//     interne visé (ex: requête GET déclenchant une action non idempotente).
func registerWebhook(webhookURL string) {
	// Réponse immédiate au client : "on s'en occupe", sans jamais exposer
	// le résultat de la requête sortante.
	go func() {
		client := &http.Client{Timeout: 5 * time.Second}
		// Aucune validation de webhookURL avant l'appel.
		resp, err := client.Head(webhookURL)
		if err == nil {
			resp.Body.Close()
		}
		// Le résultat n'est journalisé nulle part d'accessible côté
		// utilisateur : la requête a eu lieu, mais reste "aveugle".
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

	// Réponse générique immédiate, sans détail sur le résultat de la
	// requête sortante réelle.
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "webhook en cours de vérification")
}

func main() {
	http.HandleFunc("/webhooks/register", registerHandler)
	http.ListenAndServe(":8080", nil)
}
