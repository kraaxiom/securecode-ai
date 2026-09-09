// vulnerable.go - CWE-79: XSS via fichier SVG
// Le fichier SVG uploade par l'utilisateur est stocke puis servi tel
// quel avec le type MIME image/svg+xml en affichage inline. Un SVG peut
// contenir <script>, des gestionnaires d'evenements (onload) ou des
// liens javascript:, executes par le navigateur car SVG est un format
// XML qui autorise du contenu HTML actif.
package avatar

import (
	"io"
	"net/http"
	"os"
)

func ServeAvatarSVG(w http.ResponseWriter, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// VULNERABLE: contenu SVG non sanitise, servi en affichage inline
	w.Header().Set("Content-Type", "image/svg+xml")
	_, err = io.Copy(w, f)
	return err
}
