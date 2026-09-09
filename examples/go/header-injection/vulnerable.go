// vulnerable.go - CWE-113: HTTP Header Injection
// Une valeur utilisateur est copiee directement dans un en-tete de
// reponse personnalise sans validation.
package handlers

import "net/http"

func SetLangHeaderHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")

	// VULNERABLE: injection possible via CR/LF dans "lang"
	w.Header().Set("X-User-Lang", lang)
	w.Write([]byte("ok"))
}
