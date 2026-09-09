// fixed.go - Correction CWE-1236
// Chaque champ est verifie : s'il commence par un caractere pouvant
// declencher l'interpretation d'une formule par un tableur (= + - @, ou
// tabulation/CR), on prefixe le champ avec une apostrophe pour forcer
// une interpretation en tant que texte pur.
package export

import (
	"encoding/csv"
	"io"
	"strings"
)

type Contact struct {
	Name  string
	Email string
	Note  string
}

// sanitizeCSVField neutralise les caracteres declencheurs de formule.
func sanitizeCSVField(field string) string {
	if field == "" {
		return field
	}
	switch field[0] {
	case '=', '+', '-', '@', '\t', '\r':
		return "'" + field
	default:
		return field
	}
}

func WriteContactsCSV(w io.Writer, contacts []Contact) error {
	writer := csv.NewWriter(w)
	defer writer.Flush()

	writer.Write([]string{"Nom", "Email", "Note"})
	for _, c := range contacts {
		// FIXED: chaque champ passe par la sanitisation avant ecriture
		row := []string{
			sanitizeCSVField(strings.TrimSpace(c.Name)),
			sanitizeCSVField(strings.TrimSpace(c.Email)),
			sanitizeCSVField(strings.TrimSpace(c.Note)),
		}
		writer.Write(row)
	}
	return nil
}
