// vulnerable.go - CWE-89: UNION-based SQL Injection
// L'entree utilisateur est concatenee dans une requete SELECT,
// permettant d'y ajouter une clause UNION SELECT pour extraire
// des donnees d'autres tables.
package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func SearchArticleHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	// VULNERABLE: concatenation, colonne count/type non controles
	query := fmt.Sprintf("SELECT id, title FROM articles WHERE title = '%s'", title)
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	w.Write([]byte("ok"))
}
