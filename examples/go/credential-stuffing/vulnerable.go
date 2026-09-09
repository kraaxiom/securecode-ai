// vulnerable.go - CWE-307: Improper Restriction of Excessive Authentication Attempts
// Le login accepte un volume illimite de tentatives depuis n'importe quelle
// source, sans detection des patterns typiques du credential stuffing
// (grand nombre de comptes distincts testes depuis une meme IP/plage).
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

	// VULNERABLE: aucune detection d'anomalie, aucun captcha,
	// pas de verification contre une liste de mots de passe compromis
	user, ok := authenticate(req.Email, req.Password)
	if !ok {
		http.Error(w, "identifiants invalides", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": issueToken(user)})
}
