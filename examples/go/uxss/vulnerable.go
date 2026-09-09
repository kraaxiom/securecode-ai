// vulnerable.go - CWE-79: Universal Cross-Site Scripting (UXSS)
// Le serveur sert une page qui charge un script tiers depuis un CDN
// sans verification d'integrite (Subresource Integrity) et embarque un
// widget tiers dans une iframe sans attribut sandbox. Si le CDN ou le
// widget tiers est compromis, le code injecte s'execute avec les
// pleins privileges de l'origine de l'application (cookies, DOM,
// stockage local), sans qu'aucune faille ne soit presente dans le code
// applicatif lui-meme.
package page

import (
	"fmt"
	"net/http"
)

func DashboardPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// VULNERABLE: script tiers sans integrity, iframe tierce sans sandbox
	fmt.Fprint(w, `
<html>
  <head>
    <script src="https://cdn.example.com/widget.js"></script>
  </head>
  <body>
    <iframe src="https://widget-tiers.example.com/chat"></iframe>
  </body>
</html>`)
}
