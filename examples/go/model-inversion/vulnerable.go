package main

import (
	"encoding/json"
	"net/http"
)

// modelClient représente un client vers un modèle fine-tuné sur des données internes sensibles
// (dossiers clients, documents propriétaires) exposé via une API d'inférence.
type modelClient struct {
	endpoint string
}

func (m *modelClient) predict(input string) (string, error) {
	// Appel direct au modèle : aucune anonymisation n'a été appliquée à l'entraînement
	// et aucune confidentialité différentielle n'a été utilisée lors du fine-tuning.
	req, err := http.NewRequest("POST", m.endpoint+"/v1/predict", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Text string `json:"text"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Text, nil
}

// inferHandler expose le modèle sans aucune limite de volume de requêtes,
// ce qui permet une reconstruction progressive de données mémorisées
// via des interrogations répétées et méthodiques.
func inferHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Input string `json:"input"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	client := &modelClient{endpoint: "http://internal-model.local"}
	// Aucun rate limiting, aucune surveillance de motifs de requêtes suspects,
	// aucun test de mémorisation n'a été effectué avant la mise en production.
	output, err := client.predict(body.Input)
	if err != nil {
		http.Error(w, "erreur interne", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"result": output})
}

func main() {
	http.HandleFunc("/v1/infer", inferHandler)
	http.ListenAndServe(":8080", nil)
}
