# Remédiation — Redirection ouverte (Open Redirect)

## Principe
Valider les URL de redirection contre une liste blanche stricte de chemins relatifs ou de domaines autorisés. Refuser les URL absolues ou protocol-relative fournies par l'utilisateur.

## PHP
```php
// Avant — vulnérable
$redirect = $_GET['redirect'];
header("Location: $redirect");

// Après — sécurisé
$allowedPaths = ['/dashboard', '/profile', '/settings'];
$redirect = $_GET['redirect'] ?? '/dashboard';
if (!in_array($redirect, $allowedPaths, true)) {
    $redirect = '/dashboard';
}
header("Location: $redirect");
```

## PHP (Laravel)
```php
// Avant — vulnérable
return redirect($request->input('next'));

// Après — sécurisé
$next = $request->input('next');
if (!Str::startsWith($next, '/') || Str::startsWith($next, '//')) {
    $next = '/dashboard'; // rejette les URL absolues et protocol-relative
}
return redirect($next);
```

## Node.js (Express)
```js
// Avant — vulnérable
app.get('/logout', (req, res) => {
  res.redirect(req.query.returnUrl);
});

// Après — sécurisé
const ALLOWED_PATHS = new Set(['/login', '/home']);
app.get('/logout', (req, res) => {
  const returnUrl = req.query.returnUrl;
  res.redirect(ALLOWED_PATHS.has(returnUrl) ? returnUrl : '/login');
});
```

## Python (Flask)
```python
# Avant — vulnérable
next_url = request.args.get('next')
return redirect(next_url)

# Après — sécurisé
from urllib.parse import urlparse

def is_safe_url(target):
    ref = urlparse(request.host_url)
    test = urlparse(target)
    return test.netloc in ('', ref.netloc) and test.scheme in ('', 'http', 'https')

next_url = request.args.get('next', '/')
if not is_safe_url(next_url):
    next_url = '/'
return redirect(next_url)
```

## Checklist de vérification post-patch
- [ ] Toute destination de redirection issue d'un paramètre client est validée contre une liste blanche de chemins/domaines.
- [ ] Les URL absolues et protocol-relative (`//evil.com`) fournies par l'utilisateur sont rejetées pour les redirections internes.
- [ ] La comparaison de domaine se fait par égalité stricte de l'hôte, jamais par sous-chaîne (`includes`).
- [ ] Un test confirme qu'une valeur de redirection non autorisée retombe sur une destination par défaut sûre.
