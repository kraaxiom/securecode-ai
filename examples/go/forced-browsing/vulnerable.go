// vulnerable.go - CWE-425: Direct Request ('Forced Browsing')
// Les pages/fichiers sensibles (rapports internes, exports) sont exposes
// sur des chemins previsibles sans aucun controle d'acces, en comptant
// uniquement sur le fait que l'URL n'est pas publiee.
package handlers

import (
	"net/http"
)

// VULNERABLE: le repertoire "/internal/reports/" est servi statiquement
// sans authentification ni autorisation, seule l'obscurite de l'URL protege
func RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/internal/reports/", http.StripPrefix("/internal/reports/", http.FileServer(http.Dir("./data/reports"))))
}
