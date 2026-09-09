// fixed.go - Correction CWE-287
// Liste blanche stricte des redirect_uri autorises et validation du
// parametre "state" genere par le serveur, liant la requete d'autorisation
// initiale a la reponse du callback (protection anti-CSRF/vol de code).
package handlers

import (
	"net/http"
)

var allowedRedirectURIs = map[string]bool{
	"https://app.example.com/oauth/callback": true,
}

func OAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	redirectURI := r.URL.Query().Get("redirect_uri")
	state := r.URL.Query().Get("state")

	// FIXED: redirect_uri controle contre une liste blanche exacte
	if !allowedRedirectURIs[redirectURI] {
		http.Error(w, "redirect_uri non autorise", http.StatusBadRequest)
		return
	}

	// FIXED: le "state" recu doit correspondre a celui genere et stocke
	// en session au debut du flux (protection CSRF)
	expectedState, ok := getStoredState(r)
	if !ok || state != expectedState {
		http.Error(w, "parametre state invalide", http.StatusBadRequest)
		return
	}
	invalidateStoredState(r)

	token, err := exchangeCodeForToken(code, redirectURI)
	if err != nil {
		http.Error(w, "echec OAuth", http.StatusUnauthorized)
		return
	}

	user := getUserFromToken(token)
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    issueToken(user),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
