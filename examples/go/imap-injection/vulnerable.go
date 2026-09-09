// vulnerable.go - CWE-93: Injection de commande IMAP (CRLF)
// Le terme de recherche fourni par l'utilisateur est insere directement
// dans une commande IMAP brute envoyee sur le socket. Une entree
// contenant CRLF permet d'injecter une commande IMAP supplementaire
// (ex: se deplacer vers un autre dossier, executer une commande admin).
package mailsearch

import (
	"fmt"
	"net/textproto"
)

func SearchSubject(conn *textproto.Conn, tag string, subject string) error {
	// VULNERABLE: subject est concatene tel quel dans la commande IMAP
	cmd := fmt.Sprintf("%s SEARCH SUBJECT \"%s\"", tag, subject)
	_, err := conn.Cmd(cmd)
	return err
}
