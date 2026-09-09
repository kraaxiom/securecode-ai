// vulnerable.go - CWE-89: Error-based SQL Injection
// Les erreurs SQL brutes sont renvoyees au client, et l'entree utilisateur
// est concatenee, ce qui permet d'utiliser les messages d'erreur pour
// extraire des donnees (ex: CAST, EXTRACTVALUE).
package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func GetOrderHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("order_id")

	// VULNERABLE: concatenation + fuite du message d'erreur SQL brut
	query := fmt.Sprintf("SELECT total FROM orders WHERE id = %s", orderID)
	var total float64
	if err := db.QueryRow(query).Scan(&total); err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "total=%.2f", total)
}
