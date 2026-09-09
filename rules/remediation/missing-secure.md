# Remédiation — Cookie sans attribut Secure

## Principe
Définir systématiquement l'attribut `Secure` sur tout cookie de session en production, et forcer HTTPS sur l'ensemble de l'application (redirection + HSTS) pour que ce flag soit pleinement efficace et empêcher toute transmission du cookie en clair sur HTTP.

## PHP
```php
// Avant — vulnérable
setcookie('PHPSESSID', session_id(), ['path' => '/', 'httponly' => true]);

// Après — sécurisé
setcookie('PHPSESSID', session_id(), [
    'path' => '/',
    'httponly' => true,
    'secure' => true, // uniquement transmis sur HTTPS
    'samesite' => 'Lax',
]);
// php.ini
ini_set('session.cookie_secure', '1');
```

## Node.js (Express)
```js
// Avant — vulnérable
res.cookie('sid', sessionId, { httpOnly: true });

// Après — sécurisé
app.set('trust proxy', 1); // si derrière un reverse proxy TLS
res.cookie('sid', sessionId, {
  httpOnly: true,
  secure: true,
  sameSite: 'lax',
});

// Forcer HTTPS
app.use((req, res, next) => {
  if (!req.secure) return res.redirect(301, `https://${req.headers.host}${req.url}`);
  next();
});
```

## Python (Flask/Django)
```python
# Avant — vulnérable
resp.set_cookie('session_id', session_id, httponly=True)

# Après — sécurisé
resp.set_cookie('session_id', session_id, httponly=True, secure=True, samesite='Lax')

# Flask config globale
app.config['SESSION_COOKIE_SECURE'] = True

# Django settings.py
SESSION_COOKIE_SECURE = True
SECURE_SSL_REDIRECT = True
```

## Checklist de vérification post-patch
- [ ] Tous les cookies de session en production portent l'attribut `Secure`.
- [ ] L'application force une redirection HTTP vers HTTPS sur tous les endpoints.
- [ ] L'en-tête HSTS est présent pour renforcer l'application de HTTPS.
- [ ] Le flag `Secure` reste désactivé uniquement en environnement de développement local, jamais en production.
