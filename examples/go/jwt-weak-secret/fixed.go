// fixed.go - Correction CWE-326
// Le secret HMAC est un secret fort (>= 256 bits d'entropie), genere
// aleatoirement et charge depuis un coffre-fort de secrets / variable
// d'environnement, jamais code en dur.
package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// FIXED: le secret est charge depuis l'environnement (backe par un
// gestionnaire de secrets en production), jamais present dans le code source
func loadSecretKey() ([]byte, error) {
	secret := os.Getenv("JWT_HMAC_SECRET")
	if len(secret) < 32 {
		// FIXED: refus explicite si le secret est absent ou trop court
		// (< 256 bits d'entropie recommandee pour HS256)
		return nil, errors.New("JWT_HMAC_SECRET manquant ou trop faible")
	}
	return []byte(secret), nil
}

func SignToken(claims jwt.MapClaims) (string, error) {
	secretKey, err := loadSecretKey()
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
