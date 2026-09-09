// fixed.go - Correction CWE-444
// Rejet des requetes ambigues (Content-Length ET Transfer-Encoding
// presents simultanement) avant tout transfert vers le backend,
// conformement a la RFC 7230.
package handlers

import (
	"io"
	"net/http"
)

func ProxyForwardHandler(w http.ResponseWriter, r *http.Request) {
	// FIXED: rejet explicite des requetes ambigues CL/TE
	hasCL := r.Header.Get("Content-Length") != ""
	hasTE := r.Header.Get("Transfer-Encoding") != ""
	if hasCL && hasTE {
		http.Error(w, "ambiguous request", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest(r.Method, "http://backend.internal"+r.URL.Path, r.Body)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	req.Header = r.Header.Clone()
	req.Header.Del("Transfer-Encoding") // le client Go gere le chunking lui-meme

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "bad gateway", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}
