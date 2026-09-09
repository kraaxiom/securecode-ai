// vulnerable.go - CWE-307: Improper Restriction of Excessive Authentication Attempts
// La limitation est appliquee uniquement par compte, jamais par IP/source.
// Un attaquant peut donc tester un seul mot de passe frequent sur des
// milliers de comptes distincts depuis la meme source sans etre bloque.
package handlers

import (
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct{ Email, Password string }
	json.NewDecoder(r.Body).Decode(&req)

	// VULNERABLE: limitation uniquement par compte, jamais par IP globale
	if isAccountLocked(req.Email) {
		http.Error(w, "compte verrouille", http.StatusForbidden)
		return
	}

	user, ok := authenticate(req.Email, req.Password)
	if !ok {
		recordFailedAttemptPerAccount(req.Email)
		http.Error(w, "identifiants invalides", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": issueToken(user)})
}
