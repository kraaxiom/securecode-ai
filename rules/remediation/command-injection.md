# Remédiation — Command Injection

## Principe
Éviter tout passage par un interpréteur shell lors de l'exécution d'un processus externe. Utiliser les API d'exécution qui acceptent un tableau d'arguments distincts (pas de concaténation de chaîne), et valider les entrées avec une liste blanche stricte.

## PHP
```php
// Avant — vulnérable
$host = $_GET['host'];
system("ping -c 3 " . $host);

// Après — sécurisé
$host = $_GET['host'];
if (!filter_var($host, FILTER_VALIDATE_IP)) {
    throw new InvalidArgumentException('Hôte invalide');
}
$process = new Symfony\Component\Process\Process(['ping', '-c', '3', $host]);
$process->run();
```

## Node.js
```js
// Avant — vulnérable
const { exec } = require('child_process');
const host = req.query.host;
exec(`ping -c 3 ${host}`, (err, stdout) => res.send(stdout));

// Après — sécurisé
const { execFile } = require('child_process');
const net = require('net');
const host = req.query.host;
if (!net.isIP(host)) {
  return res.status(400).send('Hôte invalide');
}
execFile('ping', ['-c', '3', host], (err, stdout) => res.send(stdout));
```

## Python
```python
# Avant — vulnérable
import subprocess
host = request.args.get("host")
subprocess.call(f"ping -c 3 {host}", shell=True)

# Après — sécurisé
import ipaddress
import subprocess

host = request.args.get("host")
ipaddress.ip_address(host)  # lève ValueError si invalide
subprocess.run(["ping", "-c", "3", host], shell=False)
```

## Checklist de vérification post-patch
- [ ] Aucun appel `system`/`exec`/`popen`/`subprocess` avec `shell=True` ou concaténation de chaîne restant dans le fichier corrigé.
- [ ] Tous les arguments de commande sont passés en tableau distinct, jamais via une chaîne unique interprétée par un shell.
- [ ] Une validation stricte (liste blanche, format IP/nom de fichier) est appliquée avant tout passage à l'exécutable.
- [ ] Le processus s'exécute avec le moindre privilège possible (utilisateur dédié, sandbox/conteneur).
- [ ] Un test confirme qu'une entrée contenant des métacaractères shell est rejetée ou traitée comme une donnée inerte, pas comme du contrôle.
