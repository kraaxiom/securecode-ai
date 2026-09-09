// vulnerable.go - CWE-862: Missing Authorization
// Le endpoint d'administration ne verifie que l'authentification (l'utilisateur
// est connecte) mais jamais le role/la fonction requise, ce qui permet a
// n'importe quel utilisateur authentifie d'appeler une fonction reservee aux admins.
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

	// VULNERABLE: aucune verification du role/de la fonction du user courant,
	// seule l'authentification est controlee
	targetID := r.URL.Query().Get("user_id")
	if err := deleteUser(targetID); err != nil {
		http.Error(w, "erreur", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}
