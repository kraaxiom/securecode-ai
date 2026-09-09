// vulnerable.go - CWE-639: Authorization Bypass Through User-Controlled Key
// L'objet est recupere directement via l'identifiant fourni par le client,
// sans jamais verifier que l'utilisateur authentifie en est bien le proprietaire.
package handlers

import (
	"encoding/json"
	"net/http"
)

func GetInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	invoiceID := r.URL.Query().Get("id")

	// VULNERABLE: aucune verification que la facture appartient a l'utilisateur
	// authentifie, l'ID est utilise tel quel pour la recherche
	invoice, err := getInvoiceByID(invoiceID)
	if err != nil {
		http.Error(w, "facture introuvable", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(invoice)
}
