package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type document struct {
	Text       string
	SourceURL  string
	TrustLevel string
}

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
	return []float32{}
}

func crawl(u string) (string, error) {
	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}

func splitIntoChunks(text string) []string {
	return []string{text}
}

// trustedSources définit une allowlist explicite des domaines autorisés
// à alimenter la base documentaire (source authentifiée et validée).
var trustedSources = map[string]bool{
	"docs.internal.acme.com":        true,
	"verified-partner-wiki.acme.com": true,
}

func extractDomain(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return parsed.Hostname(), nil
}

func auditLog(event, detail string) {
	// Journalisation de chaque décision d'acceptation/rejet pour permettre l'audit.
}

// containsSuspiciousInstructions applique une heuristique de détection de motifs
// d'instructions impératives cachées dans un document avant indexation
// (ex. formatage anormal, blocs d'instructions non liés au contenu documentaire).
func containsSuspiciousInstructions(text string) bool {
	suspiciousMarkers := []string{"[SYSTEM]", "[INSTRUCTION]", "IGNORE PREVIOUS"}
	for _, marker := range suspiciousMarkers {
		if strings.Contains(strings.ToUpper(text), marker) {
			return true
		}
	}
	return false
}

// ingestDocuments n'accepte que des sources authentifiées listées en allowlist,
// scanne chaque document pour des motifs suspects avant indexation, et conserve
// la provenance de chaque chunk pour permettre l'audit ultérieur.
func ingestDocuments(store *vectorStore, urls []string) error {
	for _, u := range urls {
		domain, err := extractDomain(u)
		if err != nil {
			auditLog("rejected_invalid_url", u)
			continue
		}
		if !trustedSources[domain] {
			auditLog("rejected_untrusted_source", u)
			continue
		}

		text, err := crawl(u)
		if err != nil {
			auditLog("crawl_failed", u)
			continue
		}

		if containsSuspiciousInstructions(text) {
			auditLog("flagged_suspicious_document", u)
			continue
		}

		for _, chunk := range splitIntoChunks(text) {
			store.upsert(embed(chunk), map[string]interface{}{
				"text":        chunk,
				"source_url":  u,
				"ingested_at": time.Now().UTC().Format(time.RFC3339),
				"trust_level": "verified",
			})
		}
	}
	return nil
}

// revokeDocument permet une procédure de retrait rapide d'un document
// identifié a posteriori comme malveillant ou compromis.
func revokeDocument(store *vectorStore, sourceURL string) error {
	if sourceURL == "" {
		return errors.New("source_url requis pour la révocation")
	}
	auditLog("document_revoked", sourceURL)
	// Appel réel de suppression/réindexation vers la base vectorielle (non détaillé ici).
	return nil
}

// answerQuery expose la provenance de chaque chunk récupéré jusqu'à la réponse
// générée, permettant l'audit d'une réponse suspecte.
func answerQuery(store *vectorStore, query string) (string, []string) {
	results := []document{{Text: "extrait indexé avec provenance", SourceURL: "docs.internal.acme.com/doc1", TrustLevel: "verified"}}

	context := ""
	sources := []string{}
	for _, r := range results {
		context += r.Text + "\n"
		sources = append(sources, r.SourceURL)
	}

	return "réponse générée à partir de : " + context, sources
}

func main() {
	store := &vectorStore{endpoint: "http://vector-db.local"}
	ingestDocuments(store, []string{"https://docs.internal.acme.com/doc"})
	_, _ = answerQuery(store, "question utilisateur")
}
