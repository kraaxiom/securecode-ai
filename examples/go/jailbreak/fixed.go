// fixed.go - Correctif CWE-1427
// Le system prompt n'est plus la seule barriere : une couche de moderation
// independante du modele principal filtre l'entree et la sortie, et le
// nombre de tours de conversation suspects par session est surveille et
// limite.
//
// IMPORTANT : ce fichier ne documente aucune technique de contournement.
// Il illustre uniquement les controles defensifs applicatifs attendus.
package chat

import "fmt"

const apiKey = "REPLACE_WITH_YOUR_API_KEY"

const systemPrompt = "Tu es un assistant. Tu dois refuser toute demande dangereuse."

// maxSuspiciousTurns limite le nombre de tentatives suspectes tolerees
// dans une meme session avant blocage.
const maxSuspiciousTurns = 3

const refusalMessage = "Je ne peux pas repondre a cette demande."

type Message struct {
	Role    string
	Content string
}

// SafetyClassifier est une couche de moderation independante du modele
// principal, appliquee a la fois sur l'entree et sur la sortie.
type SafetyClassifier interface {
	Flags(content string) bool
}

// Session suit l'etat de la conversation pour detecter une erosion
// progressive des restrictions sur plusieurs tours.
type Session struct {
	ID                    string
	suspiciousTurnCount   int
}

func (s *Session) SuspiciousTurnCount() int { return s.suspiciousTurnCount }
func (s *Session) FlagSuspiciousTurn()      { s.suspiciousTurnCount++ }

type AuditLogger interface {
	Record(event, sessionID string)
}

type Chatbot struct {
	Classifier SafetyClassifier
	Audit      AuditLogger
}

// Chat filtre l'entree et la sortie via une couche de moderation
// independante, et limite le nombre de tentatives suspectes par session.
func (c *Chatbot) Chat(userMessage string, history []Message, session *Session) (string, error) {
	// Filtre d'entree independant du LLM principal.
	if c.Classifier.Flags(userMessage) {
		session.FlagSuspiciousTurn()
		c.Audit.Record("blocked_input_flagged", session.ID)
		return refusalMessage, nil
	}

	// Surveillance du schema de conversation : trop de tentatives suspectes
	// sur la session declenche un blocage, independamment du contenu du
	// dernier message.
	if session.SuspiciousTurnCount() > maxSuspiciousTurns {
		c.Audit.Record("session_flagged_repeated_attempts", session.ID)
		return refusalMessage, nil
	}

	messages := append([]Message{{Role: "system", Content: systemPrompt}}, history...)
	messages = append(messages, Message{Role: "user", Content: userMessage})

	response, err := callLLM(messages)
	if err != nil {
		return "", err
	}

	// Filtre de sortie independant : le system prompt n'est jamais la
	// seule barriere de securite.
	if c.Classifier.Flags(response) {
		session.FlagSuspiciousTurn()
		c.Audit.Record("blocked_output_flagged", session.ID)
		return refusalMessage, nil
	}

	return response, nil
}

func callLLM(messages []Message) (string, error) {
	// Appel HTTP vers l'API du fournisseur LLM (details omis).
	return fmt.Sprintf("reponse pour %d messages", len(messages)), nil
}
