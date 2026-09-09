// fixed.go - Correction CWE-330
// L'identifiant de session est genere avec un generateur pseudo-aleatoire
// cryptographiquement sur (crypto/rand), avec une entropie suffisante
// pour rendre toute prediction ou enumeration impossible en pratique.
package auth

import (
	"crypto/rand"
	"encoding/base64"
)

// FIXED: identifiant de session genere via crypto/rand, 256 bits d'entropie
func GenerateSessionID() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}
