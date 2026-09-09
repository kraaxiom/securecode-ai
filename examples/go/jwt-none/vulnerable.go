// vulnerable.go - CWE-347: Improper Verification of Cryptographic Signature
// Le callback de verification renvoie la cle sans jamais controler que
// l'algorithme n'est pas "none", ce qui permet a un attaquant d'envoyer
// un token non signe et de le faire accepter comme valide.
package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("change-me")

// VULNERABLE: aucune restriction sur les algorithmes acceptes,
// un token avec alg="none" et signature vide passe la verification
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}
