// fixed.go - Correction CWE-113
// Utilisation de http.SetCookie qui echappe correctement la valeur,
// combinee a une validation d'entree stricte.
package handlers

import (
	"net/http"
	"regexp"
)

var prefPattern = regexp.MustCompile(`^[a-zA-Z0-9_-]{1,32}$`)

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("pref")
	if !prefPattern.MatchString(value) {
		http.Error(w, "invalid preference", http.StatusBadRequest)
		return
	}

	// FIXED: API standard qui gere correctement l'encodage du cookie
	http.SetCookie(w, &http.Cookie{Name: "pref", Value: value, Path: "/", HttpOnly: true})
	w.Write([]byte("saved"))
}
