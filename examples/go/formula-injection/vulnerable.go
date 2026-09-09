// vulnerable.go - CWE-1236: Injection de formule dans un rapport tableur
// Le montant de depense saisi par un utilisateur est ecrit directement
// dans une cellule d'un rapport genere pour un tableur. Une valeur comme
// "=HYPERLINK(\"http://attacker.example\",\"cliquez\")" sera executee comme
// formule des l'ouverture du fichier par la victime (comptable, RH...).
package report

import (
	"fmt"
	"io"
)

type ExpenseLine struct {
	Employee string
	Amount   string // saisi librement par l'utilisateur
}

func WriteExpenseReport(w io.Writer, lines []ExpenseLine) {
	fmt.Fprintln(w, "Employe;Montant")
	for _, l := range lines {
		// VULNERABLE: le montant est inseré tel quel, meme s'il commence par = + - @
		fmt.Fprintf(w, "%s;%s\n", l.Employee, l.Amount)
	}
}
