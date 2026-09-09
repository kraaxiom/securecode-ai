// fixed.go - Correction CWE-287
// La verification MFA est entierement realisee cote serveur : le code OTP
// soumis est controle contre l'etat de session stocke serveur, jamais
// contre un flag fourni par le client.
package handlers

import (
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct{ Email, Password string }
	json.NewDecoder(r.Body).Decode(&req)

	user, ok := authenticate(req.Email, req.Password)
	if !ok {
		http.Error(w, "identifiants invalides", http.StatusUnauthorized)
		return
	}

	// FIXED: creation d'une session intermediaire "pending_mfa" cote serveur,
	// aucun jeton final n'est emis avant verification du second facteur
	mfaSessionID := createPendingMFASession(user.ID)
	json.NewEncoder(w).Encode(map[string]string{
		"status":         "mfa_required",
		"mfa_session_id": mfaSessionID,
	})
}

func VerifyMFAHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		MFASessionID string
		Code         string
	}
	json.NewDecoder(r.Body).Decode(&req)

	// FIXED: le code OTP est verifie serveur, lie a une session en attente
	// et a un compteur de tentatives limite
	user, ok := verifyPendingMFA(req.MFASessionID, req.Code)
	if !ok {
		http.Error(w, "code invalide ou expire", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": issueToken(user)})
}
