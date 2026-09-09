// vulnerable.go - CWE-94: Code Injection
// Une expression fournie par l'utilisateur est evaluee dynamiquement
// par un moteur d'expression, permettant l'execution de code arbitraire.
package handlers

import (
	"fmt"
	"net/http"

	"github.com/Knetic/govaluate"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	expr := r.URL.Query().Get("expr")

	// VULNERABLE: evaluation directe d'une expression utilisateur non filtree
	e, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		http.Error(w, "invalid expression", http.StatusBadRequest)
		return
	}
	result, _ := e.Evaluate(nil)
	fmt.Fprintf(w, "result=%v", result)
}
