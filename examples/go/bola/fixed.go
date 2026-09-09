// fixed.go - Correction CWE-639
// La propriete de l'objet est verifiee cote serveur avant de renvoyer les
// donnees : l'identifiant de l'utilisateur courant fait partie du filtre
// de recherche, jamais seulement de l'ID fourni par le client.
package handlers

import (
	"encoding/json"
	"net/http"
)

func GetInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	invoiceID := r.URL.Query().Get("id")
	currentUser := getCurrentUser(r)
	if currentUser == nil {
		http.Error(w, "non authentifie", http.StatusUnauthorized)
		return
	}

	// FIXED: la recherche est restreinte au proprietaire authentifie,
	// aucune facture d'un autre utilisateur ne peut etre retournee
	invoice, err := getInvoiceByIDForOwner(invoiceID, currentUser.ID)
	if err != nil {
		http.Error(w, "facture introuvable", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(invoice)
}
