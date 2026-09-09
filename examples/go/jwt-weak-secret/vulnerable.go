// vulnerable.go - CWE-326: Inadequate Encryption Strength
// Le secret HMAC utilise pour signer les JWT est court, previsible et
// code en dur, ce qui le rend cassable par force brute hors ligne.
package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

// VULNERABLE: secret court, faible entropie, code en dur dans le source
var secretKey = []byte("secret123")

func SignToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
