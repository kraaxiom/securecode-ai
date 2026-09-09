// vulnerable.go - CWE-93: CRLF Injection
// L'entree utilisateur est ecrite telle quelle dans un en-tete HTTP,
// permettant l'injection de sequences CRLF pour ajouter des en-tetes
// ou scinder la reponse.
package handlers

import "net/http"

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")

	// VULNERABLE: entree utilisateur non filtree dans un en-tete
	w.Header().Set("Location", target)
	w.WriteHeader(http.StatusFound)
}
