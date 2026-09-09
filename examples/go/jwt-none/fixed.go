// fixed.go - Correction CWE-347
// L'algorithme "none" et tout algorithme non attendu sont explicitement
// rejetes avant toute verification de signature.
package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("change-me") // en pratique : secret fort, gere via un coffre-fort de secrets

// FIXED: seule la liste blanche d'algorithmes attendus est acceptee ;
// "none" est automatiquement exclu car absent de cette liste
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}, jwt.WithValidMethods([]string{"HS256"}))
}
