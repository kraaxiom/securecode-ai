// vulnerable.go - CWE-91: Injection XML (XML Injection)
// Le document XML est construit par concatenation de chaines a partir
// de l'entree utilisateur. Une valeur contenant des balises XML
// (ex: "</user><user><role>admin</role></user>") modifie la structure
// du document et peut ajouter des elements ou attributs non prevus.
package profile

import (
	"fmt"
)

func BuildUserXML(username, bio string) string {
	// VULNERABLE: username et bio sont inseres tels quels dans le XML
	return fmt.Sprintf(`<user><name>%s</name><bio>%s</bio></user>`, username, bio)
}
