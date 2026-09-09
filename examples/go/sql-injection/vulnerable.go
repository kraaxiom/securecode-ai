// vulnerable.go - CWE-89: SQL Injection
// La requete est construite par concatenation de l'entree utilisateur,
// permettant a un attaquant d'injecter du SQL arbitraire.
package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func GetUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	// VULNERABLE: concatenation directe de l'entree utilisateur dans la requete SQL
	query := fmt.Sprintf("SELECT id, email FROM users WHERE username = '%s'", username)
	row := db.QueryRow(query)

	var id int
	var email string
	if err := row.Scan(&id, &email); err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "id=%d email=%s", id, email)
}
