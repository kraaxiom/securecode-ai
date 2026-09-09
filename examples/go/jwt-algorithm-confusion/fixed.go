// fixed.go - Correction CWE-347
// L'algorithme attendu est verifie explicitement avant d'utiliser la cle,
// empechant un attaquant de changer d'algorithme (ex: RS256 -> HS256)
// pour forger une signature valide.
package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var rsaPublicKey interface{} // cle publique RSA utilisee pour verifier RS256

// FIXED: seul l'algorithme attendu (RS256) est accepte ; toute autre
// valeur du header "alg" est rejetee explicitement
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("algorithme de signature inattendu")
		}
		return rsaPublicKey, nil
	}, jwt.WithValidMethods([]string{"RS256"}))
}
