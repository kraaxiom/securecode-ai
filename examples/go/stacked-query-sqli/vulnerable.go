// vulnerable.go - CWE-89: Stacked Query SQL Injection
// Utilisation de db.Exec avec concatenation permettant a un attaquant
// d'empiler des requetes supplementaires (ex: DROP, UPDATE) separees
// par un point-virgule.
package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func UpdateNicknameHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	nickname := r.URL.Query().Get("nickname")

	// VULNERABLE: concatenation autorisant l'empilement de requetes
	query := fmt.Sprintf("UPDATE users SET nickname = '%s' WHERE id = %s", nickname, userID)
	if _, err := db.Exec(query); err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("updated"))
}
