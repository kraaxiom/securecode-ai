// vulnerable.go - CWE-1336: Server-Side Template Injection
// Le contenu du template est construit a partir de l'entree utilisateur
// puis parse/execute, permettant l'injection de directives de template.
package handlers

import (
	"net/http"
	"text/template"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	// VULNERABLE: l'entree utilisateur devient elle-meme le template
	tmplStr := "Hello, " + name + "!"
	tmpl, err := template.New("greet").Parse(tmplStr)
	if err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
