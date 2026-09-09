// vulnerable.go - CWE-269: Improper Privilege Management
// L'endpoint de changement de role fait confiance au role cible envoye par
// le client sans verifier que l'appelant a le droit d'attribuer ce role,
// permettant a un utilisateur standard de s'auto-promouvoir administrateur.
package handlers

import (
	"encoding/json"
	"net/http"
)

func ChangeUserRoleHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID  string
		NewRole string
	}
	json.NewDecoder(r.Body).Decode(&req)

	// VULNERABLE: aucune verification que l'appelant a le droit d'attribuer
	// ce role, ni que le role cible est autorise (ex: "admin")
	if err := setUserRole(req.UserID, req.NewRole); err != nil {
		http.Error(w, "erreur", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
