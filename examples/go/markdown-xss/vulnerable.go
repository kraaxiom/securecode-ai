// vulnerable.go - CWE-79: XSS via rendu Markdown
// Le contenu Markdown soumis par l'utilisateur est converti en HTML
// avec l'option "Unsafe" activee, qui autorise le HTML brut a passer
// tel quel dans la sortie. Un utilisateur peut ainsi inserer directement
// une balise <script> ou un gestionnaire d'evenement dans son Markdown.
package blog

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

var mdVulnerable = goldmark.New(
	goldmark.WithRendererOptions(
		html.WithUnsafe(), // VULNERABLE: autorise le HTML/script brut dans la sortie
	),
)

func RenderMarkdown(source []byte) (string, error) {
	var buf bytes.Buffer
	if err := mdVulnerable.Convert(source, &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
