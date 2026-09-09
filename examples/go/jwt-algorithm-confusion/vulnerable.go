// vulnerable.go - CWE-347: Improper Verification of Cryptographic Signature
// La verification du JWT accepte n'importe quel algorithme annonce dans le
// header du token (dont "none" ou HMAC avec une cle publique RSA), ce qui
// permet a un attaquant de forger un token accepte comme valide.
package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

var rsaPublicKey interface{} // cle publique RSA utilisee pour verifier RS256

// VULNERABLE: le callback de verification ne controle pas l'algorithme
// annonce par le token, il fait confiance au header attaquant-controle
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Pas de verification de t.Method : un attaquant peut passer
		// alg=HS256 et signer avec la cle publique RSA (connue) comme secret HMAC
		return rsaPublicKey, nil
	})
}
