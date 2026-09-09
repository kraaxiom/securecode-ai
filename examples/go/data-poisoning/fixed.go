// fixed.go - Correctif CWE-349
// Chaque source de donnees d'entrainement est verifiee (provenance + hash),
// les donnees aberrantes sont filtrees avant integration, les retours
// utilisateurs passent par un echantillonnage humain, et les performances
// du modele sont comparees a un jeu de reference apres chaque entrainement.
package training

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

type DataSource struct {
	ID         string
	URL        string
	SignedHash string // hash attendu, signe par la source de confiance
}

type Example struct {
	Prompt   string
	Response string
}

const apiKey = "REPLACE_WITH_YOUR_API_KEY"

// ACCEPTABLE_DRIFT_THRESHOLD est le score minimal acceptable sur le jeu
// de reference apres un cycle d'entrainement.
const acceptableDriftThreshold = 0.85

// trustedSources liste les identifiants de sources autorisees pour le
// fine-tuning.
var trustedSources = map[string]bool{
	"internal-labeled-set":   true,
	"verified-partner-feed":  true,
}

var errModelDrift = fmt.Errorf("derive de comportement detectee apres entrainement")

// BuildTrainingSet ne conserve que les sources de confiance, verifie leur
// integrite et filtre les valeurs aberrantes avant integration.
func BuildTrainingSet(sources []DataSource) ([]Example, error) {
	var dataset []Example

	for _, src := range sources {
		if !trustedSources[src.ID] {
			recordAudit("rejected_untrusted_source", src.ID)
			continue
		}

		if err := verifyIntegrity(src); err != nil {
			recordAudit("rejected_integrity_check_failed", src.ID)
			continue
		}

		batch, err := fetchAndParse(src.URL)
		if err != nil {
			return nil, fmt.Errorf("erreur source %s: %w", src.ID, err)
		}

		batch = filterStatisticalOutliers(batch) // detection d'anomalies
		dataset = append(dataset, batch...)
	}

	return dataset, nil
}

// IngestUserFeedback isole les retours utilisateurs dans une file d'attente
// soumise a un echantillonnage et une validation humaine avant toute
// reintegration eventuelle au pipeline d'entrainement.
func IngestUserFeedback(feedback []Example) []Example {
	var approved []Example
	sample := sampleForHumanReview(feedback)
	for _, ex := range sample {
		if humanApproves(ex) {
			approved = append(approved, ex)
		} else {
			recordAudit("feedback_rejected_human_review", ex.Prompt)
		}
	}
	return approved
}

// ValidateModelAfterTraining compare les performances du modele fraichement
// entraine a un jeu de reference fixe pour detecter une derive suspecte.
func ValidateModelAfterTraining(evaluate func(goldenSet []Example) float64, goldenSet []Example) error {
	score := evaluate(goldenSet)
	if score < acceptableDriftThreshold {
		return errModelDrift
	}
	return nil
}

func verifyIntegrity(src DataSource) error {
	resp, err := http.Get(src.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	sum := sha256.Sum256(body)
	if hex.EncodeToString(sum[:]) != src.SignedHash {
		return fmt.Errorf("hash invalide pour la source %s", src.ID)
	}
	return nil
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
	return []Example{}, nil
}

func filterStatisticalOutliers(batch []Example) []Example {
	// Detection d'anomalies statistiques (motifs repetitifs suspects,
	// valeurs aberrantes) avant integration au dataset (details omis).
	return batch
}

func sampleForHumanReview(feedback []Example) []Example { return feedback }
func humanApproves(ex Example) bool                     { return false }
func recordAudit(event, detail string)                  {}
