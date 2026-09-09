// vulnerable.go - CWE-521: Weak Password Requirements
// Aucune exigence minimale de robustesse n'est imposee sur le mot de passe
// choisi par l'utilisateur : longueur minimale absente, aucune verification
// contre les mots de passe communs ou compromis.
package handlers

import (
	"encoding/json"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req struct{ Email, Password string }
	json.NewDecoder(r.Body).Decode(&req)

	// VULNERABLE: aucune politique de robustesse appliquee sur le mot de passe
	if err := createUser(req.Email, req.Password); err != nil {
		http.Error(w, "erreur creation compte", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
