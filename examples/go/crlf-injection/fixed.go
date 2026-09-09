// fixed.go - Correction CWE-93
// Rejet des caracteres de controle (CR/LF) et validation que la cible
// est une URL relative autorisee (allowlist de prefixes).
package handlers

import (
	"net/http"
	"strings"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")

	// FIXED: rejet des caracteres CR/LF et validation de la cible
	if strings.ContainsAny(target, "\r\n") || !strings.HasPrefix(target, "/") {
		http.Error(w, "invalid target", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, target, http.StatusFound)
}
