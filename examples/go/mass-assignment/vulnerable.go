// vulnerable.go - CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes
// La structure de requete est deserialisee directement dans le modele
// persiste, y compris des champs sensibles comme "Role" ou "IsAdmin"
// qu'un attaquant peut injecter dans le payload JSON.
package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID      string
	Email   string
	Name    string
	Role    string // champ sensible : ne devrait jamais etre modifiable par l'utilisateur
	IsAdmin bool   // champ sensible
}

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	currentUser := getCurrentUser(r)

	var updated User
	// VULNERABLE: deserialisation directe du JSON attaquant dans le modele,
	// y compris les champs Role et IsAdmin non censes etre exposes
	json.NewDecoder(r.Body).Decode(&updated)
	updated.ID = currentUser.ID

	saveUser(&updated)
	w.WriteHeader(http.StatusOK)
}
