# Remédiation — Cookie sans attribut HttpOnly

## Principe
Définir explicitement l'attribut `HttpOnly` sur tout cookie de session ou cookie sensible qui n'a pas besoin d'être lu par du JavaScript côté client, afin de limiter l'impact d'une éventuelle faille XSS sur le vol de session.

## PHP
```php
// Avant — vulnérable
setcookie('PHPSESSID', session_id(), ['path' => '/']);

// Après — sécurisé
setcookie('PHPSESSID', session_id(), [
    'path' => '/',
    'httponly' => true,
    'secure' => true,
    'samesite' => 'Lax',
]);
// Ou globalement dans php.ini / avant session_start():
ini_set('session.cookie_httponly', '1');
```

## Node.js (Express)
```js
// Avant — vulnérable
res.cookie('sid', sessionId, { path: '/' });

// Après — sécurisé
res.cookie('sid', sessionId, {
  path: '/',
  httpOnly: true,
  secure: true,
  sameSite: 'lax',
});

// Ou via express-session
app.use(session({
  secret: process.env.SESSION_SECRET,
  cookie: { httpOnly: true, secure: true, sameSite: 'lax' },
}));
```

## Python (Flask/Django)
```python
# Avant — vulnérable
resp.set_cookie('session_id', session_id)

# Après — sécurisé
resp.set_cookie('session_id', session_id, httponly=True, secure=True, samesite='Lax')

# Flask config globale
app.config['SESSION_COOKIE_HTTPONLY'] = True

# Django settings.py
SESSION_COOKIE_HTTPONLY = True
```

## Checklist de vérification post-patch
- [ ] Tous les cookies de session émis par l'application portent l'attribut `HttpOnly`.
- [ ] Le réglage global de session du framework impose `HttpOnly` par défaut.
- [ ] Les outils de dev navigateur confirment que `document.cookie` ne retourne plus le cookie de session.
- [ ] Combiné avec `Secure` et `SameSite` sur le même cookie.
