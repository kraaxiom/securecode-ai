// fixed.go - Correction CWE-78
// Appel direct du binaire avec arguments separes (pas de shell),
// et validation stricte du format de l'entree (allowlist regex).
package handlers

import (
	"net/http"
	"os/exec"
	"regexp"
)

var hostPattern = regexp.MustCompile(`^[a-zA-Z0-9.\-]{1,253}$`)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")
	if !hostPattern.MatchString(host) {
		http.Error(w, "invalid host", http.StatusBadRequest)
		return
	}

	// FIXED: pas de shell, arguments passes individuellement
	cmd := exec.Command("ping", "-c", "1", host)
	out, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "ping failed", http.StatusInternalServerError)
		return
	}
	w.Write(out)
}
