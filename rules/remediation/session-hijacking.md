# Remédiation — Session Hijacking

## Principe
Fixer une durée de vie de session raisonnable, invalider la session côté serveur à la déconnexion (pas seulement le cookie client), ne jamais journaliser l'identifiant de session en clair, et protéger le transport avec `HttpOnly`, `Secure` et `SameSite`.

## PHP
```php
// Avant — vulnérable
function logout() {
    setcookie('PHPSESSID', '', time() - 3600); // supprime seulement le cookie côté client
}

// Après — sécurisé
ini_set('session.gc_maxlifetime', 1800); // 30 min

function logout() {
    session_start();
    $_SESSION = [];
    session_destroy(); // invalide la session côté serveur
    setcookie('PHPSESSID', '', time() - 3600, '/', '', true, true);
}
```

## Node.js (Express)
```js
// Avant — vulnérable
app.get('/logout', (req, res) => {
  res.clearCookie('sid');
  res.redirect('/');
});

// Après — sécurisé
app.use(session({
  secret: process.env.SESSION_SECRET,
  cookie: { maxAge: 30 * 60 * 1000, httpOnly: true, secure: true, sameSite: 'lax' },
}));

app.get('/logout', (req, res) => {
  req.session.destroy((err) => { // invalide la session côté serveur (store)
    if (err) return res.status(500).end();
    res.clearCookie('sid');
    res.redirect('/');
  });
});
```

## Python (Flask/Django)
```python
# Avant — vulnérable
@app.route('/logout')
def logout():
    resp = redirect('/')
    resp.delete_cookie('session_id')
    return resp

# Après — sécurisé (Flask)
app.config['PERMANENT_SESSION_LIFETIME'] = timedelta(minutes=30)

@app.route('/logout')
def logout():
    session.clear()  # invalide côté serveur
    resp = redirect('/')
    resp.delete_cookie('session_id')
    return resp

# Django : SESSION_COOKIE_AGE = 1800, puis
from django.contrib.auth import logout as django_logout
def logout_view(request):
    django_logout(request)  # flush() la session côté serveur
    return redirect('/')
```

## Checklist de vérification post-patch
- [ ] La déconnexion détruit la session côté serveur, pas seulement le cookie client.
- [ ] Une durée de vie de session raisonnable est configurée avec renouvellement périodique.
- [ ] Aucun log applicatif n'affiche l'identifiant de session en clair.
- [ ] Les cookies de session portent `HttpOnly`, `Secure` et `SameSite`.
