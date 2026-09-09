// fixed.go - Correction CWE-94
// Aucune evaluation dynamique de code utilisateur : les operations
// autorisees sont limitees a une allowlist explicite traitee cote
// application (pas d'interpreteur generique expose).
package handlers

import (
	"encoding/json"
	"net/http"
)

type calcRequest struct {
	Op string  `json:"op"` // doit etre "add", "sub", "mul", "div"
	A  float64 `json:"a"`
	B  float64 `json:"b"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	var req calcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// FIXED: logique explicite, aucune evaluation de code arbitraire
	var result float64
	switch req.Op {
	case "add":
		result = req.A + req.B
	case "sub":
		result = req.A - req.B
	case "mul":
		result = req.A * req.B
	case "div":
		if req.B == 0 {
			http.Error(w, "division by zero", http.StatusBadRequest)
			return
		}
		result = req.A / req.B
	default:
		http.Error(w, "unsupported operation", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]float64{"result": result})
}
