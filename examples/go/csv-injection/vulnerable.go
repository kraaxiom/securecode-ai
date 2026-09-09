// vulnerable.go - CWE-1236: Injection de formule CSV
// Les champs fournis par l'utilisateur sont ecrits tels quels dans un
// export CSV. Si un champ commence par =, +, - ou @, le tableur qui
// ouvrira le fichier (Excel, LibreOffice, Google Sheets) l'interpretera
// comme une formule, ce qui peut declencher DDE/macro ou exfiltrer des
// donnees vers un serveur externe.
package export

import (
	"encoding/csv"
	"io"
)

type Contact struct {
	Name  string
	Email string
	Note  string
}

func WriteContactsCSV(w io.Writer, contacts []Contact) error {
	writer := csv.NewWriter(w)
	defer writer.Flush()

	writer.Write([]string{"Nom", "Email", "Note"})
	for _, c := range contacts {
		// VULNERABLE: aucune neutralisation des caracteres declencheurs de formule
		writer.Write([]string{c.Name, c.Email, c.Note})
	}
	return nil
}
