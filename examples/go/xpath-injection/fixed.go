// fixed.go - Correction CWE-643
// La bibliotheque XPath utilisee ne supporte pas nativement les
// variables liees : on echappe donc strictement chaque valeur en la
// convertissant en litteral XPath sur, via concat() lorsque la valeur
// contient elle-meme une apostrophe. Recommandation structurelle :
// ne jamais utiliser XPath comme mecanisme d'authentification, mais
// une base de donnees avec hash de mot de passe compare cote applicatif.
package auth

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

// xpathLiteral construit un litteral XPath 1.0 sur, meme si la valeur
// contient des apostrophes (impossible a echapper directement en XPath 1.0).
func xpathLiteral(value string) string {
	if !strings.Contains(value, "'") {
		return "'" + value + "'"
	}
	parts := strings.Split(value, "'")
	quoted := make([]string, 0, len(parts)*2)
	for i, p := range parts {
		if i > 0 {
			quoted = append(quoted, `"'"`)
		}
		quoted = append(quoted, "'"+p+"'")
	}
	return "concat(" + strings.Join(quoted, ", ") + ")"
}

func FindUserNode(doc *xmlquery.Node, username, password string) *xmlquery.Node {
	// FIXED: chaque valeur est convertie en litteral XPath sur avant insertion
	expr := fmt.Sprintf("//user[username=%s and password=%s]",
		xpathLiteral(username), xpathLiteral(password))
	return xmlquery.FindOne(doc, expr)
}
