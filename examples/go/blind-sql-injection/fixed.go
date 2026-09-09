// fixed.go - Correction CWE-89
// Requete parametree : l'entree ne peut plus modifier la structure
// logique de la clause WHERE.
package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
)

func CheckAccountHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	// FIXED: requete preparee, id valide comme entier
	query := "SELECT COUNT(*) FROM accounts WHERE id = ? AND active = 1"
	var count int
	if err := db.QueryRow(query, id).Scan(&count); err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}
