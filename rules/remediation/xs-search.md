# Remédiation — XS-Search (Cross-Site Search)

## Principe
Exiger une méthode POST avec vérification d'origine/jeton pour toute recherche portant sur des données sensibles, ajouter `Cross-Origin-Resource-Policy: same-origin` sur les réponses de recherche, et normaliser le temps/la taille de réponse pour éviter toute corrélation avec le volume de résultats.

## PHP
```php
// Avant — vulnérable : recherche en GET, chargeable via une balise cross-origin
if ($_SERVER['REQUEST_METHOD'] === 'GET' && isset($_GET['q'])) {
    echo json_encode(searchRecords($_GET['q']));
}

// Après — sécurisé : POST + vérification d'origine + en-tête CORP
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $origin = $_SERVER['HTTP_ORIGIN'] ?? '';
    if (!in_array($origin, ['https://app.example.com'], true)) {
        http_response_code(403);
        exit;
    }
    header('Cross-Origin-Resource-Policy: same-origin');
    $q = json_decode(file_get_contents('php://input'), true)['q'] ?? '';
    echo json_encode(searchRecords($q));
}
```

## Node.js (Express)
```js
// Avant — vulnérable : GET sans vérification d'origine ni CORP
app.get('/search', (req, res) => {
  res.json(searchRecords(req.query.q));
});

// Après — sécurisé
const ALLOWED_ORIGIN = 'https://app.example.com';
app.post('/search', (req, res) => {
  if (req.headers.origin !== ALLOWED_ORIGIN) {
    return res.status(403).end();
  }
  res.set('Cross-Origin-Resource-Policy', 'same-origin');
  res.json(searchRecords(req.body.q));
});
```

## Python (Flask) — normalisation du temps de réponse
```python
# Avant — vulnérable : temps de réponse corrélé au nombre de résultats
@app.route('/search', methods=['GET'])
def search():
    results = search_records(request.args.get('q'))
    return jsonify(results)

# Après — sécurisé : POST + origine vérifiée + padding de temps constant
import time

ALLOWED_ORIGIN = "https://app.example.com"

@app.route('/search', methods=['POST'])
def search():
    if request.headers.get('Origin') != ALLOWED_ORIGIN:
        abort(403)
    start = time.monotonic()
    results = search_records(request.json.get('q'))
    elapsed = time.monotonic() - start
    target = 0.3
    if elapsed < target:
        time.sleep(target - elapsed)
    response = jsonify(results)
    response.headers['Cross-Origin-Resource-Policy'] = 'same-origin'
    return response
```

## Checklist de vérification post-patch
- [ ] Les endpoints de recherche/filtrage sur données sensibles exigent POST + vérification d'origine ou jeton anti-CSRF.
- [ ] `Cross-Origin-Resource-Policy: same-origin` est présent sur les réponses de recherche.
- [ ] Le temps de réponse et/ou la taille de réponse ne varient plus de façon exploitable selon le nombre de résultats.
- [ ] Test : une requête de recherche déclenchée depuis une origine tierce via balise passive échoue (403) et ne fuit aucune information temporelle exploitable.
