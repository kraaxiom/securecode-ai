# Remédiation — Remote File Inclusion (RFI)

## Principe
Ne jamais permettre qu'une valeur pouvant contenir un schéma d'URL (`http://`, `https://`, `ftp://`, `php://`) atteigne une fonction d'inclusion dynamique. Désactiver `allow_url_include` au niveau du runtime PHP et restreindre `allow_url_fopen`, en complément d'une whitelist stricte de valeurs autorisées côté application.

## PHP
```php
// Avant — vulnérable (php.ini: allow_url_include=On)
$module = $_GET['module'];
include($module . '.php');

// Après — sécurisé
$allowedModules = ['dashboard', 'profile', 'settings'];
$module = $_GET['module'] ?? 'dashboard';
if (!in_array($module, $allowedModules, true)) {
    http_response_code(400);
    exit('Module invalide');
}
include(__DIR__ . '/modules/' . $module . '.php');
```

```ini
; php.ini — configuration recommandée
allow_url_include = Off
allow_url_fopen = Off
```

## Node.js
```js
// Avant — vulnérable (chargement dynamique par URL/chemin fourni par le client)
const mod = require(req.query.module);

// Après — sécurisé
const ALLOWED = {
  dashboard: require('./modules/dashboard'),
  profile: require('./modules/profile'),
};
const mod = ALLOWED[req.query.module];
if (!mod) return res.status(400).send('Module invalide');
```

## Python (Flask/Django)
```python
# Avant — vulnérable
module_name = request.args.get('module')
module = importlib.import_module(module_name)

# Après — sécurisé
ALLOWED_MODULES = {"dashboard", "profile", "settings"}
module_name = request.args.get('module', 'dashboard')
if module_name not in ALLOWED_MODULES:
    abort(400)
module = importlib.import_module(f"app.modules.{module_name}")
```

## Checklist de vérification post-patch
- [ ] `allow_url_include` est désactivé (`Off`) dans la configuration PHP de production.
- [ ] `allow_url_fopen` est désactivé ou strictement nécessaire et justifié.
- [ ] Aucune valeur pouvant contenir un schéma d'URL n'atteint une fonction d'inclusion/chargement dynamique.
- [ ] Un mapping fermé (whitelist) remplace toute construction dynamique du chemin/module inclus.
