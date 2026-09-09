// fixed.go - Correction CWE-79
// Le parametre de recherche est insere via html/template, dont
// l'echappement contextuel automatique neutralise tout balisage ou
// script avant de l'ecrire dans la reponse HTML.
package search

import (
	"html/template"
	"net/http"
)

var searchTmpl = template.Must(template.New("search").Parse(
	`<html><body><p>Resultats pour : {{.Query}}</p></body></html>`,
))

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// FIXED: html/template echappe automatiquement selon le contexte HTML
	searchTmpl.Execute(w, struct{ Query string }{Query: q})
}
