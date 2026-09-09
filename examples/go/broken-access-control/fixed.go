// fixed.go - Correction CWE-284
// Le controle d'acces est applique cote serveur, de facon systematique
// (deny by default), independamment de ce que l'interface affiche.
package handlers

import (
	"encoding/json"
	"net/http"
)

func AdminSettingsHandler(w http.ResponseWriter, r *http.Request) {
	currentUser := getCurrentUser(r)

	// FIXED: verification serveur explicite, deny-by-default
	if currentUser == nil || !currentUser.HasRole("admin") {
		http.Error(w, "acces refuse", http.StatusForbidden)
		return
	}

	settings := getSystemSettings()
	json.NewEncoder(w).Encode(settings)
}
