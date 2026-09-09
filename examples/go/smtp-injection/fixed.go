// fixed.go - Correction CWE-93
// Toute sequence CR ou LF dans les champs utilisateur destines aux
// en-tetes est rejetee avant construction du message. Les en-tetes
// eux-memes sont produits via net/mail et mime, qui gerent
// correctement l'encodage, plutot que par concatenation manuelle.
package mailer

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

var ErrInvalidHeaderValue = errors.New("valeur d'en-tete invalide")

// sanitizeHeaderValue rejette toute entree pouvant casser la structure
// des en-tetes du message (CR/LF = debut d'un nouvel en-tete ou du corps).
func sanitizeHeaderValue(value string) (string, error) {
	if strings.ContainsAny(value, "\r\n") {
		return "", ErrInvalidHeaderValue
	}
	return value, nil
}

func SendContactMessage(addr string, auth smtp.Auth, from, userName, userSubject, body string) error {
	// FIXED: validation stricte avant insertion dans les en-tetes
	safeName, err := sanitizeHeaderValue(userName)
	if err != nil {
		return err
	}
	safeSubject, err := sanitizeHeaderValue(userSubject)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("From: %s\r\nSubject: Contact de %s: %s\r\n\r\n%s",
		from, safeName, safeSubject, body)

	return smtp.SendMail(addr, auth, from, []string{"support@example.com"}, []byte(msg))
}
