// vulnerable.go - CWE-444: HTTP Request Smuggling
// Le serveur amont fait confiance a la fois a Content-Length et a
// Transfer-Encoding fournis par le client sans les valider, ouvrant
// la voie a une desynchronisation entre proxy et backend.
package handlers

import (
	"io"
	"net/http"
)

func ProxyForwardHandler(w http.ResponseWriter, r *http.Request) {
	// VULNERABLE: transmission brute des en-tetes Content-Length et
	// Transfer-Encoding sans normalisation ni rejet des ambiguites
	req, _ := http.NewRequest(r.Method, "http://backend.internal"+r.URL.Path, r.Body)
	req.Header = r.Header // copie complete, y compris CL et TE potentiellement contradictoires

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "bad gateway", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}
