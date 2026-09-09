package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type toolCall struct {
	Name string                 `json:"name"`
	Args map[string]interface{} `json:"args"`
}

type tool func(args map[string]interface{}) (interface{}, error)

var toolRegistry = map[string]tool{
	"send_email":   func(args map[string]interface{}) (interface{}, error) { return "envoyé", nil },
	"query_db":     func(args map[string]interface{}) (interface{}, error) { return "résultat", nil },
	"call_webhook": func(args map[string]interface{}) (interface{}, error) { return "appelé", nil },
}

// toolSchema décrit précisément les arguments attendus et leur type pour un outil,
// utilisée pour rejeter tout argument hors périmètre avant exécution.
type toolSchema struct {
	requiredArgs map[string]string // nom -> type attendu ("string", "number", ...)
}

var toolSchemas = map[string]toolSchema{
	"send_email": {requiredArgs: map[string]string{"to": "string", "subject": "string", "body": "string"}},
	"query_db":   {requiredArgs: map[string]string{"query_id": "string"}},
}

// trustedToolSources n'autorise le chargement de définitions d'outils que depuis
// des registres internes vérifiés et revus, jamais depuis une source externe arbitraire.
var trustedToolSources = map[string]bool{
	"internal-registry": true,
}

func auditLog(event, detail string) {
	log.Printf("[audit] %s: %s", event, detail)
}

// loadToolDefinitions refuse toute source de définitions d'outils qui n'est pas
// explicitement listée comme source de confiance vérifiée.
func loadToolDefinitions(sourceID string) error {
	if !trustedToolSources[sourceID] {
		auditLog("rejected_untrusted_tool_source", sourceID)
		return errors.New("source d'outil non fiable")
	}
	return nil
}

// trustLevel représente le niveau de confiance du contenu en cours de traitement
// (ex. "user_direct" vs "retrieved_content" vs "tool_output").
type trustLevel string

const (
	trustUserDirect  trustLevel = "user_direct"
	trustLowConfidence trustLevel = "low_confidence"
)

// allowedToolsByTrust applique une liste blanche contextuelle : le contenu de
// faible confiance (par exemple la sortie réinjectée d'un outil précédent)
// ne peut déclencher qu'un sous-ensemble restreint d'outils.
var allowedToolsByTrust = map[trustLevel]map[string]bool{
	trustUserDirect:    {"send_email": true, "query_db": true, "call_webhook": true},
	trustLowConfidence: {"query_db": true},
}

// validateArgs vérifie chaque argument contre le schéma déclaré de l'outil,
// indépendamment de ce que le modèle a généré.
func validateArgs(call toolCall) (map[string]interface{}, error) {
	schema, ok := toolSchemas[call.Name]
	if !ok {
		return nil, errors.New("outil sans schéma déclaré : exécution refusée")
	}
	validated := map[string]interface{}{}
	for name := range schema.requiredArgs {
		val, present := call.Args[name]
		if !present {
			return nil, errors.New("argument requis manquant : " + name)
		}
		validated[name] = val
	}
	return validated, nil
}

// executeToolCall applique le schéma strict et la liste blanche contextuelle
// avant toute exécution, et journalise chaque appel réel pour permettre l'audit.
func executeToolCall(call toolCall, level trustLevel) (interface{}, error) {
	if !allowedToolsByTrust[level][call.Name] {
		auditLog("blocked_tool_not_in_allowlist", call.Name)
		return nil, errors.New("outil non autorisé pour ce niveau de confiance")
	}

	validatedArgs, err := validateArgs(call)
	if err != nil {
		auditLog("blocked_invalid_tool_args", call.Name)
		return nil, err
	}

	fn, ok := toolRegistry[call.Name]
	if !ok {
		return nil, errors.New("outil inconnu")
	}

	result, err := fn(validatedArgs)
	if err == nil {
		auditLog("tool_executed", call.Name)
	}
	return result, err
}

// pipelineHandler abaisse le niveau de confiance dès qu'un résultat d'outil
// est réinjecté comme entrée d'un appel suivant, limitant ainsi l'ensemble
// des outils exécutables à partir de contenu non directement fourni par
// l'utilisateur authentifié.
func pipelineHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ToolCalls []toolCall `json:"tool_calls"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	level := trustUserDirect
	var lastResult interface{}

	for i, call := range body.ToolCalls {
		if i > 0 {
			// À partir du deuxième appel, les paramètres peuvent avoir été
			// influencés par la sortie du premier outil : confiance réduite.
			level = trustLowConfidence
		}

		result, err := executeToolCall(call, level)
		if err != nil {
			http.Error(w, "appel d'outil rejeté", http.StatusForbidden)
			return
		}
		lastResult = result
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"result": lastResult})
}

func main() {
	http.HandleFunc("/pipeline", pipelineHandler)
	http.ListenAndServe(":8080", nil)
}
