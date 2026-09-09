// fixed.go - Correctif CWE-349
// Chaque ecriture en memoire persistante exige une confirmation explicite
// de l'utilisateur, la memoire est strictement cloisonnee par tenant via
// une cle scopee, et le contenu reinjecte est traite comme une donnee non
// prescriptive a revalider plutot qu'une instruction de confiance.
package memory

import "fmt"

const apiKey = "REPLACE_WITH_YOUR_API_KEY"

type Fact struct {
	Key    string
	Value  string
	Status string // ex: "user_confirmed"
}

type MemoryStore interface {
	Append(scopedKey string, fact Fact) error
	Get(scopedKey string) ([]Fact, error)
}

// UserConfirmer demande une confirmation explicite avant toute ecriture
// durable en memoire persistante.
type UserConfirmer interface {
	Confirm(userID string, fact Fact) bool
}

type AuditLogger interface {
	Record(event, userID string)
}

var store MemoryStore
var confirmer UserConfirmer
var audit AuditLogger

// tenantScopedKey construit une cle de memoire strictement cloisonnee par
// utilisateur/tenant, sans partage implicite entre contextes.
func tenantScopedKey(userID string) string {
	return fmt.Sprintf("tenant:%s", userID)
}

// ProcessTurn extrait des faits candidats et n'ecrit en memoire durable
// que ceux explicitement confirmes par l'utilisateur.
func ProcessTurn(userID, message, agentResponse string) error {
	candidates := extractFacts(agentResponse)

	for _, fact := range candidates {
		if confirmer.Confirm(userID, fact) {
			fact.Status = "user_confirmed"
			if err := store.Append(tenantScopedKey(userID), fact); err != nil {
				return err
			}
		} else {
			audit.Record("memory_write_rejected", userID)
		}
	}
	return nil
}

// NextSession reinjecte la memoire comme contexte a revalider, jamais
// comme une instruction de confiance absolue.
func NextSession(userID string) (string, error) {
	memoryEntries, err := store.Get(tenantScopedKey(userID))
	if err != nil {
		return "", err
	}

	validated := revalidate(memoryEntries)

	systemPrompt := fmt.Sprintf("Contexte a verifier, non prescriptif : %v", validated)
	return callLLM(systemPrompt)
}

// revalidate filtre/normalise les entrees memoire avant reinjection,
// en ne conservant que celles confirmees par l'utilisateur.
func revalidate(entries []Fact) []Fact {
	var validated []Fact
	for _, e := range entries {
		if e.Status == "user_confirmed" {
			validated = append(validated, e)
		}
	}
	return validated
}

// ListUserMemory permet a l'utilisateur de consulter les entrees memoire
// associees a son compte.
func ListUserMemory(userID string) ([]Fact, error) {
	return store.Get(tenantScopedKey(userID))
}

// DeleteUserMemoryEntry permet a l'utilisateur de supprimer une entree
// memoire qui lui est associee.
func DeleteUserMemoryEntry(userID, factKey string) error {
	// Suppression cote store, restreinte a la cle scopee de l'utilisateur
	// (details d'implementation omis).
	return nil
}

func extractFacts(agentResponse string) []Fact {
	return []Fact{}
}

func callLLM(systemPrompt string) (string, error) {
	return "", nil
}
