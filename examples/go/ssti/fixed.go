// fixed.go - Correction CWE-1336
// Le template est une constante statique et l'entree utilisateur
// n'est jamais qu'une donnee injectee dans un champ, pas du code
// de template.
package handlers

import (
	"net/http"
	"text/template"
)

var greetTmpl = template.Must(template.New("greet").Parse("Hello, {{.Name}}!"))

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	// FIXED: le nom est une donnee, jamais interpretee comme du template
	greetTmpl.Execute(w, map[string]string{"Name": name})
}
