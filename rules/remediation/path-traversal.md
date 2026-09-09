# Remédiation — Path Traversal

## Principe
Ne jamais construire un chemin de fichier (inclusion, lecture, écriture, extraction d'archive) directement à partir d'une entrée utilisateur. Toujours résoudre le chemin en forme canonique (absolue) et vérifier qu'il reste strictement sous le répertoire racine autorisé, ou utiliser un mapping fermé (ID -> chemin fixe) plutôt qu'un chemin fourni par le client.

## PHP
```php
// Avant — vulnérable
$page = $_GET['page'];
include('/var/www/app/pages/' . $page . '.php');

// Après — sécurisé
$allowed = ['home', 'contact', 'about'];
$page = $_GET['page'] ?? 'home';
if (!in_array($page, $allowed, true)) {
    http_response_code(400);
    exit('Page invalide');
}
include('/var/www/app/pages/' . $page . '.php');
```

## Node.js (Express)
```js
// Avant — vulnérable
app.get('/file', (req, res) => {
  res.sendFile(path.join('/var/www/uploads', req.query.name));
});

// Après — sécurisé
const baseDir = path.resolve('/var/www/uploads');
app.get('/file', (req, res) => {
  const requested = path.resolve(baseDir, req.query.name);
  if (!requested.startsWith(baseDir + path.sep)) {
    return res.status(400).send('Chemin invalide');
  }
  res.sendFile(requested);
});
```

## Python (Flask)
```python
# Avant — vulnérable
@app.route('/file')
def get_file():
    name = request.args.get('name')
    return send_file(os.path.join('/var/www/uploads', name))

# Après — sécurisé
BASE_DIR = os.path.realpath('/var/www/uploads')

@app.route('/file')
def get_file():
    name = request.args.get('name', '')
    requested = os.path.realpath(os.path.join(BASE_DIR, name))
    if not requested.startswith(BASE_DIR + os.sep):
        abort(400)
    return send_file(requested)
```

## Checklist de vérification post-patch
- [ ] Le chemin final est résolu en forme canonique (`realpath`/`path.resolve`/`os.path.realpath`) avant toute opération fichier.
- [ ] Une vérification confirme que le chemin résolu commence strictement par le répertoire racine autorisé (avec séparateur, pas de simple `startsWith` sans séparateur).
- [ ] Un test confirme qu'une entrée contenant `../../../etc/passwd` (ou équivalent Windows) est rejetée.
- [ ] Un test confirme qu'un chemin absolu ou un caractère nul dans l'entrée est rejeté.
- [ ] Les noms de fichiers stockés/générés côté serveur ne dépendent pas du nom fourni par le client (préférer un UUID).
