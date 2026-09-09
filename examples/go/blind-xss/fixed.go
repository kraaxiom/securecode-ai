// fixed.go - Correction CWE-79
// On remplace la construction manuelle de HTML par html/template, dont
// l'echappement contextuel automatique neutralise tout balisage ou
// script present dans le message avant de l'inserer dans la page
// consultee par l'administrateur.
package support

import (
	"html/template"
	"net/http"
)

type Ticket struct {
	ID      string
	Message string
}

var ticketTmpl = template.Must(template.New("ticket").Parse(
	`<div class="ticket"><h3>Ticket {{.ID}}</h3><p>{{.Message}}</p></div>`,
))

func RenderAdminTicket(w http.ResponseWriter, t Ticket) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// FIXED: html/template echappe automatiquement le contenu selon le contexte HTML
	return ticketTmpl.Execute(w, t)
}
