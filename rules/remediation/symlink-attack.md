# Remédiation — Symlink Attack

## Principe
Éviter toute fenêtre TOCTOU entre la vérification d'un fichier (`stat`/`lstat`) et son utilisation. Utiliser des opérations atomiques d'ouverture exclusive (`O_EXCL`/`O_NOFOLLOW` ou équivalents), vérifier explicitement qu'un chemin n'est pas un lien symbolique avant une opération sensible, et éviter les répertoires partagés en écriture par des tiers non fiables.

## PHP
```php
// Avant — vulnérable (TOCTOU : vérification puis écriture séparées)
$path = '/tmp/shared/' . $filename;
if (!file_exists($path)) {
    file_put_contents($path, $data);
}

// Après — sécurisé (création atomique exclusive, refus si lien symbolique)
$path = '/var/app/private-tmp/' . basename($filename);
$fp = fopen($path, 'x'); // échoue si le fichier existe déjà (atomique)
if ($fp === false) {
    throw new RuntimeException('Impossible de créer le fichier de façon sécurisée');
}
if (is_link($path)) {
    fclose($fp);
    unlink($path);
    throw new RuntimeException('Lien symbolique détecté');
}
fwrite($fp, $data);
fclose($fp);
```

## Node.js
```js
// Avant — vulnérable
if (!fs.existsSync(targetPath)) {
  fs.writeFileSync(targetPath, data);
}

// Après — sécurisé
const fd = fs.openSync(targetPath, fs.constants.O_CREAT | fs.constants.O_EXCL | fs.constants.O_WRONLY);
try {
  const stats = fs.lstatSync(targetPath);
  if (stats.isSymbolicLink()) {
    throw new Error('Lien symbolique détecté');
  }
  fs.writeSync(fd, data);
} finally {
  fs.closeSync(fd);
}
```

## Python
```python
# Avant — vulnérable
if not os.path.exists(target_path):
    with open(target_path, 'w') as f:
        f.write(data)

# Après — sécurisé
import os

fd = os.open(target_path, os.O_CREAT | os.O_EXCL | os.O_WRONLY | os.O_NOFOLLOW, 0o600)
try:
    with os.fdopen(fd, 'w') as f:
        f.write(data)
except OSError as e:
    raise RuntimeError("Création sécurisée impossible") from e
```

## Checklist de vérification post-patch
- [ ] Toute création de fichier dans un répertoire partagé utilise une ouverture atomique exclusive (`O_EXCL`), pas un `stat` suivi d'une écriture séparée.
- [ ] `O_NOFOLLOW` (ou vérification explicite `is_link`/`isSymbolicLink`) est utilisé pour refuser d'opérer sur un lien symbolique.
- [ ] Les répertoires temporaires/partagés utilisés par le service ne sont pas accessibles en écriture par des utilisateurs/processus non fiables.
- [ ] Les permissions du fichier créé sont restrictives dès la création (pas de fenêtre en mode monde-lisible/inscriptible).
