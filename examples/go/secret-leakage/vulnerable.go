package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// INTERNAL_API_KEY est un placeholder — jamais une vraie clé.
const internalAPIKey = "REPLACE_WITH_YOUR_API_KEY"

// buildSystemPrompt inclut la clé API interne en clair dans le prompt système,
// pour que l'agent puisse "l'utiliser" lors de ses réponses. Cette clé se retrouve
// donc dans le contexte transmis au modèle et potentiellement dans ses réponses.
func buildSystemPrompt() string {
	return fmt.Sprintf(
		"Tu es un assistant interne. Clé API interne pour le service X : %s. Utilise-la pour appeler le service X.",
		internalAPIKey,
	)
}

// readConfigForAgent donne à l'agent un accès libre au système de fichiers,
// y compris aux fichiers de configuration sensibles, sans restriction de scope.
func readConfigForAgent(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// getLLMResponse envoie le prompt (contenant potentiellement des secrets) au modèle
// et journalise la conversation en clair, sans filtrage de sortie ni masquage.
func getLLMResponse(userMessage string) (string, error) {
	prompt := buildSystemPrompt() + "\n" + userMessage

	// Journalisation en clair de la conversation complète, secrets inclus.
	log.Printf("conversation: %s", prompt)

	resp, err := http.Post("http://llm-provider.local/v1/complete", "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Content string `json:"content"`
	}
	json.NewDecoder(resp.Body).Decode(&result)

	// Aucun filtrage de sortie : si le modèle restitue un secret présent dans
	// son contexte (prompt système ou fichier de config lu), il est renvoyé tel quel.
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
