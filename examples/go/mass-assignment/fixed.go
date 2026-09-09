// fixed.go - Correction CWE-915
// Utilisation d'un DTO explicite ne contenant que les champs autorises a
// etre modifies par l'utilisateur : les champs sensibles (Role, IsAdmin)
// ne peuvent jamais etre affectes depuis la requete client.
package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID      string
	Email   string
	Name    string
	Role    string
	IsAdmin bool
}

// FIXED: DTO d'entree explicite, liste blanche des champs modifiables
type UpdateProfileRequest struct {
	Email string
	Name  string
}

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	currentUser := getCurrentUser(r)

	var req UpdateProfileRequest
	json.NewDecoder(r.Body).Decode(&req)

	// FIXED: seuls les champs explicitement autorises sont copies sur le
	// modele persiste, Role et IsAdmin restent inaccessibles au client
	currentUser.Email = req.Email
	currentUser.Name = req.Name

	saveUser(currentUser)
	w.WriteHeader(http.StatusOK)
}
