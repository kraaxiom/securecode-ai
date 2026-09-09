// fixed.go - Correction CWE-89
// Requete parametree : l'entree ne peut plus fermer la chaine
// litterale ni ajouter une clause UNION.
package handlers

import (
	"database/sql"
	"net/http"
)

func SearchArticleHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	// FIXED: requete preparee
	query := "SELECT id, title FROM articles WHERE title = ?"
	rows, err := db.Query(query, title)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	w.Write([]byte("ok"))
}
