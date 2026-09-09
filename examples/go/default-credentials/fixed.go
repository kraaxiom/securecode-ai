// fixed.go - Correction CWE-1392
// Aucun mot de passe par defaut connu : un mot de passe temporaire
// aleatoire est genere et un changement obligatoire est impose
// avant toute utilisation du compte.
package setup

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
)

const defaultAdminUser = "admin"

// FIXED: generation d'un mot de passe temporaire aleatoire et unique
func generateTemporaryPassword() string {
	buf := make([]byte, 24)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(buf)
}

func InitAdminAccount(db *sql.DB) (string, error) {
	tempPassword := generateTemporaryPassword()

	// FIXED: le compte est marque "must_change_password" et le mot de
	// passe temporaire n'est jamais reutilisable apres la premiere connexion
	_, err := db.Exec(
		"INSERT INTO users (username, password_hash, role, must_change_password) VALUES (?, ?, 'admin', 1)",
		defaultAdminUser, hashPassword(tempPassword),
	)
	if err != nil {
		return "", err
	}
	// Le mot de passe temporaire est communique une seule fois hors bande
	// (ex: canal securise a l'operateur), jamais stocke en clair.
	return tempPassword, nil
}
