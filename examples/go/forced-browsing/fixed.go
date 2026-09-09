// fixed.go - Correction CWE-425
// Chaque acces au repertoire sensible passe par un handler qui applique
// une authentification et une autorisation explicites avant de servir
// le fichier, sans jamais reposer sur l'obscurite de l'URL.
package handlers

import (
	"net/http"
)

// FIXED: middleware d'authentification/autorisation applique avant de
// servir tout fichier du repertoire sensible
func requireInternalAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUser := getCurrentUser(r)
		if currentUser == nil || !currentUser.HasRole("internal-staff") {
			http.Error(w, "acces refuse", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RegisterRoutes(mux *http.ServeMux) {
	reportsHandler := http.StripPrefix("/internal/reports/", http.FileServer(http.Dir("./data/reports")))
	mux.Handle("/internal/reports/", requireInternalAccess(reportsHandler))
}
