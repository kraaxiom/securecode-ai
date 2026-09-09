// vulnerable.go - CWE-349: Acceptance of Extraneous Untrusted Data With Trust
// Le contenu soumis par les utilisateurs est directement transforme en
// embedding puis indexe dans la base vectorielle, sans moderation, sans
// limite de debit par source, et sans surveillance de la distribution
// des vecteurs.
package rag

import (
	"crypto/rand"
	"encoding/hex"
)

// Vector represente un embedding.
type Vector []float32

const apiKey = "REPLACE_WITH_YOUR_API_KEY"

// VectorStore represente la base vectorielle utilisee pour la recherche
// semantique.
type VectorStore interface {
	Upsert(id string, vector Vector, metadata map[string]string) error
}

// EmbeddingModel genere l'embedding d'un texte via l'API du fournisseur LLM.
type EmbeddingModel interface {
	Embed(text string) (Vector, error)
}

type Indexer struct {
	Embeddings EmbeddingModel
	Store      VectorStore
}

// IndexContent transforme le contenu utilisateur en embedding et l'indexe
// directement dans la base vectorielle.
func (idx *Indexer) IndexContent(userContent string, sourceID string) error {
	// VULNERABLE : aucune moderation du contenu avant generation de
	// l'embedding, aucune limite de frequence/volume par source, aucune
	// detection de clustering ou de densite anormale dans l'espace vectoriel.
	vector, err := idx.Embeddings.Embed(userContent)
	if err != nil {
		return err
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
