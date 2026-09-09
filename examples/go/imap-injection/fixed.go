// fixed.go - Correction CWE-93
// On rejette toute entree contenant CR ou LF, puis on transmet la
// valeur comme "IMAP literal" ({N}\r\n<donnees>) plutot que comme chaine
// entre guillemets construite par concatenation : le serveur IMAP lit
// alors exactement N octets, sans jamais interpreter leur contenu comme
// une nouvelle commande.
package mailsearch

import (
	"errors"
	"fmt"
	"net/textproto"
	"strings"
)

var ErrInvalidSearchTerm = errors.New("terme de recherche invalide")

func SearchSubject(conn *textproto.Conn, tag string, subject string) error {
	// FIXED: refus explicite de toute sequence CRLF dans l'entree
	if strings.ContainsAny(subject, "\r\n") {
		return ErrInvalidSearchTerm
	}

	// FIXED: transmission via literal IMAP {N} pour eviter toute
	// interpretation de metacaracteres par le parseur de commandes
	id, err := conn.Cmd("%s SEARCH SUBJECT {%d}", tag, len(subject))
	if err != nil {
		return err
	}
	conn.StartResponse(id)
	defer conn.EndResponse(id)

	if _, err := fmt.Fprint(conn.W, subject+"\r\n"); err != nil {
		return err
	}
	return conn.W.Flush()
}
