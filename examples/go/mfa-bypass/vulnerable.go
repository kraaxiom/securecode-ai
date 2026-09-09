// vulnerable.go - CWE-287: Improper Authentication
// L'etape MFA est controlee par un parametre envoye par le client :
// un attaquant qui a deja les identifiants peut simplement omettre ou
// falsifier ce parametre pour contourner la verification du code OTP.
package handlers

import (
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email      string
		Password   string
		MFAVerified bool // VULNERABLE: la confiance est placee dans le client
	}
	json.NewDecoder(r.Body).Decode(&req)

	user, ok := authenticate(req.Email, req.Password)
	if !ok {
		http.Error(w, "identifiants invalides", http.StatusUnauthorized)
		return
	}

	// VULNERABLE: le serveur fait confiance au flag envoye par le client
	// au lieu de verifier lui-meme le code OTP
	if req.MFAVerified {
		json.NewEncoder(w).Encode(map[string]string{"token": issueToken(user)})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "mfa_required"})
}
