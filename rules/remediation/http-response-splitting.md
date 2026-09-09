# Remédiation — HTTP Response Splitting

## Principe
Ne jamais construire un en-tête HTTP (notamment `Location` ou `Set-Cookie`) ou une redirection à partir d'une entrée utilisateur non filtrée. Rejeter ou supprimer systématiquement les caractères CR/LF (`\r`, `\n`) avant toute écriture dans un en-tête, et privilégier les API modernes du framework qui encodent nativement ces valeurs plutôt que d'écrire le flux HTTP à la main.

## PHP
```php
// Avant — vulnérable
$redirectUrl = $_GET['next'];
header("Location: " . $redirectUrl);

// Après — sécurisé
$allowed = ['/dashboard', '/profile', '/account'];
$redirectUrl = $_GET['next'] ?? '/dashboard';
if (!in_array($redirectUrl, $allowed, true)) {
    $redirectUrl = '/dashboard';
}
// PHP >= 5.1.2 rejette déjà les CR/LF dans header(), la liste blanche
// élimine en plus tout risque de redirection ouverte.
header("Location: " . $redirectUrl);
```

## Node.js (Express)
```js
// Avant — vulnérable
app.get('/redirect', (req, res) => {
  res.setHeader('Location', req.query.next);
  res.status(302).end();
});

// Après — sécurisé
const ALLOWED_PATHS = new Set(['/dashboard', '/profile', '/account']);

app.get('/redirect', (req, res) => {
  const next = ALLOWED_PATHS.has(req.query.next) ? req.query.next : '/dashboard';
  // res.redirect() encode/valide la valeur via l'API Express plutôt
  // qu'une écriture brute d'en-tête, et Node rejette nativement les CR/LF.
  res.redirect(302, next);
});
```

## Python (Flask)
```python
# Avant — vulnérable
@app.route("/redirect")
def redirect_view():
    next_url = request.args.get("next")
    response = make_response("", 302)
    response.headers["Location"] = next_url
    return response

# Après — sécurisé
from urllib.parse import urlparse

ALLOWED_PATHS = {"/dashboard", "/profile", "/account"}

@app.route("/redirect")
def redirect_view():
    next_url = request.args.get("next", "/dashboard")
    parsed = urlparse(next_url)
    # Rejette toute valeur contenant un schéma/hôte (redirection ouverte)
    # ou des caractères de contrôle, puis vérifie la liste blanche.
    if parsed.scheme or parsed.netloc or next_url not in ALLOWED_PATHS:
        next_url = "/dashboard"
    return redirect(next_url, code=302)
```

## Checklist de vérification post-patch
- [ ] Aucune valeur d'entrée utilisateur n'est écrite directement dans un en-tête HTTP (`Location`, `Set-Cookie`, en-tête personnalisé) sans passer par une validation ou une liste blanche.
- [ ] Les caractères `\r` et `\n` sont rejetés ou impossibles à injecter dans toute valeur d'en-tête construite par l'application.
- [ ] Les redirections utilisent l'API de haut niveau du framework (`res.redirect()`, `header()` PHP moderne, `redirect()` Flask) plutôt qu'une écriture brute du flux de réponse.
- [ ] Un test confirme qu'une entrée contenant `%0d%0a` ou `\r\n` ne modifie pas le nombre d'en-têtes renvoyés par le serveur.
- [ ] Le framework/runtime est à jour vers une version qui bloque nativement l'injection CR/LF dans les en-têtes.
