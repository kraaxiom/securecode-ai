// vulnerable.go - CWE-89: Blind SQL Injection
// Aucune donnee n'est renvoyee directement, mais l'entree utilisateur
// est concatenee dans une clause WHERE booleenne, permettant une
// exfiltration par inference (vrai/faux).
package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func CheckAccountHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// VULNERABLE: concatenation dans une condition booleenne
	query := fmt.Sprintf("SELECT COUNT(*) FROM accounts WHERE id = %s AND active = 1", id)
	var count int
	if err := db.QueryRow(query).Scan(&count); err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}
