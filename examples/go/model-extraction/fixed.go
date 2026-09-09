// fixed.go - Correctif CWE-200
// L'endpoint applique un quota strict par cle API, ne renvoie plus les
// logits bruts sans besoin metier, et surveille les patterns de requetes
// pour detecter une extraction systematique du modele.
package api

import (
	"encoding/json"
	"net/http"
)

const apiKey = "REPLACE_WITH_YOUR_API_KEY"

// requestsPerHourLimit est le quota applique par cle API sur l'endpoint
// d'inference.
const requestsPerHourLimit = 100

type InferRequest struct {
	Input string `json:"input"`
}

type InferResult struct {
	Text string `json:"text"`
	// Les logits ne sont plus exposes par defaut : ils ne sont calcules et
	// renvoyes que si un cas d'usage metier explicite l'exige, via un canal
	// distinct et audite.
}

// RateLimiter applique un quota strict par cle API sur les endpoints
// d'inference.
type RateLimiter interface {
	Allow(apiKey string, limitPerHour int) bool
}

// UsageMonitor detecte les patterns de requetes systematiques evoquant une
// tentative d'extraction de modele (volume, regularite anormale).
type UsageMonitor interface {
	DetectsExtractionPattern(apiKey string) bool
}

type AuditLogger interface {
	Record(event, apiKey string)
}

type InferService struct {
	RateLimit RateLimiter
	Usage     UsageMonitor
	Audit     AuditLogger
}

// InferHandler applique le quota, detecte les patterns d'usage anormaux,
// et ne renvoie que la sortie textuelle necessaire au cas d'usage.
func (s *InferService) InferHandler(w http.ResponseWriter, r *http.Request) {
	clientAPIKey := r.Header.Get("X-API-Key")
	if clientAPIKey == "" {
		http.Error(w, "cle API requise", http.StatusUnauthorized)
		return
	}

	if !s.RateLimit.Allow(clientAPIKey, requestsPerHourLimit) {
		s.Audit.Record("rate_limit_exceeded", clientAPIKey)
		http.Error(w, "quota depasse", http.StatusTooManyRequests)
		return
	}

	if s.Usage.DetectsExtractionPattern(clientAPIKey) {
		s.Audit.Record("suspected_model_extraction", clientAPIKey)
		http.Error(w, "usage anormal detecte", http.StatusTooManyRequests)
		return
	}

	var req InferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "requete invalide", http.StatusBadRequest)
		return
	}

	// returnLogits=false : aucune information interne du modele exposee
	// au-dela de la sortie textuelle necessaire au cas d'usage.
	result := predict(req.Input, false)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"output": result.Text,
	})
}

func predict(input string, returnLogits bool) InferResult {
	return InferResult{Text: ""}
}
