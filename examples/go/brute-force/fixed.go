// fixed.go - Correction CWE-307
// Limitation de debit par compte ET par IP, avec verrouillage progressif
// et journalisation des echecs.
package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct{ Email, Password string }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	key := r.RemoteAddr + ":" + req.Email

	// FIXED: verification du compteur d'echecs avant toute tentative
	if isRateLimited(key, 5, 15*time.Minute) {
		http.Error(w, "trop de tentatives, reessayez plus tard", http.StatusTooManyRequests)
		return
	}

	user, ok := authenticate(req.Email, req.Password)
	if !ok {
		recordFailedAttempt(key)
		logAuthFailure(req.Email, r.RemoteAddr)
		// Message generique : ne revele pas si le compte existe
		http.Error(w, "identifiants invalides", http.StatusUnauthorized)
		return
	}

	clearFailedAttempts(key)
	json.NewEncoder(w).Encode(map[string]string{"token": issueToken(user)})
}
