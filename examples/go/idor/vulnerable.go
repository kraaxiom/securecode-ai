// vulnerable.go - CWE-639: Authorization Bypass Through User-Controlled Key
// Le document est recupere uniquement a partir de l'identifiant fourni par
// le client dans l'URL, sans jamais verifier que l'utilisateur authentifie
// a le droit d'y acceder (Insecure Direct Object Reference).
package handlers

import (
	"encoding/json"
	"net/http"
)

func GetDocumentHandler(w http.ResponseWriter, r *http.Request) {
	docID := r.URL.Query().Get("doc_id")

	// VULNERABLE: reference directe non securisee, aucun controle de
	// propriete ou de droit d'acces sur le document demande
	doc, err := getDocumentByID(docID)
	if err != nil {
		http.Error(w, "document introuvable", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(doc)
}
