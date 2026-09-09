// vulnerable.go - CWE-79: DOM-based Cross-Site Scripting
// Le serveur sert une page contenant un script inline qui lit
// directement location.search cote client et l'insere via innerHTML.
// La donnee ne transite jamais par le serveur (le fragment de requete
// peut meme rester local au navigateur), mais la page servie par ce
// handler contient le sink vulnerable qui permet l'execution du script
// injecte par l'URL.
package widget

import (
	"fmt"
	"net/http"
)

func WidgetPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// VULNERABLE: le script inline utilise innerHTML avec une source non fiable
	fmt.Fprint(w, `
<html>
  <body>
    <div id="result"></div>
    <script>
      var params = new URLSearchParams(window.location.search);
      var q = params.get('q');
      document.getElementById('result').innerHTML = 'Recherche : ' + q;
    </script>
  </body>
</html>`)
}
