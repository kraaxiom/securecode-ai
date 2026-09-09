// vulnerable.go - CWE-349: Acceptance of Extraneous Untrusted Data With Trust
// Le pipeline de fine-tuning ingere des donnees depuis des sources externes
// sans verifier leur provenance ni leur integrite, et reintegre les retours
// utilisateurs directement dans le jeu d'entrainement sans filtrage.
package training

import (
	"fmt"
	"io"
	"net/http"
)

// DataSource represente une source de donnees pour le fine-tuning.
type DataSource struct {
	ID  string
	URL string
}

// Example est un exemple d'entrainement (prompt/reponse).
type Example struct {
	Prompt   string
	Response string
}

const apiKey = "REPLACE_WITH_YOUR_API_KEY"

// BuildTrainingSet telecharge et agrege les donnees de chaque source pour
// constituer le jeu de fine-tuning.
func BuildTrainingSet(sources []DataSource) ([]Example, error) {
	var dataset []Example

	for _, src := range sources {
		// VULNERABLE : aucune verification de provenance ni d'integrite
		// (pas de hash/signature attendue), aucune source de confiance
		// definie, aucune detection d'anomalies statistiques.
		batch, err := fetchAndParse(src.URL)
		if err != nil {
			return nil, fmt.Errorf("erreur source %s: %w", src.ID, err)
		}
		dataset = append(dataset, batch...)
	}

	return dataset, nil
}

// IngestUserFeedback reintegre directement les retours utilisateurs dans
// le jeu d'entrainement, sans echantillonnage humain ni filtrage.
func IngestUserFeedback(feedback []Example) []Example {
	// VULNERABLE : les retours utilisateurs (potentiellement manipules par
	// un attaquant en soumettant massivement des exemples biaises) sont
	// reinjectes tels quels dans le pipeline d'entrainement.
	return feedback
}

func fetchAndParse(url string) ([]Example, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	_ = body
	// Parsing des donnees brutes en exemples d'entrainement (details omis).
	return []Example{}, nil
}
