// vulnerable.go - CWE-89: Time-based Blind SQL Injection
// L'entree utilisateur est concatenee dans la requete ; un attaquant
// peut injecter des fonctions de delai (SLEEP/WAITFOR) pour deduire
// des informations via le temps de reponse.
package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func FilterReportHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	// VULNERABLE: concatenation directe, aucune limite de temps d'execution
	query := fmt.Sprintf("SELECT id FROM reports WHERE category = '%s'", category)
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	w.Write([]byte("ok"))
}
