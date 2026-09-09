// vulnerable.go - CWE-235: HTTP Parameter Pollution
// Le code prend naivement la premiere valeur d'un parametre repete
// sans que la logique metier (ex: montant) ne soit coherente avec
// ce que valide un composant en amont (WAF, proxy) qui pourrait lire
// une autre occurrence.
package handlers

import (
	"net/http"
	"strconv"
)

func ApplyDiscountHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// VULNERABLE: en cas de parametres dupliques (?amount=10&amount=99999),
	// le comportement differe selon le composant qui lit la requete
	amountStr := r.Form.Get("amount")
	amount, _ := strconv.Atoi(amountStr)
	w.Write([]byte("discount applied: " + strconv.Itoa(amount)))
}
