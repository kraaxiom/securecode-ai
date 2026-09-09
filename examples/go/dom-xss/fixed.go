// fixed.go - Correction CWE-79
// Le script inline sert servi par ce handler utilise textContent au
// lieu de innerHTML pour inserer la valeur issue de location.search :
// la donnee reste toujours traitee comme du texte, jamais interpretee
// comme du HTML actif, quelle que soit la source (non fiable par
// definition car controlee par l'utilisateur via l'URL).
package widget

import (
	"fmt"
	"net/http"
)

func WidgetPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Security-Policy", "script-src 'self'")
	// FIXED: textContent au lieu de innerHTML -> jamais interprete comme HTML
	fmt.Fprint(w, `
<html>
  <body>
    <div id="result"></div>
    <script>
      var params = new URLSearchParams(window.location.search);
      var q = params.get('q');
      document.getElementById('result').textContent = 'Recherche : ' + q;
    </script>
  </body>
</html>`)
}
