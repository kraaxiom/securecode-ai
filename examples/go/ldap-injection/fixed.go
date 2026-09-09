// fixed.go - Correction CWE-90
// La bibliotheque go-ldap fournit ldap.EscapeFilter, qui echappe les
// caracteres speciaux du langage de filtre LDAP (RFC 4515 : * ( ) \ et
// NUL). L'entree utilisateur passe systematiquement par cette fonction
// avant d'etre inseree dans le filtre.
package auth

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

func FindUser(conn *ldap.Conn, baseDN, uid string) (*ldap.SearchResult, error) {
	// FIXED: echappement dedie des caracteres speciaux du filtre LDAP
	safeUID := ldap.EscapeFilter(uid)
	filter := fmt.Sprintf("(uid=%s)", safeUID)

	req := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"dn", "cn", "mail"},
		nil,
	)
	return conn.Search(req)
}
