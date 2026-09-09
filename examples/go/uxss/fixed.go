// fixed.go - Correction CWE-79
// On ajoute une Content Security Policy stricte via en-tete HTTP, une
// verification d'integrite (Subresource Integrity) sur le script tiers,
// et un attribut sandbox minimal sur l'iframe tierce. Ces mesures
// n'eliminent pas une compromission du tiers lui-meme, mais reduisent
// fortement la surface d'exposition de l'application si cela se produit.
package page

import (
	"fmt"
	"net/http"
)

func DashboardPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// FIXED: CSP stricte en defense en profondeur
	w.Header().Set("Content-Security-Policy",
		"default-src 'self'; script-src 'self' https://cdn.example.com; frame-src https://widget-tiers.example.com")

	// FIXED: integrity + crossorigin sur le script tiers, sandbox minimal sur l'iframe
	fmt.Fprint(w, `
<html>
  <head>
    <script src="https://cdn.example.com/widget.js"
            integrity="sha384-REPLACE-WITH-VERIFIED-HASH"
            crossorigin="anonymous"></script>
  </head>
  <body>
    <iframe src="https://widget-tiers.example.com/chat"
            sandbox="allow-scripts allow-same-origin"
            referrerpolicy="no-referrer"></iframe>
  </body>
</html>`)
}
