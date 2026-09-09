// vulnerable.go - CWE-78: OS Command Injection
// L'entree utilisateur est passee a un shell via "sh -c", permettant
// l'injection de metacaracteres shell (;, |, &&) pour executer des
// commandes arbitraires.
package handlers

import (
	"net/http"
	"os/exec"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")

	// VULNERABLE: entree utilisateur interpolee dans une commande shell
	cmd := exec.Command("sh", "-c", "ping -c 1 "+host)
	out, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "ping failed", http.StatusInternalServerError)
		return
	}
	w.Write(out)
}
