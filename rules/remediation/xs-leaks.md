# Remédiation — XS-Leaks (Cross-Site Leaks)

## Principe
Définir `Cross-Origin-Opener-Policy: same-origin`, `Cross-Origin-Embedder-Policy: require-corp` et `Cross-Origin-Resource-Policy: same-origin` sur les réponses sensibles, uniformiser les comportements observables entre états authentifié/non authentifié, et utiliser `SameSite=Strict`/`Lax` sur les cookies de session.

## PHP
```php
// Avant — vulnérable : aucun en-tête cross-origin, cookie sans SameSite
setcookie('session', $sessionId, ['httponly' => true]);

// Après — sécurisé
header('Cross-Origin-Opener-Policy: same-origin');
header('Cross-Origin-Embedder-Policy: require-corp');
header('Cross-Origin-Resource-Policy: same-origin');
setcookie('session', $sessionId, [
    'httponly' => true,
    'samesite' => 'Strict',
    'secure' => true,
]);
```

## Node.js (Express)
```js
// Avant — vulnérable
app.use(session({ cookie: { httpOnly: true } }));

// Après — sécurisé
app.use((req, res, next) => {
  res.set('Cross-Origin-Opener-Policy', 'same-origin');
  res.set('Cross-Origin-Embedder-Policy', 'require-corp');
  res.set('Cross-Origin-Resource-Policy', 'same-origin');
  next();
});
app.use(session({
  cookie: { httpOnly: true, sameSite: 'strict', secure: true },
}));
```

## Python (Flask)
```python
# Avant — vulnérable
app.config['SESSION_COOKIE_HTTPONLY'] = True

# Après — sécurisé
app.config['SESSION_COOKIE_HTTPONLY'] = True
app.config['SESSION_COOKIE_SAMESITE'] = 'Strict'
app.config['SESSION_COOKIE_SECURE'] = True

@app.after_request
def add_isolation_headers(response):
    response.headers['Cross-Origin-Opener-Policy'] = 'same-origin'
    response.headers['Cross-Origin-Embedder-Policy'] = 'require-corp'
    response.headers['Cross-Origin-Resource-Policy'] = 'same-origin'
    return response
```

## Comportement observable uniforme (exemple: endpoint dont le statut varie selon l'authentification)
```js
// Avant — vulnérable : redirection uniquement si authentifié (fuite observable cross-origin)
app.get('/dashboard', (req, res) => {
  if (!req.user) return res.redirect('/login');
  res.render('dashboard');
});

// Après — atténué : réponse 200 systématique, contenu conditionné côté rendu, pas de redirection distinctive
app.get('/dashboard', (req, res) => {
  res.status(200);
  if (!req.user) return res.render('login-prompt');
  res.render('dashboard');
});
```

## Checklist de vérification post-patch
- [ ] `Cross-Origin-Opener-Policy: same-origin` et `Cross-Origin-Resource-Policy: same-origin` sont présents sur les réponses sensibles.
- [ ] Les cookies de session utilisent `SameSite=Strict` ou `Lax` et `Secure`.
- [ ] Les comportements observables (statut HTTP, redirection, taille) entre états authentifié/non authentifié ont été revus pour les endpoints sensibles.
- [ ] Test : une requête cross-origin passive (balise `<img>`/`<script>`) vers un endpoint sensible ne permet plus de déduire l'état d'authentification.
