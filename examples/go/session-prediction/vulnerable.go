// vulnerable.go - CWE-330: Use of Insufficiently Random Values
// L'identifiant de session est genere a partir d'un compteur sequentiel et
// de l'horodatage, des valeurs previsibles qu'un attaquant peut deviner
// ou enumerer pour usurper la session d'un autre utilisateur.
package auth

import (
	"fmt"
	"time"
)

var sessionCounter int64

// VULNERABLE: identifiant de session previsible (compteur + timestamp)
func GenerateSessionID() string {
	sessionCounter++
	return fmt.Sprintf("%d-%d", time.Now().Unix(), sessionCounter)
}
