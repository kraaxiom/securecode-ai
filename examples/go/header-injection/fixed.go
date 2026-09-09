// fixed.go - Correction CWE-113
// Validation stricte via allowlist des valeurs de langue acceptees ;
// aucune valeur brute non filtree n'atteint l'API d'en-tetes.
package handlers

import "net/http"

var allowedLangs = map[string]bool{"fr": true, "en": true, "es": true}

func SetLangHeaderHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")

	// FIXED: allowlist stricte, pas de caracteres de controle possibles
	if !allowedLangs[lang] {
		lang = "en"
	}
	w.Header().Set("X-User-Lang", lang)
	w.Write([]byte("ok"))
}
