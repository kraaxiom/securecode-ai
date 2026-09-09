# Remédiation — Web Cache Poisoning

## Principe
Ne jamais faire confiance aux en-têtes contrôlables par le client (`X-Forwarded-Host`, `X-Forwarded-Scheme`, `X-Original-URL`) pour générer du contenu reflété dans la réponse. Valider contre une liste blanche d'hôtes, et s'assurer que la clé de cache du CDN/reverse proxy inclut tout en-tête influençant réellement le rendu.

## PHP
```php
// Avant — vulnérable : hôte reflété sans validation
$host = $_SERVER['HTTP_X_FORWARDED_HOST'] ?? $_SERVER['HTTP_HOST'];
echo '<link rel="canonical" href="https://' . $host . '/page">';

// Après — sécurisé : liste blanche stricte
$allowedHosts = ['www.example.com', 'app.example.com'];
$host = $_SERVER['HTTP_HOST'];
if (!in_array($host, $allowedHosts, true)) {
    $host = 'www.example.com';
}
echo '<link rel="canonical" href="https://' . htmlspecialchars($host, ENT_QUOTES) . '/page">';
```

## Node.js (Express)
```js
// Avant — vulnérable
app.get('/reset-link', (req, res) => {
  const host = req.headers['x-forwarded-host'] || req.hostname;
  res.send(`https://${host}/reset?token=${token}`);
});

// Après — sécurisé
const ALLOWED_HOSTS = new Set(['www.example.com', 'app.example.com']);
app.get('/reset-link', (req, res) => {
  const host = ALLOWED_HOSTS.has(req.hostname) ? req.hostname : 'www.example.com';
  res.set('Vary', 'Host');
  res.send(`https://${host}/reset?token=${token}`);
});
```

## Python (Flask)
```python
# Avant — vulnérable
host = request.headers.get('X-Forwarded-Host', request.host)
return f"https://{host}/reset?token={token}"

# Après — sécurisé
ALLOWED_HOSTS = {"www.example.com", "app.example.com"}
host = request.host if request.host in ALLOWED_HOSTS else "www.example.com"
response = make_response(f"https://{host}/reset?token={token}")
response.headers['Vary'] = 'Host'
return response
```

## Configuration cache/CDN (Nginx en frontal)
```nginx
# Avant — vulnérable : clé de cache basée uniquement sur l'URL
proxy_cache_key $scheme$host$request_uri;

# Après — sécurisé : normaliser/supprimer les en-têtes non fiables avant le backend,
# et inclure Vary dans la logique de cache si le contenu dépend d'un en-tête
proxy_set_header X-Forwarded-Host "";
proxy_set_header X-Forwarded-Scheme "";
proxy_cache_key $scheme$host$request_uri;
add_header Vary "Host" always;
```

## Checklist de vérification post-patch
- [ ] Aucune génération d'URL/contenu à partir d'en-têtes `X-Forwarded-*`/`X-Original-URL` sans validation contre une liste blanche.
- [ ] La clé de cache du CDN/reverse proxy inclut tout en-tête influençant réellement la réponse, ou ces en-têtes sont supprimés avant le backend.
- [ ] Un en-tête `Vary` approprié est présent sur les réponses dont le contenu dépend d'un en-tête spécifique.
- [ ] Les réponses personnalisées/sensibles portent `Cache-Control: private, no-store`.
