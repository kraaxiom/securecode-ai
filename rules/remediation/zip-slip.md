# Remédiation — Zip Slip

## Principe
Pour chaque entrée d'une archive (ZIP, TAR...), résoudre le chemin de destination en forme canonique et vérifier qu'il reste strictement sous le répertoire d'extraction avant toute écriture. Rejeter toute entrée dont le nom contient des séquences `..`, un chemin absolu, ou qui pointe vers un lien symbolique.

## PHP
```php
// Avant — vulnérable
$zip = new ZipArchive();
$zip->open($archivePath);
$zip->extractTo('/var/app/uploads/extracted/');

// Après — sécurisé
$destDir = realpath('/var/app/uploads/extracted');
$zip = new ZipArchive();
$zip->open($archivePath);
for ($i = 0; $i < $zip->numFiles; $i++) {
    $name = $zip->getNameIndex($i);
    $target = $destDir . DIRECTORY_SEPARATOR . $name;
    $resolved = str_replace('\\', '/', $target);
    if (strpos($resolved, $destDir) !== 0) {
        throw new RuntimeException("Entrée d'archive suspecte: $name");
    }
    $zip->extractTo($destDir, [$name]);
}
```

## Node.js
```js
// Avant — vulnérable
const zip = new AdmZip(archivePath);
zip.extractAllTo('/var/app/uploads/extracted', true);

// Après — sécurisé
const path = require('path');
const destDir = path.resolve('/var/app/uploads/extracted');
const zip = new AdmZip(archivePath);
for (const entry of zip.getEntries()) {
  const target = path.resolve(destDir, entry.entryName);
  if (!target.startsWith(destDir + path.sep)) {
    throw new Error(`Entrée d'archive suspecte: ${entry.entryName}`);
  }
  zip.extractEntryTo(entry, destDir, false, true);
}
```

## Python
```python
# Avant — vulnérable
with zipfile.ZipFile(archive_path) as z:
    z.extractall('/var/app/uploads/extracted')

# Après — sécurisé
import os
import zipfile

dest_dir = os.path.realpath('/var/app/uploads/extracted')

with zipfile.ZipFile(archive_path) as z:
    for member in z.namelist():
        target = os.path.realpath(os.path.join(dest_dir, member))
        if not target.startswith(dest_dir + os.sep):
            raise ValueError(f"Entrée d'archive suspecte: {member}")
        z.extract(member, dest_dir)
```

## Java (référence — bibliothèque courante affectée par zip-slip)
```java
// Après — sécurisé
Path destDir = Paths.get("/var/app/uploads/extracted").toRealPath();
Path target = destDir.resolve(entry.getName()).normalize();
if (!target.startsWith(destDir)) {
    throw new IOException("Entrée d'archive suspecte: " + entry.getName());
}
```

## Checklist de vérification post-patch
- [ ] Chaque entrée d'archive a son chemin de destination résolu en forme canonique avant écriture.
- [ ] Le chemin résolu est vérifié comme sous-chemin strict du répertoire d'extraction (avec séparateur final, pas simple préfixe de chaîne).
- [ ] Les entrées de type lien symbolique dans l'archive sont détectées et rejetées.
- [ ] Une limite de taille totale/nombre d'entrées est appliquée pour éviter les attaques par zip bomb en complément.
