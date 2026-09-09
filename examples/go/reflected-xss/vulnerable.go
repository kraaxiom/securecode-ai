// vulnerable.go - CWE-79: Reflected Cross-Site Scripting
// Le parametre de recherche est reinjecte tel quel dans la reponse
// HTML via fmt.Fprintf, sans aucun echappement. Un lien contenant
// ?q=<script>...</script> execute le script dans le navigateur de la
// victime des qu'elle clique dessus.
package search

import (
	"fmt"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// VULNERABLE: q est reflete tel quel dans le HTML de la reponse
	fmt.Fprintf(w, "<html><body><p>Resultats pour : %s</p></body></html>", q)
}
