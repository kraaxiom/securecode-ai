// fixed.go - Correction CWE-307
// Ajout d'une limitation globale par IP/source en complement de la
// limitation par compte, capable de detecter un faible nombre d'echecs
// repartis sur un grand nombre de comptes distincts (signature du
// password spraying).
package handlers

import (
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct{ Email, Password string }
	json.NewDecoder(r.Body).Decode(&req)

	// FIXED: limitation globale par IP, independante du compte cible
	if isIPRateLimited(r.RemoteAddr) {
		http.Error(w, "trop de tentatives depuis cette source", http.StatusTooManyRequests)
		return
	}
	if isAccountLocked(req.Email) {
		http.Error(w, "compte verrouille", http.StatusForbidden)
		return
	}

	user, ok := authenticate(req.Email, req.Password)
	if !ok {
		recordFailedAttemptPerAccount(req.Email)
		// FIXED: le compteur global par IP est incremente meme si le
		// compte cible est different a chaque essai, ce qui detecte le spray
		recordFailedAttemptPerIP(r.RemoteAddr)
		http.Error(w, "identifiants invalides", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": issueToken(user)})
}
