// fixed.go - Correction CWE-97
// On echappe le HTML (protection XSS classique) puis on neutralise
// explicitement la sequence de directive SSI "<!--#" residuelle. La
// vraie correction structurelle est de ne jamais ecrire de contenu
// utilisateur dans un fichier interprete par SSI et d'utiliser un
// moteur de templates applicatif a la place ; l'echappement reste ici
// en defense en profondeur.
package comments

import (
	"fmt"
	"html"
	"os"
	"strings"
)

func sanitizeSSI(input string) string {
	escaped := html.EscapeString(input)
	// FIXED: neutralise toute sequence de directive SSI residuelle
	return strings.ReplaceAll(escaped, "<!--#", "&lt;!--#")
}

func AppendComment(path, author, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	safeAuthor := sanitizeSSI(author)
	safeText := sanitizeSSI(text)

	// FIXED: en production, preferer un moteur de templates applicatif
	// (html/template) et desactiver SSI sur le repertoire de contenu utilisateur
	_, err = fmt.Fprintf(f, "<p><strong>%s</strong>: %s</p>\n", safeAuthor, safeText)
	return err
}
