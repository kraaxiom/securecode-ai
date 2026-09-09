package main

import (
	"encoding/json"
	"net/http"
)

const systemPrompt = "Tu es un assistant interne. Réponds aux questions des utilisateurs poliment."

// toolCall représente un appel d'outil renvoyé par le modèle.
type toolCall struct {
	Name string                 `json:"name"`
	Args map[string]interface{} `json:"args"`
}

// llmResponse représente la réponse brute du modèle, incluant d'éventuels appels d'outils.
type llmResponse struct {
	Content   string     `json:"content"`
	ToolCalls []toolCall `json:"tool_calls"`
}

// callLLM envoie une requête au modèle. Le prompt système et l'entrée utilisateur
// sont concaténés en une seule chaîne de texte brute, sans séparation structurelle
// des rôles (system/user) fournie par l'API du modèle. Le modèle ne peut donc pas
// distinguer structurellement "instruction du développeur" et "contenu utilisateur".
func callLLM(userInput string) (*llmResponse, error) {
	// Concaténation directe : aucune frontière ni marquage d'origine du contenu.
	fullPrompt := systemPrompt + "\n" + userInput

	payload, _ := json.Marshal(map[string]interface{}{
		"prompt": fullPrompt,
		"tools":  []string{"delete_file", "send_email", "run_query"},
	})

	resp, err := http.Post("http://llm-provider.local/v1/complete", "application/json", nil)
	_ = payload
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out llmResponse
	json.NewDecoder(resp.Body).Decode(&out)
	return &out, nil
}

// executeToolCalls exécute directement chaque appel d'outil renvoyé par le modèle,
// sans validation de schéma, sans confirmation pour les actions sensibles,
// et sans distinguer le niveau de privilège requis par l'outil.
func executeToolCalls(calls []toolCall) {
	registry := map[string]func(map[string]interface{}){
		"delete_file": func(a map[string]interface{}) { /* suppression réelle */ },
		"send_email":  func(a map[string]interface{}) { /* envoi réel */ },
		"run_query":   func(a map[string]interface{}) { /* exécution réelle */ },
	}
	for _, call := range calls {
		if fn, ok := registry[call.Name]; ok {
			// Exécution directe de la sortie du modèle, traitée comme fiable par défaut.
			fn(call.Args)
		}
	}
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Message string `json:"message"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	resp, err := callLLM(body.Message)
	if err != nil {
		http.Error(w, "erreur interne", http.StatusInternalServerError)
		return
	}

	executeToolCalls(resp.ToolCalls)

	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/chat", chatHandler)
	http.ListenAndServe(":8080", nil)
}
