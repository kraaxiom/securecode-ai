// fixed.go - Correction CWE-79
// L'option Unsafe est retiree : goldmark echappe alors par defaut tout
// HTML brut present dans la source Markdown. En defense en profondeur
// supplementaire (pour les cas ou du HTML doit malgre tout etre
// autorise), on passe la sortie dans bluemonday avec une politique
// d'allowlist stricte avant de la servir.
package blog

import (
	"bytes"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
)

// FIXED: pas de html.WithUnsafe() -> le HTML brut de la source est echappe
var mdSafe = goldmark.New()

var htmlSanitizer = bluemonday.UGCPolicy() // allowlist stricte (balises de contenu utilisateur)

func RenderMarkdown(source []byte) (string, error) {
	var buf bytes.Buffer
	if err := mdSafe.Convert(source, &buf); err != nil {
		return "", err
	}
	// FIXED: sanitisation supplementaire en defense en profondeur
	return htmlSanitizer.Sanitize(buf.String()), nil
}
