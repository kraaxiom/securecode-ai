// vulnerable.go - CWE-284: Improper Access Control
// Le controle d'acces repose uniquement sur l'affichage cote client (un
// lien "Admin" cache dans l'UI) : l'API elle-meme n'applique aucune
// verification serveur, donc l'endpoint reste accessible directement.
package handlers

import (
	"encoding/json"
	"net/http"
)

func AdminSettingsHandler(w http.ResponseWriter, r *http.Request) {
	// VULNERABLE: aucun controle d'acces serveur, la "protection" repose
	// uniquement sur le fait que le lien n'est pas affiche aux non-admins
	settings := getSystemSettings()
	json.NewEncoder(w).Encode(settings)
}
