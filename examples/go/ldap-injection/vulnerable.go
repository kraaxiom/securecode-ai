// vulnerable.go - CWE-90: Injection LDAP
// Le filtre de recherche LDAP est construit par concatenation directe
// de l'entree utilisateur. Une valeur comme "*)(uid=*))(|(uid=*" permet
// de modifier la logique du filtre et de contourner l'authentification
// ou d'enumerer des entrees non autorisees.
package auth

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

func FindUser(conn *ldap.Conn, baseDN, uid string) (*ldap.SearchResult, error) {
	// VULNERABLE: uid est injecte tel quel dans le filtre LDAP
	filter := fmt.Sprintf("(uid=%s)", uid)

	req := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"dn", "cn", "mail"},
		nil,
	)
	return conn.Search(req)
}
