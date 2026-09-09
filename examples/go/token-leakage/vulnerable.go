package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// llmProviderAPIKey est codée en dur dans le binaire backend et journalisée
// en clair en cas d'erreur — un placeholder, jamais une vraie clé.
const llmProviderAPIKey = "REPLACE_WITH_YOUR_API_KEY"

// askAssistant appelle directement le fournisseur LLM en transmettant la clé
// d'API en clair dans l'en-tête, sans passer par un proxy backend dédié qui
// isolerait le secret d'un éventuel client non fiable.
func askAssistant(question string) (*http.Response, error) {
	req, err := http.NewRequest("POST", "https://llm-provider.local/v1/messages", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+llmProviderAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// debugHandler journalise l'intégralité des en-têtes de la requête, y compris
// l'en-tête Authorization contenant le jeton, sans aucun masquage.
func debugHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := askAssistant("question de test")
	if err != nil {
		http.Error(w, "erreur interne", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Journalisation de débogage exposant l'en-tête d'autorisation en clair.
	log.Printf("réponse du fournisseur, en-têtes complets: %v", resp.Header)

	w.Write([]byte("ok"))
}

func assistantHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Question string `json:"question"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	resp, err := askAssistant(body.Question)
	if err != nil {
		http.Error(w, "erreur interne", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	json.NewEncoder(w).Encode(map[string]string{"status": "envoyé"})
}

func main() {
	http.HandleFunc("/api/assistant", assistantHandler)
	http.HandleFunc("/debug", debugHandler)
	http.ListenAndServe(":8080", nil)
}
