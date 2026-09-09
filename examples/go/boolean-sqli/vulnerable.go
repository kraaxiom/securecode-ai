// vulnerable.go - CWE-89: Boolean-based SQL Injection
// La reponse HTTP differe selon la verite de la condition SQL injectee,
// ce qui permet une extraction de donnees bit par bit.
package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func SearchProductHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	// VULNERABLE: concatenation directe dans le LIKE
	query := fmt.Sprintf("SELECT id FROM products WHERE name LIKE '%%%s%%'", name)
	rows, err := db.Query(query)
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
