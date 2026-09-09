// vulnerable.go - CWE-287: Improper Authentication
// Le callback OAuth n'a pas de liste blanche de redirect_uri et n'utilise
// aucun parametre "state" pour lier la requete d'autorisation a la reponse,
// ouvrant la voie a du vol de code d'autorisation et du CSRF sur le login.
package handlers

import (
	"net/http"
)

func OAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	redirectURI := r.URL.Query().Get("redirect_uri")

	// VULNERABLE: redirect_uri controle par le client, non valide contre
	// une liste blanche -> vol du code d'autorisation possible
	token, err := exchangeCodeForToken(code, redirectURI)
	if err != nil {
		http.Error(w, "echec OAuth", http.StatusUnauthorized)
		return
	}

	// VULNERABLE: aucune verification du parametre "state" -> CSRF possible
	// sur le flux de connexion
	user := getUserFromToken(token)
	http.SetCookie(w, &http.Cookie{Name: "session", Value: issueToken(user)})
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
