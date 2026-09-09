// fixed.go - Correction CWE-89
// Requete parametree unique : le pilote n'autorise pas l'empilement
// de plusieurs instructions dans un seul appel prepare.
package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
)

func UpdateNicknameHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	nickname := r.URL.Query().Get("nickname")

	// FIXED: requete preparee avec parametres lies
	query := "UPDATE users SET nickname = ? WHERE id = ?"
	if _, err := db.Exec(query, nickname, userID); err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("updated"))
}
