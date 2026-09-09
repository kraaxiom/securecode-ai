// vulnerable.go - CWE-79: Stored Cross-Site Scripting
// Le commentaire soumis par un utilisateur est stocke tel quel, puis
// reaffiche sans echappement a tous les visiteurs de la page. Le
// script injecte une fois s'execute pour chaque visiteur qui consulte
// la page, contrairement au XSS reflechi qui necessite un lien piege.
package comments

import (
	"fmt"
	"net/http"
)

type Comment struct {
	Author string
	Body   string
}

func RenderComments(w http.ResponseWriter, comments []Comment) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<ul>")
	for _, c := range comments {
		// VULNERABLE: c.Body provient du stockage et n'est jamais echappe
		fmt.Fprintf(w, "<li><strong>%s</strong>: %s</li>", c.Author, c.Body)
	}
	fmt.Fprint(w, "</ul>")
}
