// fixed.go - Correction CWE-862
// Verification explicite de l'autorisation fonctionnelle (role/permission)
// requise pour l'action, en plus de l'authentification.
package handlers

import (
	"encoding/json"
	"net/http"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	currentUser := getCurrentUser(r)
	if currentUser == nil {
		http.Error(w, "non authentifie", http.StatusUnauthorized)
		return
	}

	// FIXED: verification explicite de la permission fonctionnelle requise
	if !currentUser.HasPermission("users:delete") {
		http.Error(w, "acces refuse", http.StatusForbidden)
		return
	}

	targetID := r.URL.Query().Get("user_id")
	if err := deleteUser(targetID); err != nil {
		http.Error(w, "erreur", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}
