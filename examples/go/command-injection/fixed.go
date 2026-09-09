// fixed.go - Correction CWE-78
// Execution directe sans shell, arguments passes separement,
// et verification que le fichier reste dans le repertoire autorise.
package handlers

import (
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
)

const uploadDir = "/var/app/uploads"

func CompressHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	cleanPath := filepath.Join(uploadDir, filepath.Base(filename))

	// FIXED: empeche la traversee de repertoire et l'injection shell
	if !strings.HasPrefix(cleanPath, uploadDir) {
		http.Error(w, "invalid file", http.StatusBadRequest)
		return
	}
	cmd := exec.Command("tar", "-czf", "archive.tar.gz", cleanPath)
	if err := cmd.Run(); err != nil {
		http.Error(w, "compression failed", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("done"))
}
