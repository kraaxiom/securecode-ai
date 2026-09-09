// fixed.go - Correction CWE-611
// On abandonne l'appel a un outil externe qui resout les entites, et
// on utilise encoding/xml de la bibliotheque standard Go, qui ne
// resout jamais les entites externes ni les DTD par defaut : il n'y a
// aucune option a desactiver, le comportement sur est celui par defaut.
package importer

import (
	"bytes"
	"encoding/xml"
)

type Document struct {
	XMLName xml.Name `xml:"document"`
	Title   string   `xml:"title"`
	Content string   `xml:"content"`
}

func NormalizeUserXML(userXML []byte) (*Document, error) {
	// FIXED: encoding/xml ne resout ni les entites externes ni les DTD
	decoder := xml.NewDecoder(bytes.NewReader(userXML))

	// Defense en profondeur explicite : n'autoriser aucune entite personnalisee
	decoder.Entity = map[string]string{}
	decoder.Strict = true

	var doc Document
	if err := decoder.Decode(&doc); err != nil {
		return nil, err
	}
	return &doc, nil
}
