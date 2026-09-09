// vulnerable.go - CWE-917: Expression Language Injection
// Une regle metier est construite en concatenant l'entree utilisateur
// dans une expression evaluee dynamiquement par un moteur de regles.
package handlers

import (
	"fmt"
	"net/http"

	"github.com/Knetic/govaluate"
)

func DiscountRuleHandler(w http.ResponseWriter, r *http.Request) {
	condition := r.URL.Query().Get("condition") // ex: "amount > 100"

	// VULNERABLE: expression utilisateur evaluee directement
	expr := "amount > 0 && (" + condition + ")"
	e, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		http.Error(w, "invalid rule", http.StatusBadRequest)
		return
	}
	result, _ := e.Evaluate(map[string]interface{}{"amount": 150})
	fmt.Fprintf(w, "eligible=%v", result)
}
