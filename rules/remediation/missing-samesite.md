# Remédiation — Cookie sans attribut SameSite

## Principe
Définir `SameSite=Lax` (ou `Strict` si aucune navigation cross-site légitime n'est nécessaire) sur tout cookie de session, afin de limiter l'envoi automatique du cookie lors de requêtes cross-site et de réduire l'exposition aux attaques CSRF. Réserver `SameSite=None` aux cas justifiés, toujours associé à `Secure`.

## PHP
```php
// Avant — vulnérable
setcookie('PHPSESSID', session_id(), ['path' => '/']);

// Après — sécurisé
setcookie('PHPSESSID', session_id(), [
    'path' => '/',
    'samesite' => 'Lax',
    'secure' => true,
    'httponly' => true,
]);
// php.ini
ini_set('session.cookie_samesite', 'Lax');
```

## Node.js (Express)
```js
// Avant — vulnérable
res.cookie('sid', sessionId, { httpOnly: true, secure: true });

// Après — sécurisé
res.cookie('sid', sessionId, {
  httpOnly: true,
  secure: true,
  sameSite: 'lax', // ou 'strict' selon le besoin fonctionnel
});
```

## Python (Flask/Django)
```python
# Avant — vulnérable
resp.set_cookie('session_id', session_id, httponly=True, secure=True)

# Après — sécurisé
resp.set_cookie('session_id', session_id, httponly=True, secure=True, samesite='Lax')

# Django settings.py
SESSION_COOKIE_SAMESITE = 'Lax'
```

## Checklist de vérification post-patch
- [ ] Tous les cookies de session portent explicitement `SameSite=Lax` ou `Strict`.
- [ ] Tout usage de `SameSite=None` est justifié fonctionnellement et associé à `Secure`.
- [ ] Une protection CSRF applicative (token anti-CSRF) reste en place en complément.
- [ ] Un test confirme qu'une requête cross-site forgée n'entraîne plus l'envoi automatique du cookie de session.
