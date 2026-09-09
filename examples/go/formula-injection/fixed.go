// fixed.go - Correction CWE-1236
// On valide d'abord que le montant est numerique (c'est la seule donnee
// legitime attendue dans ce champ) ; a defaut, on neutralise tout
// caractere declencheur de formule avant ecriture, en defense en
// profondeur pour les autres champs texte libres.
package report

import (
	"fmt"
	"io"
	"regexp"
)

type ExpenseLine struct {
	Employee string
	Amount   string
}

var numericAmount = regexp.MustCompile(`^-?[0-9]+(\.[0-9]{1,2})?$`)

// sanitizeFormulaField neutralise = + - @ en debut de champ.
func sanitizeFormulaField(field string) string {
	if field == "" {
		return field
	}
	switch field[0] {
	case '=', '+', '-', '@':
		return "'" + field
	default:
		return field
	}
}

func WriteExpenseReport(w io.Writer, lines []ExpenseLine) error {
	fmt.Fprintln(w, "Employe;Montant")
	for _, l := range lines {
		// FIXED: le montant doit respecter un format numerique strict
		if !numericAmount.MatchString(l.Amount) {
			return fmt.Errorf("montant invalide pour %s", l.Employee)
		}
		employee := sanitizeFormulaField(l.Employee)
		fmt.Fprintf(w, "%s;%s\n", employee, l.Amount)
	}
	return nil
}
