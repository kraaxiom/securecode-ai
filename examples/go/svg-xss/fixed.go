// fixed.go - Correction CWE-79
// Le SVG est sanitise avant stockage avec une politique HTML/SVG
// dediee de bluemonday (suppression de <script>, des gestionnaires
// d'evenements et des schemas javascript:), et servi avec
// Content-Disposition: attachment pour forcer le telechargement plutot
// que l'affichage inline dans le contexte de l'origine de l'application.
package avatar

import (
	"io"
	"net/http"
	"os"

	"github.com/microcosm-cc/bluemonday"
)

// SanitizeSVG neutralise le contenu actif avant stockage.
func SanitizeSVG(raw []byte) []byte {
	// FIXED: politique dediee retirant script/gestionnaires d'evenements/javascript:
	policy := bluemonday.NewPolicy()
	policy.AllowElements("svg", "path", "circle", "rect", "g", "polygon", "line")
	policy.AllowAttrs("d", "cx", "cy", "r", "width", "height", "viewBox", "fill", "stroke").Globally()
	return policy.SanitizeBytes(raw)
}

func ServeAvatarSVG(w http.ResponseWriter, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// FIXED: telechargement force plutot qu'affichage inline dans l'origine de l'app
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Content-Disposition", "attachment; filename=\"avatar.svg\"")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	_, err = io.Copy(w, f) // le fichier a deja ete sanitise via SanitizeSVG avant stockage
	return err
}
