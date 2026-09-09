// vulnerable.go - CWE-97: Injection Server-Side Includes (SSI)
// Le commentaire utilisateur est ecrit tel quel dans un fichier .shtml
// servi par un serveur web avec SSI active (ex: Apache mod_include).
// Une valeur contenant "<!--#exec cmd=\"...\" -->" sera interpretee et
// executee cote serveur a chaque affichage de la page.
package comments

import (
	"fmt"
	"os"
)

func AppendComment(path, author, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// VULNERABLE: text peut contenir une directive SSI active
	_, err = fmt.Fprintf(f, "<p><strong>%s</strong>: %s</p>\n", author, text)
	return err
}
