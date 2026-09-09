// vulnerable.go - CWE-611: XML External Entity (XXE) Injection
// Le document XML fourni par l'utilisateur est valide/normalise en
// invoquant l'outil externe xmllint avec l'option --noent, qui resout
// et substitue les entites externes declarees dans une DTD. Un document
// contenant <!ENTITY xxe SYSTEM "file:///etc/passwd"> permet donc de
// lire des fichiers arbitraires du serveur ou d'atteindre des services
// internes (SSRF).
package importer

import (
	"bytes"
	"os/exec"
)

func NormalizeUserXML(userXML []byte) ([]byte, error) {
	// VULNERABLE: --noent resout les entites externes de la DTD fournie
	cmd := exec.Command("xmllint", "--noent", "-")
	cmd.Stdin = bytes.NewReader(userXML)
	return cmd.Output()
}
