// vulnerable.go - CWE-113: HTTP Response Splitting
// L'entree utilisateur est inseree dans un en-tete puis la reponse
// est renvoyee, permettant de scinder la reponse HTTP en y injectant
// un corps ou des en-tetes arbitraires.
package handlers

import "net/http"

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("pref")

	// VULNERABLE: valeur brute utilisee dans un en-tete Set-Cookie manuel
	w.Header().Set("Set-Cookie", "pref="+value)
	w.Write([]byte("saved"))
}
