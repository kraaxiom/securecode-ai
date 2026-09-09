package main

import (
	"encoding/json"
	"net/http"
)

// toolCall représente un appel d'outil tel que renvoyé par le modèle (function calling).
type toolCall struct {
	Name string                 `json:"name"`
	Args map[string]interface{} `json:"args"`
}

// tool est la signature d'une fonction exécutable par l'agent.
type tool func(args map[string]interface{}) (interface{}, error)

// toolRegistry associe chaque nom d'outil à son implémentation réelle.
var toolRegistry = map[string]tool{
	"send_email":   func(args map[string]interface{}) (interface{}, error) { return "envoyé", nil },
	"query_db":     func(args map[string]interface{}) (interface{}, error) { return "résultat", nil },
	"call_webhook": func(args map[string]interface{}) (interface{}, error) { return "appelé", nil },
}

// loadToolDefinitions charge dynamiquement des définitions d'outils depuis une
// source externe (registre de plugins tiers) sans vérifier qu'elle provient
// d'une source de confiance.
func loadToolDefinitions(sourceURL string) ([]byte, error) {
	resp, err := http.Get(sourceURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var raw []byte
	json.NewDecoder(resp.Body).Decode(&raw)
	return raw, nil
}

// executeToolCall transmet directement les arguments générés par le modèle
// à l'implémentation de l'outil, sans validation de schéma, sans liste blanche
// contextuelle selon le niveau de confiance du contenu traité, et sans
// journalisation des arguments réels utilisés.
func executeToolCall(call toolCall) (interface{}, error) {
	fn, ok := toolRegistry[call.Name]
	if !ok {
		return nil, nil
	}
	// Exécution directe : les arguments proviennent du texte généré par le modèle
	// (potentiellement influencé par du contenu externe non fiable, ex. résultat
	// d'un outil précédent réinjecté sans nettoyage) sans aucune étape de contrôle.
	return fn(call.Args)
}

// pipelineHandler enchaîne plusieurs outils où la sortie de l'un alimente
// directement les paramètres du suivant, sans étape de validation intermédiaire.
func pipelineHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ToolCalls []toolCall `json:"tool_calls"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	var lastResult interface{}
	for _, call := range body.ToolCalls {
		result, err := executeToolCall(call)
		if err != nil {
			http.Error(w, "erreur d'exécution d'outil", http.StatusInternalServerError)
			return
		}
		lastResult = result // réinjecté tel quel comme entrée du prochain appel, sans nettoyage
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"result": lastResult})
}

func main() {
	http.HandleFunc("/pipeline", pipelineHandler)
	http.ListenAndServe(":8080", nil)
}
