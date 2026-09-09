// vulnerable.go - CWE-1333: Denial of Service par expression reguliere
// Le moteur regexp de Go est base sur RE2 (complexite lineaire garantie,
// pas de backtracking catastrophique), mais l'absence de limite de
// taille sur l'entree reste un vecteur de deni de service : une chaine
// de plusieurs Mo appliquee a une regex complexe consomme un temps CPU
// et une memoire proportionnels a sa taille, sans aucune borne.
package validate

import (
	"net/http"
	"regexp"
)

var emailPattern = regexp.MustCompile(`^([a-zA-Z0-9._%-]+)+@([a-zA-Z0-9.-]+)+\.[a-zA-Z]{2,}$`)

func ValidateEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	// VULNERABLE: aucune limite de taille avant application de la regex
	if !emailPattern.MatchString(email) {
		http.Error(w, "email invalide", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
