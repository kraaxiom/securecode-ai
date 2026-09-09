// fixed.go - Correction CWE-91
// Le document est produit via encoding/xml.Marshal sur une struct
// typee : le marshaller applique automatiquement l'echappement des
// caracteres speciaux XML (<, >, &, ', ") dans le contenu des elements,
// ce qui empeche toute donnee utilisateur d'etre interpretee comme du
// balisage.
package profile

import (
	"encoding/xml"
)

type User struct {
	XMLName xml.Name `xml:"user"`
	Name    string   `xml:"name"`
	Bio     string   `xml:"bio"`
}

func BuildUserXML(username, bio string) (string, error) {
	// FIXED: encoding/xml echappe automatiquement le contenu des elements
	u := User{Name: username, Bio: bio}
	out, err := xml.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
