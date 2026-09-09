// vulnerable.go - CWE-78: Command Injection
// Un nom de fichier fourni par l'utilisateur est concatene dans une
// commande d'archivage lancee via le shell.
package handlers

import (
	"net/http"
	"os/exec"
)

func CompressHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")

	// VULNERABLE: interpolation dans une commande shell
	cmd := exec.Command("sh", "-c", "tar -czf archive.tar.gz "+filename)
	if err := cmd.Run(); err != nil {
		http.Error(w, "compression failed", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("done"))
}
