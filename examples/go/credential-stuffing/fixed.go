// fixed.go - Correction CWE-307
// Detection des patterns de credential stuffing : limitation par IP sur le
// nombre de comptes distincts testes, verification des mots de passe
// compromis, et defi CAPTCHA / MFA en cas d'anomalie.
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

	// FIXED: detection d'un volume anormal de comptes distincts testes
	// depuis la meme source (signature de credential stuffing)
	if isSuspiciousLoginVelocity(r.RemoteAddr) {
		requireCaptcha(w, r)
		return
	}

	// FIXED: refus des mots de passe presents dans des fuites connues
	if isKnownCompromisedPassword(req.Password) {
		http.Error(w, "mot de passe compromis, veuillez le reinitialiser", http.StatusForbidden)
		return
	}

	user, ok := authenticate(req.Email, req.Password)
	if !ok {
		recordFailedAttempt(r.RemoteAddr, req.Email)
		http.Error(w, "identifiants invalides", http.StatusUnauthorized)
		return
	}

	// FIXED: MFA obligatoire si le contexte de connexion est inhabituel
	if requiresStepUpMFA(user, r) {
		json.NewEncoder(w).Encode(map[string]string{"status": "mfa_required"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": issueToken(user)})
}
