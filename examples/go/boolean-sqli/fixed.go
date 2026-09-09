// fixed.go - Correction CWE-89
// Le motif LIKE est construit cote application et passe comme parametre
// lie, jamais concatene dans la requete SQL.
package handlers

import (
	"database/sql"
	"net/http"
)

func SearchProductHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	// FIXED: le motif est un parametre, pas du SQL concatene
	pattern := "%" + name + "%"
	query := "SELECT id FROM products WHERE name LIKE ?"
	rows, err := db.Query(query, pattern)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	if rows.Next() {
		w.Write([]byte("found"))
	} else {
		w.Write([]byte("not found"))
	}
}
