// vulnerable.go - CWE-93: Injection d'en-tetes SMTP
// Le sujet et le nom fournis par l'utilisateur sont concatenes
// directement dans les en-tetes du message brut. Une valeur contenant
// CRLF permet d'injecter des en-tetes SMTP supplementaires (Bcc, Cc)
// ou de terminer prematurement les en-tetes pour injecter un corps de
// message arbitraire (spam relay, phishing).
package mailer

import (
	"fmt"
	"net/smtp"
)

func SendContactMessage(addr string, auth smtp.Auth, from, userName, userSubject, body string) error {
	// VULNERABLE: userName et userSubject sont inseres sans validation
	msg := fmt.Sprintf("From: %s\r\nSubject: Contact de %s: %s\r\n\r\n%s",
		from, userName, userSubject, body)

	return smtp.SendMail(addr, auth, from, []string{"support@example.com"}, []byte(msg))
}
