// fixed.go - Correction CWE-639
// Le droit d'acces au document est verifie explicitement cote serveur
// (propriete ou partage explicite) avant de le renvoyer, quel que soit
// l'identifiant fourni par le client.
package handlers

import (
	"encoding/json"
	"net/http"
)

func GetDocumentHandler(w http.ResponseWriter, r *http.Request) {
	docID := r.URL.Query().Get("doc_id")
	currentUser := getCurrentUser(r)
	if currentUser == nil {
		http.Error(w, "non authentifie", http.StatusUnauthorized)
		return
	}

	doc, err := getDocumentByID(docID)
	if err != nil {
		http.Error(w, "document introuvable", http.StatusNotFound)
		return
	}

	// FIXED: verification explicite du droit d'acces (proprietaire ou
	// partage explicite) avant de renvoyer le contenu
	if !userCanAccessDocument(currentUser.ID, doc) {
		http.Error(w, "document introuvable", http.StatusNotFound) // pas de fuite d'existence
		return
	}
	json.NewEncoder(w).Encode(doc)
}
