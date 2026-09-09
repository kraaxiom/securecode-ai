// fixed.go - Correction CWE-89
// Requete parametree + timeout de contexte pour limiter l'impact
// de toute tentative d'injection de delai.
package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"time"
)

func FilterReportHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// FIXED: requete preparee + timeout borne
	query := "SELECT id FROM reports WHERE category = ?"
	rows, err := db.QueryContext(ctx, query, category)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	w.Write([]byte("ok"))
}
