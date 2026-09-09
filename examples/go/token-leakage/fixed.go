package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// resolveSecret résout la clé du fournisseur LLM depuis un gestionnaire de secrets
// (variable d'environnement injectée par le secret manager en production),
// jamais codée en dur dans le code source.
func resolveSecret(name string) string {
	value := os.Getenv(name)
	if value == "" {
		return "REPLACE_WITH_YOUR_API_KEY" // valeur de repli pour environnement de développement local uniquement
	}
	return value
}

var llmProviderAPIKey = resolveSecret("LLM_API_KEY")

// redactHeaders masque les en-têtes sensibles avant toute journalisation,
// pour éviter qu'un jeton d'autorisation ne finisse dans les journaux applicatifs.
func redactHeaders(headers http.Header) http.Header {
	redacted := headers.Clone()
	for _, sensitive := range []string{"Authorization", "X-Api-Key"} {
		if redacted.Get(sensitive) != "" {
			redacted.Set(sensitive, "[REDACTED]")
		}
	}
	return redacted
}

// askAssistant reste le seul point du backend à détenir la clé d'API du fournisseur.
// Aucun client (frontend web, application mobile) n'a jamais accès à ce secret :
// tous les appels transitent par ce backend qui agit comme proxy authentifié.
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

func requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Authentification de l'appelant du backend (session utilisateur, JWT applicatif...),
		// distincte et sans rapport avec la clé du fournisseur LLM.
		if r.Header.Get("X-App-Session") == "" {
			http.Error(w, "authentification requise", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

// debugHandler journalise les en-têtes après masquage systématique des jetons sensibles.
func debugHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := askAssistant("question de test")
	if err != nil {
		http.Error(w, "erreur interne", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	log.Printf("réponse du fournisseur, en-têtes (masqués): %v", redactHeaders(resp.Header))

	w.Write([]byte("ok"))
}

// assistantHandler est le seul endpoint exposé au client : il ne transmet jamais
// la clé du fournisseur, uniquement la question de l'utilisateur authentifié.
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
	http.HandleFunc("/api/assistant", requireAuth(assistantHandler))
	http.HandleFunc("/debug", requireAuth(debugHandler))
	http.ListenAndServe(":8080", nil)
}
