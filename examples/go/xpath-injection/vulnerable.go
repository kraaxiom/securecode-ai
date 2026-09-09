// vulnerable.go - CWE-643: Injection XPath
// L'expression XPath utilisee pour authentifier un utilisateur est
// construite par concatenation directe. Une entree comme
// "' or '1'='1" dans le mot de passe permet de contourner la
// verification et de s'authentifier sans connaitre le mot de passe.
package auth

import (
	"fmt"

	"github.com/antchfx/xmlquery"
)

func FindUserNode(doc *xmlquery.Node, username, password string) *xmlquery.Node {
	// VULNERABLE: username et password sont inseres tels quels dans l'expression XPath
	expr := fmt.Sprintf("//user[username='%s' and password='%s']", username, password)
	return xmlquery.FindOne(doc, expr)
}
