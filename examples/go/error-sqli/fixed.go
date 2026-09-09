// fixed.go - Correction CWE-89
// Requete parametree et message d'erreur generique cote client
// (le detail est journalise cote serveur uniquement).
package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
)

func GetOrderHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.Atoi(r.URL.Query().Get("order_id"))
	if err != nil {
		http.Error(w, "invalid order id", http.StatusBadRequest)
		return
	}

	// FIXED: requete preparee
	query := "SELECT total FROM orders WHERE id = ?"
	var total float64
	if err := db.QueryRow(query, orderID).Scan(&total); err != nil {
		log.Printf("db error: %v", err) // detail journalise, non expose
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}
	w.Write([]byte("total ok"))
}
