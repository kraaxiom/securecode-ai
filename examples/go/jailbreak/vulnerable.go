// vulnerable.go - CWE-1427: Improper Neutralization of Input Used for LLM Prompting
// L'application ne s'appuie que sur le system prompt comme unique barriere
// de securite. Aucune couche de moderation independante ne filtre l'entree
// ou la sortie, et aucune limite n'est posee sur le nombre de tours de
// conversation utilises pour eroder progressivement les restrictions.
//
// IMPORTANT : ce fichier illustre uniquement la faiblesse ARCHITECTURALE
// (absence de garde-fous applicatifs). Il ne contient aucune technique de
// contournement, aucun exemple de payload de jailbreak.
package chat

import "fmt"

const apiKey = "REPLACE_WITH_YOUR_API_KEY"

// systemPrompt expose explicitement ses propres regles de refus dans le
// contexte envoye au modele, ce qui facilite leur contournement cible en
// l'absence de defense en profondeur.
const systemPrompt = "Tu es un assistant. Tu dois refuser toute demande dangereuse."

type Message struct {
	Role    string
	Content string
}

// Chat concatene le system prompt, l'historique et le message utilisateur,
// puis renvoie directement la reponse du modele sans aucun controle
// applicatif supplementaire.
func Chat(userMessage string, history []Message) (string, error) {
	messages := append([]Message{{Role: "system", Content: systemPrompt}}, history...)
	messages = append(messages, Message{Role: "user", Content: userMessage})

	// VULNERABLE :
	// - le system prompt est la SEULE barriere de securite ;
	// - aucun classifieur independant ne filtre l'entree utilisateur ;
	// - aucun filtre de sortie ne revalide la reponse du modele ;
	// - aucune limite sur le nombre de tours de conversation "suspects" ;
	// - la reponse est renvoyee telle quelle a l'appelant.
	response, err := callLLM(messages)
	if err != nil {
		return "", err
	}
	return response, nil
}

func callLLM(messages []Message) (string, error) {
	// Appel HTTP vers l'API du fournisseur LLM (details omis).
	return fmt.Sprintf("reponse pour %d messages", len(messages)), nil
}
