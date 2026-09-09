// fixed.go - Correction CWE-917
// Les regles sont selectionnees parmi un ensemble predefini (allowlist),
// aucune expression utilisateur n'est evaluee dynamiquement.
package handlers

import (
	"fmt"
	"net/http"
)

var allowedRules = map[string]func(amount float64) bool{
	"over100": func(amount float64) bool { return amount > 100 },
	"over500": func(amount float64) bool { return amount > 500 },
}

func DiscountRuleHandler(w http.ResponseWriter, r *http.Request) {
	ruleName := r.URL.Query().Get("rule")

	// FIXED: selection dans une allowlist, pas d'evaluation dynamique
	rule, ok := allowedRules[ruleName]
	if !ok {
		http.Error(w, "unknown rule", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "eligible=%v", rule(150))
}
