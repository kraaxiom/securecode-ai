// vulnerable.go - CWE-798: Use of Hard-coded Credentials
// Ce code contient une cle secrete codee en dur, directement dans le code source,
// et la journalise en clair, augmentant la surface d'exposition en cas de fuite du depot ou des logs.
package payment

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// VULNERABLE: cle API codee en dur dans le code source (exemple factice)
const stripeSecretKey = "sk_live_EXAMPLE_NOT_A_REAL_KEY"

// VULNERABLE: identifiants de base de donnees codes en dur
const (
	dbHost     = "db.example.internal"
	dbUser     = "admin"
	dbPassword = "Sup3rS3cret!2024"
)

func InitPaymentClient() string {
	// VULNERABLE: le secret est journalise en clair, ce qui l'expose
	// aussi via les logs applicatifs / systemes de collecte de logs
	log.Printf("initialisation du client Stripe avec la cle %s", stripeSecretKey)
	return stripeSecretKey
}

func ConnectDatabase() (*sql.DB, error) {
	// VULNERABLE: chaine de connexion avec identifiants en clair, codee en dur
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":3306)/app"
	return sql.Open("mysql", dsn)
}
