# Remédiation — OS Command Injection

## Principe
Éviter tout appel shell construit à partir d'une entrée utilisateur. Privilégier les API d'exécution qui prennent un tableau d'arguments (sans interprétation shell), et valider les entrées selon une liste blanche stricte (ex: extension de fichier, format numérique) quand un appel externe reste nécessaire.

## PHP
```php
// Avant — vulnérable
$host = $_GET['host'];
system("ping -c 4 " . $host);

// Après — sécurisé
$host = $_GET['host'];
if (!filter_var($host, FILTER_VALIDATE_IP) && !preg_match('/^[a-zA-Z0-9.-]+$/', $host)) {
    throw new InvalidArgumentException('Hôte invalide');
}
$escaped = escapeshellarg($host);
system("ping -c 4 " . $escaped);
// Préférable encore : proc_open avec un tableau d'arguments, sans passer par le shell
```

## JavaScript / Node.js (child_process)
```js
// Avant — vulnérable
const { exec } = require('child_process');
const host = req.query.host;
exec(`ping -c 4 ${host}`, callback);

// Après — sécurisé
const { execFile } = require('child_process');
const host = req.query.host;
if (!/^[a-zA-Z0-9.-]+$/.test(host)) {
  return res.status(400).json({ error: 'Hôte invalide' });
}
execFile('ping', ['-c', '4', host], callback); // pas d'interprétation shell
```

## Python (subprocess)
```python
# Avant — vulnérable
host = request.args.get('host')
subprocess.run(f"ping -c 4 {host}", shell=True)

# Après — sécurisé
import re
host = request.args.get('host', '')
if not re.match(r'^[a-zA-Z0-9.\-]+$', host):
    abort(400, "Hôte invalide")
subprocess.run(["ping", "-c", "4", host], shell=False)
```

## Checklist de vérification post-patch
- [ ] Aucun appel `exec`/`system`/`subprocess.run(shell=True)` ne reçoit une chaîne construite à partir d'une entrée utilisateur.
- [ ] Les appels externes utilisent une API à tableau d'arguments (`execFile`, `proc_open`, `subprocess.run` avec liste) plutôt que l'interprétation shell.
- [ ] Une liste blanche stricte valide le format des valeurs transmises en argument de commande.
- [ ] Un test confirme qu'une entrée contenant des méta-caractères shell (`;`, `|`, `&&`, backtick) est rejetée avant l'exécution.
- [ ] Le processus externe s'exécute avec les privilèges minimaux nécessaires (utilisateur dédié, sandboxing).
