// vulnerable.go - CWE-307: Improper Restriction of Excessive Authentication Attempts
// Aucune limitation du nombre de tentatives : un attaquant peut tester
// un nombre illimite de mots de passe sans etre bloque.
package handlers

import (
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct{ Email, Password string }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// VULNERABLE: aucune limitation de debit, aucun compteur d'echecs
	user, ok := authenticate(req.Email, req.Password)
	if !ok {
		http.Error(w, "identifiants invalides", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": issueToken(user)})
}
