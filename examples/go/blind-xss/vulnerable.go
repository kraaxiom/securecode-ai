// vulnerable.go - CWE-79: Blind Cross-Site Scripting
// Le message de support soumis par un utilisateur externe est stocke
// puis affiche tel quel, sans echappement, dans le tableau de bord
// interne consulte par un administrateur. L'attaquant ne voit jamais
// le resultat directement ("blind") : le script injecte s'execute dans
// le contexte de session de l'administrateur qui consulte le ticket.
package support

import (
	"fmt"
	"net/http"
)

type Ticket struct {
	ID      string
	Message string
}

func RenderAdminTicket(w http.ResponseWriter, t Ticket) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// VULNERABLE: t.Message est injecte tel quel dans le HTML du dashboard admin
	fmt.Fprintf(w, "<div class=\"ticket\"><h3>Ticket %s</h3><p>%s</p></div>", t.ID, t.Message)
}
