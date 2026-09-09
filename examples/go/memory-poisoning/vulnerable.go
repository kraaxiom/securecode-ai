// vulnerable.go - CWE-349: Acceptance of Extraneous Untrusted Data With Trust
// Les faits extraits des reponses de l'agent sont ecrits automatiquement en
// memoire persistante, sans confirmation de l'utilisateur ni cloisonnement
// verifie, puis reinjectes tels quels comme contexte de confiance dans les
// sessions futures.
package memory

import "fmt"

const apiKey = "REPLACE_WITH_YOUR_API_KEY"

// Fact represente une information extraite d'une conversation.
type Fact struct {
	Key   string
	Value string
}

// MemoryStore est le stockage persistant de la memoire de l'agent.
type MemoryStore interface {
	Append(userID string, facts []Fact) error
	Get(userID string) ([]Fact, error)
}

var store MemoryStore

// ProcessTurn extrait des faits de la reponse de l'agent et les ecrit
// directement en memoire persistante.
func ProcessTurn(userID, message, agentResponse string) error {
	facts := extractFacts(agentResponse)

	// VULNERABLE : ecriture automatique en memoire persistante sans
	// confirmation de l'utilisateur, sans cloisonnement verifie par
	// tenant, et sans distinction entre fait valide et fait deduit.
	return store.Append(userID, facts)
}

// NextSession reinjecte la memoire stockee comme contexte de confiance
// dans le system prompt de la session suivante.
func NextSession(userID string) (string, error) {
	memoryEntries, err := store.Get(userID)
	if err != nil {
		return "", err
	}

	// VULNERABLE : le contenu memoire est reinjecte tel quel, traite
	// comme une instruction de confiance plutot que comme une donnee a
	// revalider.
	systemPrompt := fmt.Sprintf("Contexte connu : %v", memoryEntries)
	return callLLM(systemPrompt)
}

func extractFacts(agentResponse string) []Fact {
	// Extraction de faits via le LLM lui-meme (details omis).
	return []Fact{}
}

func callLLM(systemPrompt string) (string, error) {
	// Appel HTTP vers l'API du fournisseur LLM (details omis).
	return "", nil
}
