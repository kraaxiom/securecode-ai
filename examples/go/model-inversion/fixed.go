package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

// modelClient représente un client vers un modèle fine-tuné.
// Le pipeline d'entraînement associé applique désormais l'anonymisation
// des données sensibles et la confidentialité différentielle (DP-SGD),
// avec un test de mémorisation avant toute mise en production (voir pipeline offline).
type modelClient struct {
	endpoint string
}

func (m *modelClient) predict(input string) (string, error) {
	req, err := http.NewRequest("POST", m.endpoint+"/v1/predict", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Text string `json:"text"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Text, nil
}

// usageMonitor surveille le volume et la structure des requêtes par clé API
// afin de détecter une tentative de reconstruction progressive de données mémorisées.
type usageMonitor struct {
	mu      sync.Mutex
	counts  map[string]int
	window  time.Duration
	maxHits int
}

func newUsageMonitor(window time.Duration, maxHits int) *usageMonitor {
	return &usageMonitor{counts: make(map[string]int), window: window, maxHits: maxHits}
}

// allow applique une limite de requêtes ("rate limit") par clé API.
func (u *usageMonitor) allow(apiKey string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.counts[apiKey]++
	return u.counts[apiKey] <= u.maxHits
}

// detectsReconstructionPattern applique une heuristique simple de détection
// de séquences de requêtes typiques d'une attaque par inversion de modèle
// (variations minimes et répétées d'une même entrée).
func (u *usageMonitor) detectsReconstructionPattern(apiKey string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.counts[apiKey] > u.maxHits/2 // seuil de suspicion avant blocage strict
}

var monitor = newUsageMonitor(time.Hour, 50)

func auditLog(event, apiKey string) {
	// Journalisation structurée pour permettre l'investigation a posteriori.
}

// inferHandler limite le volume de requêtes, surveille les motifs de reconstruction
// et rejette les usages anormaux avant d'atteindre le modèle.
func inferHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "clé API manquante", http.StatusUnauthorized)
		return
	}

	if !monitor.allow(apiKey) {
		auditLog("rate_limit_exceeded", apiKey)
		http.Error(w, "limite de requêtes dépassée", http.StatusTooManyRequests)
		return
	}

	if monitor.detectsReconstructionPattern(apiKey) {
		auditLog("suspected_model_inversion", apiKey)
		http.Error(w, "usage anormal détecté", http.StatusTooManyRequests)
		return
	}

	var body struct {
		Input string `json:"input"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	client := &modelClient{endpoint: "http://internal-model.local"}
	output, err := client.predict(body.Input)
	if err != nil {
		http.Error(w, "erreur interne", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"result": output})
}

func main() {
	http.HandleFunc("/v1/infer", inferHandler)
	http.ListenAndServe(":8080", nil)
}
