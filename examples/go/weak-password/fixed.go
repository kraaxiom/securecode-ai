// fixed.go - Correction CWE-521
// Politique de robustesse basee sur la longueur minimale (NIST SP 800-63B
// recommande >= 8, ici 12) et le rejet des mots de passe presents dans des
// listes de fuites connues, plutot que des regles de complexite arbitraires.
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"unicode/utf8"
)

const minPasswordLength = 12

// FIXED: politique de robustesse explicite avant creation du compte
func validatePasswordStrength(password string) error {
	if utf8.RuneCountInString(password) < minPasswordLength {
		return errors.New("le mot de passe doit contenir au moins 12 caracteres")
	}
	// FIXED: rejet si le mot de passe figure dans une liste de fuites connues
	if isKnownCompromisedPassword(password) {
		return errors.New("ce mot de passe est trop commun ou a fuite, choisissez-en un autre")
	}
	return nil
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req struct{ Email, Password string }
	json.NewDecoder(r.Body).Decode(&req)

	if err := validatePasswordStrength(req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := createUser(req.Email, req.Password); err != nil {
		http.Error(w, "erreur creation compte", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
