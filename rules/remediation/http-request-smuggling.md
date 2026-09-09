# Remédiation — HTTP Request Smuggling

## Principe
Éliminer toute ambiguïté d'interprétation entre les couches HTTP en chaîne (proxy/frontal et backend) sur la délimitation des requêtes. Rejeter explicitement les requêtes présentant à la fois `Content-Length` et `Transfer-Encoding`, et privilégier HTTP/2 de bout en bout quand c'est possible. Ce correctif se situe surtout au niveau infrastructure/configuration, mais le code applicatif doit aussi rejeter les requêtes ambiguës qu'il reçoit directement.

## PHP (validation applicative défensive)
```php
// Avant — vulnérable
// L'application fait confiance aux en-têtes transmis par le proxy sans
// vérifier leur cohérence, ouvrant la porte à une divergence d'interprétation.
$length = $_SERVER['HTTP_CONTENT_LENGTH'] ?? null;
$encoding = $_SERVER['HTTP_TRANSFER_ENCODING'] ?? null;
traiterRequete($length, $encoding);

// Après — sécurisé
$length = $_SERVER['HTTP_CONTENT_LENGTH'] ?? null;
$encoding = $_SERVER['HTTP_TRANSFER_ENCODING'] ?? null;
if ($length !== null && $encoding !== null) {
    http_response_code(400);
    exit('Requête ambiguë refusée (Content-Length et Transfer-Encoding conjoints)');
}
traiterRequete($length, $encoding);
```

## JavaScript / Node.js (serveur HTTP brut)
```js
// Avant — vulnérable
// Le serveur applicatif ne vérifie pas la cohérence des en-têtes reçus,
// ce qui peut être exploité si un proxy en amont interprète différemment
// la même requête.
const server = http.createServer((req, res) => {
  handleRequest(req, res);
});

// Après — sécurisé
const server = http.createServer((req, res) => {
  const hasContentLength = 'content-length' in req.headers;
  const hasTransferEncoding = 'transfer-encoding' in req.headers;
  if (hasContentLength && hasTransferEncoding) {
    res.writeHead(400);
    return res.end('Requête ambiguë refusée');
  }
  handleRequest(req, res);
});
```

## Python (WSGI/Flask, middleware défensif)
```python
# Avant — vulnérable
# L'application traite la requête sans vérifier la cohérence des en-têtes
# de longueur/encodage transmis par la chaîne de proxys.
@app.route("/api", methods=["POST"])
def handle():
    return process(request)

# Après — sécurisé
@app.before_request
def rejeter_requetes_ambigues():
    if "Content-Length" in request.headers and "Transfer-Encoding" in request.headers:
        abort(400, "Requête ambiguë refusée")

@app.route("/api", methods=["POST"])
def handle():
    return process(request)
```

## Checklist de vérification post-patch
- [ ] Le proxy/load balancer en amont est configuré pour rejeter les requêtes contenant à la fois `Content-Length` et `Transfer-Encoding`.
- [ ] La connexion entre le proxy et le backend utilise HTTP/2 de bout en bout, ou à défaut normalise strictement les requêtes HTTP/1.1.
- [ ] L'application applicative rejette elle-même les requêtes ambiguës reçues directement, en défense en profondeur.
- [ ] Les serveurs HTTP et proxys de la chaîne sont maintenus à jour avec les derniers correctifs de parsing.
- [ ] Un test d'intégration confirme qu'une requête avec en-têtes `Content-Length`/`Transfer-Encoding` conjoints est rejetée avec un code 400 à chaque étage de la chaîne.
