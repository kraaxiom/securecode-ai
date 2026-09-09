// vulnerable.go - CWE-1392: Use of Default Credentials
// Le compte administrateur est cree avec un identifiant et un mot de passe
// par defaut codes en dur, jamais forces a etre change au premier demarrage.
package setup

import "database/sql"

// VULNERABLE: identifiants par defaut connus, jamais invalides
const (
	defaultAdminUser = "admin"
	defaultAdminPass = "admin123"
)

func InitAdminAccount(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO users (username, password_hash, role) VALUES (?, ?, 'admin')",
		defaultAdminUser, hashPassword(defaultAdminPass),
	)
	return err
}
