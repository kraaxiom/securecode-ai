// fixed.go - Correction CWE-79
// Le rendu passe par html/template : chaque commentaire est echappe
// selon le contexte HTML au moment de l'affichage, quel que soit ce qui
// a ete stocke en base. L'echappement a l'affichage (plutot qu'a
// l'ecriture) evite aussi de corrompre les donnees stockees.
package comments

import (
	"html/template"
	"net/http"
)

type Comment struct {
	Author string
	Body   string
}

var commentsTmpl = template.Must(template.New("comments").Parse(`<ul>
{{range .}}<li><strong>{{.Author}}</strong>: {{.Body}}</li>{{end}}
</ul>`))

func RenderComments(w http.ResponseWriter, comments []Comment) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// FIXED: echappement automatique a l'affichage, quel que soit le contenu stocke
	return commentsTmpl.Execute(w, comments)
}
