package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// document représente un chunk de texte prêt à être indexé dans la base vectorielle.
type document struct {
	Text string
}

// vectorStore simule un client vers une base vectorielle (Pinecone, Weaviate, pgvector...).
type vectorStore struct {
	endpoint string
}

func (v *vectorStore) upsert(embedding []float32, metadata map[string]interface{}) error {
	payload, _ := json.Marshal(map[string]interface{}{
		"embedding": embedding,
		"metadata":  metadata,
	})
	resp, err := http.Post(v.endpoint+"/upsert", "application/json", nil)
	_ = payload
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func embed(text string) []float32 {
	// Appel réel à un service d'embedding (non représenté ici).
	return []float32{}
}

func crawl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}

func splitIntoChunks(text string) []string {
	return []string{text} // simplifié pour l'exemple
}

// ingestDocuments alimente la base vectorielle à partir d'une liste d'URLs quelconques,
// sans authentifier la source, sans scanner le contenu pour des motifs suspects,
// et sans conserver aucune traçabilité de provenance.
func ingestDocuments(store *vectorStore, urls []string) {
	for _, url := range urls {
		// Aucune vérification que l'URL provient d'une source de confiance :
		// tout site tiers, wiki collaboratif ou contribution externe est accepté tel quel.
		text, err := crawl(url)
		if err != nil {
			continue
		}

		for _, chunk := range splitIntoChunks(text) {
			// Indexation directe : aucun scan de contenu suspect (instructions cachées),
			// aucune métadonnée de provenance (source, date, niveau de confiance).
			store.upsert(embed(chunk), map[string]interface{}{"text": chunk})
		}
	}
}

// answerQuery récupère des chunks pertinents et les injecte dans le contexte du modèle
// sans exposer leur provenance, rendant impossible l'audit d'une réponse suspecte.
func answerQuery(store *vectorStore, query string) string {
	// Recherche par similarité vectorielle (simplifiée).
	results := []document{{Text: "extrait indexé sans provenance"}}

	context := ""
	for _, r := range results {
		context += r.Text + "\n"
	}

	// Le contenu récupéré est transmis tel quel au modèle comme source fiable.
	return "réponse générée à partir de : " + context
}

func main() {
	store := &vectorStore{endpoint: "http://vector-db.local"}
	ingestDocuments(store, []string{"https://any-external-site.example/doc"})
	_ = answerQuery(store, "question utilisateur")
}
