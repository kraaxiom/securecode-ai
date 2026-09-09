# Remédiation — CRLF Injection

## Principe
Ne jamais injecter directement une valeur utilisateur dans un en-tête HTTP, une ligne de log ou une commande de protocole texte. Utiliser les API de framework qui gèrent les en-têtes de façon sûre, et rejeter/filtrer toute entrée contenant `\r` ou `\n`.

## PHP
```php
// Avant — vulnérable
$redirect = $_GET['next'];
header("Location: " . $redirect);

// Après — sécurisé
$redirect = $_GET['next'];
if (preg_match('/[\r\n]/', $redirect) || !str_starts_with($redirect, '/')) {
    $redirect = '/';
}
header("Location: " . $redirect);
```

## Node.js (Express)
```js
// Avant — vulnérable
app.get('/redirect', (req, res) => {
  res.setHeader('Location', req.query.next);
  res.status(302).end();
});

// Après — sécurisé
app.get('/redirect', (req, res) => {
  const next = req.query.next;
  const safe = typeof next === 'string' && /^\/[^\r\n]*$/.test(next) ? next : '/';
  res.redirect(302, safe); // Express encode déjà l'en-tête via res.redirect
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
import re

@app.route("/redirect")
def redirect_view():
    next_url = request.args.get("next", "/")
    if not re.fullmatch(r"/[^\r\n]*", next_url):
        next_url = "/"
    return redirect(next_url, code=302)
```

## Checklist de vérification post-patch
- [ ] Aucune valeur utilisateur n'est écrite directement dans un en-tête de réponse sans passer par l'API sûre du framework.
- [ ] Toute entrée destinée à un en-tête, un log ou une commande protocolaire est vérifiée pour l'absence de `\r`/`\n` ou rejetée.
- [ ] Un test confirme qu'une entrée contenant une séquence CRLF est neutralisée (rejetée ou encodée) sans casser le comportement légitime.
- [ ] Les journaux applicatifs utilisent un format structuré (JSON) qui échappe automatiquement les caractères de contrôle.
