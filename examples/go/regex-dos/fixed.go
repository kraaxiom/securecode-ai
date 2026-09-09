// fixed.go - Correction CWE-1333
// On impose une limite stricte de longueur avant tout traitement par
// regexp (defense en profondeur, meme si RE2 garantit une complexite
// lineaire) et on simplifie le motif pour supprimer les groupes repetes
// imbriques inutiles, qui n'apportent rien a la validation.
package validate

import (
	"net/http"
	"regexp"
)

const maxEmailLength = 254 // limite RFC 5321

var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func ValidateEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	// FIXED: limite de taille appliquee avant tout traitement regex
	if len(email) == 0 || len(email) > maxEmailLength {
		http.Error(w, "email invalide", http.StatusBadRequest)
		return
	}

	// FIXED: motif non ambigu, sans groupes repetes imbriques
	if !emailPattern.MatchString(email) {
		http.Error(w, "email invalide", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
