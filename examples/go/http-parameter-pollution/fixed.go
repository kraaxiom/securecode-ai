// fixed.go - Correction CWE-235
// Rejet explicite de tout parametre duplique : un seul champ "amount"
// est accepte, sinon la requete est refusee.
package handlers

import (
	"net/http"
	"strconv"
)

func ApplyDiscountHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// FIXED: refuse toute ambiguite de parametre duplique
	values := r.Form["amount"]
	if len(values) != 1 {
		http.Error(w, "duplicate or missing parameter", http.StatusBadRequest)
		return
	}
	amount, err := strconv.Atoi(values[0])
	if err != nil || amount < 0 {
		http.Error(w, "invalid amount", http.StatusBadRequest)
		return
	}
	w.Write([]byte("discount applied: " + strconv.Itoa(amount)))
}
