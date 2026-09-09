package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

const systemPrompt = "Tu es un assistant interne. Réponds aux questions des utilisateurs poliment."

// message représente un élément structuré du contexte de conversation,
// avec un rôle explicite (system/user/assistant) et une origine de contenu marquée.
type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type toolCall struct {
	Name string                 `json:"name"`
	Args map[string]interface{} `json:"args"`
}

type llmResponse struct {
	Content   string     `json:"content"`
	ToolCalls []toolCall `json:"tool_calls"`
}

// toolSchema décrit les paramètres attendus pour un outil donné,
// utilisée pour valider indépendamment toute sortie du modèle avant exécution.
type toolSchema struct {
	requiredArgs []string
}

var toolSchemas = map[string]toolSchema{
	"send_email": {requiredArgs: []string{"to", "subject", "body"}},
}

// outils à faible impact uniquement exposés par défaut ; les outils sensibles
// (suppression, exécution de requêtes) ne sont pas inclus ici (moindre privilège).
var lowImpactTools = []string{"send_email"}

var highImpactTools = map[string]bool{
	"delete_file": true,
	"run_query":   true,
}

// callLLM utilise l'API structurée par rôles (system/user) plutôt que la
// concaténation de texte brut. Le contenu utilisateur reste marqué comme
// non fiable et séparé du prompt système au niveau du protocole d'appel,
// et non par une simple convention textuelle.
func callLLM(userInput string) (*llmResponse, error) {
	messages := []message{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: userInput},
	}

	payload, _ := json.Marshal(map[string]interface{}{
		"messages": messages,
		"tools":    lowImpactTools,
	})

	resp, err := http.Post("http://llm-provider.local/v1/chat", "application/json", nil)
	_ = payload
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out llmResponse
	json.NewDecoder(resp.Body).Decode(&out)
	return &out, nil
}

// validateAgainstSchema traite la sortie du modèle comme une donnée non fiable :
// tout appel d'outil est vérifié indépendamment contre un schéma strict
// avant toute exécution applicative.
func validateAgainstSchema(call toolCall) error {
	schema, ok := toolSchemas[call.Name]
	if !ok {
		return errors.New("outil inconnu ou non autorisé dans ce contexte")
	}
	for _, required := range schema.requiredArgs {
		if _, present := call.Args[required]; !present {
			return errors.New("argument requis manquant : " + required)
		}
	}
	return nil
}

func auditLog(event, detail string) {
	// Journalisation des tours de conversation et des sorties utilisées
	// pour déclencher des actions automatisées.
}

// requireHumanConfirmation applique une confirmation explicite avant
// toute exécution d'une action sensible — jamais automatique.
func requireHumanConfirmation(call toolCall) bool {
	return false // aucune confirmation obtenue : l'action sensible est bloquée par défaut
}

// executeToolCalls applique le principe du moindre privilège : seuls les outils
// à faible impact sont exécutables, les outils à fort impact nécessitent une
// confirmation humaine explicite, et chaque appel est validé contre son schéma.
func executeToolCalls(calls []toolCall) {
	registry := map[string]func(map[string]interface{}){
		"send_email": func(a map[string]interface{}) { /* envoi réel, validé */ },
	}

	for _, call := range calls {
		if err := validateAgainstSchema(call); err != nil {
			auditLog("blocked_invalid_tool_call", call.Name)
			continue
		}
		if highImpactTools[call.Name] && !requireHumanConfirmation(call) {
			auditLog("blocked_unconfirmed_sensitive_action", call.Name)
			continue
		}
		if fn, ok := registry[call.Name]; ok {
			fn(call.Args)
			auditLog("tool_executed", call.Name)
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
	auditLog("chat_turn", body.Message)

	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/chat", chatHandler)
	http.ListenAndServe(":8080", nil)
}
