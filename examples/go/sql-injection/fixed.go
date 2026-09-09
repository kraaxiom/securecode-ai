// fixed.go - Correction CWE-89
// Utilisation d'une requete parametree (placeholder) : le pilote SQL
// separe strictement le code SQL des donnees utilisateur.
package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func GetUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	// FIXED: requete preparee avec parametre lie
	query := "SELECT id, email FROM users WHERE username = ?"
	row := db.QueryRow(query, username)

	var id int
	var email string
	if err := row.Scan(&id, &email); err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "id=%d email=%s", id, email)
}
