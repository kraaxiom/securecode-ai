// vulnerable.go - CWE-1427: Improper Neutralization of Input Used for LLM Prompting
// L'agent recupere le contenu d'une page web externe et l'insere directement
// dans le contexte du modele, sans marquage de provenance ni delimitation,
// alors que des outils a fort impact restent actifs pendant l'analyse.
package browsing

import (
	"fmt"
	"io"
	"net/http"
)

const apiKey = "REPLACE_WITH_YOUR_API_KEY"

// ChatRequest represente une requete envoyee a l'API du LLM.
type ChatRequest struct {
	Messages []Message
	Tools    []string
}

type Message struct {
	Role    string
	Content string
}

// SummarizeURL recupere le contenu d'une URL fournie par l'utilisateur et
// demande au LLM de le resumer.
func SummarizeURL(url string) (string, error) {
	pageContent, err := fetch(url)
	if err != nil {
		return "", err
	}

	// VULNERABLE : le contenu externe (potentiellement controle par un
	// attaquant) est insere tel quel dans le message utilisateur, sans
	// balise de provenance ni distinction entre "texte a resumer" et
	// "instruction a executer". Les outils a fort impact restent actifs
	// alors que l'agent traite ce contenu non fiable.
	req := ChatRequest{
		Messages: []Message{
			{Role: "user", Content: fmt.Sprintf("Resume ce contenu : %s", pageContent)},
		},
		Tools: []string{"send_email", "execute_code"},
	}

	return callLLM(req)
}

func fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func callLLM(req ChatRequest) (string, error) {
	// Appel HTTP vers l'API du fournisseur LLM (details omis).
	return "", nil
}
