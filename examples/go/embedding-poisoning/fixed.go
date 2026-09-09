// fixed.go - Correctif CWE-349
// Le contenu utilisateur passe par une moderation avant generation de
// l'embedding, une limite de debit par source freine les campagnes
// automatisees, et les vecteurs a densite anormale sont mis en
// quarantaine plutot qu'indexes.
package rag

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Vector []float32

const apiKey = "REPLACE_WITH_YOUR_API_KEY"

// rateLimitPerSource limite le nombre de vecteurs indexes par source et
// par heure, pour freiner les campagnes d'empoisonnement automatisees.
const rateLimitPerSource = 50

var errRejectedModeration = fmt.Errorf("contenu rejete par la moderation")
var errRateLimitExceeded = fmt.Errorf("limite de debit depassee pour cette source")

type VectorStore interface {
	Upsert(id string, vector Vector, metadata map[string]string) error
	IsAnomalousDensity(vector Vector) (bool, error)
	Quarantine(vector Vector, metadata map[string]string) error
}

type EmbeddingModel interface {
	Embed(text string) (Vector, error)
}

// ContentModeration verifie que le contenu utilisateur respecte les
// regles de moderation avant indexation.
type ContentModeration interface {
	IsAllowed(content string) bool
}

// RateLimiter suit le volume de vecteurs indexes par source.
type RateLimiter interface {
	Exceeded(sourceID string, limit int) bool
}

type AuditLogger interface {
	Record(event, sourceID string)
}

type Indexer struct {
	Embeddings EmbeddingModel
	Store      VectorStore
	Moderation ContentModeration
	RateLimit  RateLimiter
	Audit      AuditLogger
}

// IndexContent modere, limite en debit, embed puis verifie la densite
// avant d'indexer le contenu utilisateur.
func (idx *Indexer) IndexContent(userContent string, sourceID string) error {
	if !idx.Moderation.IsAllowed(userContent) {
		idx.Audit.Record("rejected_content_moderation", sourceID)
		return errRejectedModeration
	}

	if idx.RateLimit.Exceeded(sourceID, rateLimitPerSource) {
		idx.Audit.Record("rejected_rate_limit", sourceID)
		return errRateLimitExceeded
	}

	vector, err := idx.Embeddings.Embed(userContent)
	if err != nil {
		return err
	}

	anomalous, err := idx.Store.IsAnomalousDensity(vector)
	if err != nil {
		return err
	}
	if anomalous {
		// Proximite anormale avec des requetes frequentes : mise en
		// quarantaine plutot qu'indexation directe.
		idx.Audit.Record("flagged_anomalous_embedding", sourceID)
		return idx.Store.Quarantine(vector, map[string]string{"source": sourceID})
	}

	id, err := newID()
	if err != nil {
		return err
	}

	return idx.Store.Upsert(id, vector, map[string]string{"source": sourceID})
}

func newID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
