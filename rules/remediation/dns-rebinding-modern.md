# Remédiation — DNS Rebinding

## Principe
Ne jamais fonder une décision d'accès uniquement sur l'en-tête `Host`/`Origin` d'une requête reçue par un service local ou interne. Exiger une authentification explicite et valider strictement l'en-tête `Host` contre une liste de valeurs attendues côté serveur.

## Node.js (service local / API de développement)
```js
// Avant — vulnérable : aucune vérification, accessible depuis n'importe quelle origine
const http = require('http');
http.createServer((req, res) => {
  res.end(JSON.stringify(getInternalStatus()));
}).listen(3000, '127.0.0.1');

// Après — sécurisé : validation stricte du Host + authentification
const ALLOWED_HOSTS = new Set(['localhost:3000', '127.0.0.1:3000']);
http.createServer((req, res) => {
  if (!ALLOWED_HOSTS.has(req.headers.host)) {
    res.writeHead(403); return res.end('Forbidden: invalid Host header');
  }
  const token = req.headers['x-local-token'];
  if (token !== process.env.LOCAL_API_TOKEN) {
    res.writeHead(401); return res.end('Unauthorized');
  }
  res.end(JSON.stringify(getInternalStatus()));
}).listen(3000, '127.0.0.1');
```

## Python (Flask, API locale)
```python
# Avant — vulnérable
@app.route('/internal/status')
def status():
    return jsonify(get_internal_status())

# Après — sécurisé
ALLOWED_HOSTS = {"localhost:5000", "127.0.0.1:5000"}

@app.before_request
def check_host_and_auth():
    if request.host not in ALLOWED_HOSTS:
        abort(403, "Invalid Host header")
    if request.headers.get('X-Local-Token') != current_app.config['LOCAL_API_TOKEN']:
        abort(401)

@app.route('/internal/status')
def status():
    return jsonify(get_internal_status())
```

## Go
```go
// Avant — vulnérable
http.HandleFunc("/internal/status", func(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(getInternalStatus())
})

// Après — sécurisé
var allowedHosts = map[string]bool{"localhost:8080": true, "127.0.0.1:8080": true}

http.HandleFunc("/internal/status", func(w http.ResponseWriter, r *http.Request) {
    if !allowedHosts[r.Host] {
        http.Error(w, "invalid Host header", http.StatusForbidden)
        return
    }
    if r.Header.Get("X-Local-Token") != os.Getenv("LOCAL_API_TOKEN") {
        http.Error(w, "unauthorized", http.StatusUnauthorized)
        return
    }
    json.NewEncoder(w).Encode(getInternalStatus())
})
```

## Checklist de vérification post-patch
- [ ] Tout service local/interne exige une authentification explicite (jeton, session), indépendamment de l'origine réseau supposée.
- [ ] L'en-tête `Host` est validé côté serveur contre une liste stricte de valeurs attendues.
- [ ] Un test confirme qu'une requête avec un `Host` non listé est rejetée (403) avant tout traitement métier.
- [ ] Le TTL DNS et le pinning ne sont pas la seule protection retenue (la validation applicative reste indépendante du réseau).
