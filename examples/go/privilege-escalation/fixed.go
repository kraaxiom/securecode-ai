// fixed.go - Correction CWE-269
// Seul un utilisateur disposant du privilege d'administration des roles
// peut modifier le role d'un compte, et il ne peut jamais s'attribuer un
// role superieur au sien (principe du moindre privilege).
package handlers

import (
	"encoding/json"
	"net/http"
)

var assignableRoles = map[string]bool{"member": true, "viewer": true}

func ChangeUserRoleHandler(w http.ResponseWriter, r *http.Request) {
	currentUser := getCurrentUser(r)
	if currentUser == nil || !currentUser.HasPermission("roles:manage") {
		// FIXED: seul un titulaire explicite du privilege peut modifier des roles
		http.Error(w, "acces refuse", http.StatusForbidden)
		return
	}

	var req struct {
		UserID  string
		NewRole string
	}
	json.NewDecoder(r.Body).Decode(&req)

	// FIXED: le role cible doit appartenir a une liste blanche de roles
	// attribuables ; un administrateur ne peut pas etre cree via cet endpoint
	if !assignableRoles[req.NewRole] {
		http.Error(w, "role non autorise via cet endpoint", http.StatusForbidden)
		return
	}

	// FIXED: interdiction de modifier son propre role (evite l'auto-elevation)
	if req.UserID == currentUser.ID {
		http.Error(w, "impossible de modifier son propre role", http.StatusForbidden)
		return
	}

	if err := setUserRole(req.UserID, req.NewRole); err != nil {
		http.Error(w, "erreur", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
