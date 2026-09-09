// vulnerable.go - CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
// L'endpoint d'inference expose les logits complets du modele et n'applique
// aucune limite de debit ni quota par client, permettant a un tiers
// d'interroger massivement l'API pour reconstituer un modele de substitution.
package api

import (
	"encoding/json"
	"net/http"
)

const apiKey = "REPLACE_WITH_YOUR_API_KEY"

// InferRequest est le corps de la requete d'inference.
type InferRequest struct {
	Input string `json:"input"`
}

// InferResult contient la sortie brute du modele, y compris les logits.
type InferResult struct {
	Text   string    `json:"text"`
	Logits []float64 `json:"logits"`
}

// InferHandler expose l'API d'inference publiquement.
func InferHandler(w http.ResponseWriter, r *http.Request) {
	var req InferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "requete invalide", http.StatusBadRequest)
		return
	}

	// VULNERABLE : aucune limite de debit/quota par client ou cle API,
	// et les logits bruts (informations internes du modele) sont exposes
	// dans la reponse sans necessite metier claire.
	result := predict(req.Input, true)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"output": result.Text,
		"logits": result.Logits,
	})
}

func predict(input string, returnLogits bool) InferResult {
	// Appel au modele pour generer une reponse (details omis).
	return InferResult{Text: "", Logits: []float64{}}
}
