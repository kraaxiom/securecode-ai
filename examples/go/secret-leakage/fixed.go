package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// buildSystemPrompt ne contient plus aucun secret en clair : l'agent est informé
// qu'il doit passer par un outil dédié, dont le secret réel est résolu
// uniquement côté application, jamais exposé dans le contexte du modèle.
func buildSystemPrompt() string {
	return "Tu es un assistant interne. Pour appeler le service X, utilise l'outil callServiceX " +
		"(le secret est géré par l'application et n'est jamais exposé dans ce contexte)."
}

// allowedConfigDirs restreint strictement les répertoires accessibles à l'agent,
// excluant explicitement tout répertoire de configuration sensible.
var allowedConfigDirs = []string{"/app/public-config"}

func isWithinAllowedDirs(path string) bool {
	for _, dir := range allowedConfigDirs {
		if strings.HasPrefix(path, dir) {
			return true
		}
	}
	return false
}

func auditLog(event, detail string) {
	log.Printf("[audit] %s: %s", event, detail)
}

// readConfigForAgent refuse tout accès en dehors de l'allowlist explicite
// de répertoires publics non sensibles.
func readConfigForAgent(path string) (string, error) {
	if !isWithinAllowedDirs(path) {
		auditLog("blocked_sensitive_file_access", path)
		return "", errors.New("accès refusé à ce chemin")
	}
	return "", nil // lecture réelle omise, non pertinente pour l'exemple
}

// resolveSecret résout un secret depuis un gestionnaire de secrets dédié
// (vault, secret manager), jamais depuis une variable en dur ou un prompt.
func resolveSecret(name string) string {
	// Résolution réelle via un secret manager (Vault, AWS Secrets Manager, etc.).
	return "REPLACE_WITH_YOUR_API_KEY"
}

// callServiceX est l'outil dédié : le secret est injecté côté application,
// après la décision du modèle, jamais transmis dans le contexte conversationnel.
func callServiceX(args map[string]interface{}) (*http.Response, error) {
	req, err := http.NewRequest("POST", "http://service-x.local/api", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+resolveSecret("service_x"))
	return http.DefaultClient.Do(req)
}

var secretPattern = regexp.MustCompile(`(?i)(api[_-]?key|secret|token|bearer)\s*[:=]\s*\S+`)

// containsSecret filtre la sortie du modèle pour détecter des motifs de secrets
// avant tout renvoi à l'utilisateur.
func containsSecret(text string) bool {
	return secretPattern.MatchString(text)
}

// getLLMResponse envoie un prompt système sans secret, masque les secrets
// potentiels dans les journaux, et filtre la sortie du modèle avant renvoi.
func getLLMResponse(userMessage string) (string, error) {
	prompt := buildSystemPrompt() + "\n" + userMessage

	// Journalisation avec masquage : le message utilisateur est loggé,
	// mais aucun secret n'est jamais présent dans le prompt système désormais.
	auditLog("conversation_started", "message reçu")

	resp, err := http.Post("http://llm-provider.local/v1/complete", "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Content string `json:"content"`
	}
	json.NewDecoder(resp.Body).Decode(&result)

	if containsSecret(result.Content) {
		auditLog("blocked_secret_in_output", "réponse filtrée")
		return "réponse indisponible : contenu sensible détecté", nil
	}

	return result.Content, nil
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Message string `json:"message"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	reply, err := getLLMResponse(body.Message)
	if err != nil {
		http.Error(w, "erreur interne", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"reply": reply})
}

func main() {
	http.HandleFunc("/chat", chatHandler)
	http.ListenAndServe(":8080", nil)
}
