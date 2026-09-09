# Remédiation — ZIP Bomb (bombe de décompression)

## Principe
Vérifier la taille décompressée annoncée et le ratio de compression avant toute extraction complète, limiter le nombre/profondeur de fichiers imbriqués, et décompresser dans un environnement isolé avec quotas stricts de mémoire, disque, CPU et timeout, de préférence en streaming.

## PHP
```php
// Avant — vulnérable
$zip = new ZipArchive();
$zip->open($_FILES['file']['tmp_name']);
$zip->extractTo('uploads/extracted/'); // aucune limite de taille/ratio
$zip->close();

// Après — sécurisé
const MAX_UNCOMPRESSED_SIZE = 50 * 1024 * 1024; // 50 Mo
const MAX_COMPRESSION_RATIO = 100;
const MAX_ENTRIES = 1000;

$zip = new ZipArchive();
if ($zip->open($_FILES['file']['tmp_name']) !== true) {
    throw new RuntimeException('Archive invalide');
}

if ($zip->numFiles > MAX_ENTRIES) {
    $zip->close();
    throw new RuntimeException('Trop de fichiers dans l\'archive');
}

$totalUncompressed = 0;
for ($i = 0; $i < $zip->numFiles; $i++) {
    $stat = $zip->statIndex($i);
    $totalUncompressed += $stat['size'];

    $ratio = $stat['comp_size'] > 0 ? $stat['size'] / $stat['comp_size'] : 0;
    if ($ratio > MAX_COMPRESSION_RATIO) {
        $zip->close();
        throw new RuntimeException('Ratio de compression suspect détecté');
    }
    if (strpos($stat['name'], '..') !== false) { // bonus : anti path traversal
        $zip->close();
        throw new RuntimeException('Chemin d\'entrée invalide');
    }
}

if ($totalUncompressed > MAX_UNCOMPRESSED_SIZE) {
    $zip->close();
    throw new RuntimeException('Taille décompressée trop importante');
}

$zip->extractTo('/var/app_extract_sandbox/'); // environnement isolé, quotas disque/CPU
$zip->close();
```

## Node.js (yauzl — extraction en streaming avec limites)
```js
// Avant — vulnérable
const AdmZip = require('adm-zip');
new AdmZip(req.file.path).extractAllTo('public/uploads/extracted', true); // charge tout en mémoire

// Après — sécurisé
const yauzl = require('yauzl');

const MAX_UNCOMPRESSED_SIZE = 50 * 1024 * 1024;
const MAX_ENTRIES = 1000;

yauzl.open(req.file.path, { lazyEntries: true }, (err, zipfile) => {
  if (err) throw new Error('Archive invalide');

  let total = 0;
  let count = 0;

  zipfile.readEntry();
  zipfile.on('entry', (entry) => {
    count++;
    total += entry.uncompressedSize;

    if (count > MAX_ENTRIES || total > MAX_UNCOMPRESSED_SIZE) {
      zipfile.close();
      throw new Error('Archive dépasse les limites autorisées');
    }
    if (entry.fileName.includes('..')) {
      zipfile.close();
      throw new Error('Chemin d\'entrée invalide');
    }
    zipfile.readEntry(); // traitement incrémental, pas de chargement global
  });
});
```

## Python (zipfile avec vérification préalable)
```python
# Avant — vulnérable
import zipfile
with zipfile.ZipFile(uploaded_path) as z:
    z.extractall('static/uploads/extracted')  # aucune limite

# Après — sécurisé
import zipfile

MAX_UNCOMPRESSED_SIZE = 50 * 1024 * 1024
MAX_COMPRESSION_RATIO = 100
MAX_ENTRIES = 1000

with zipfile.ZipFile(uploaded_path) as z:
    infos = z.infolist()
    if len(infos) > MAX_ENTRIES:
        raise ValueError('Trop de fichiers dans l\'archive')

    total_uncompressed = sum(i.file_size for i in infos)
    if total_uncompressed > MAX_UNCOMPRESSED_SIZE:
        raise ValueError('Taille décompressée trop importante')

    for info in infos:
        if info.compress_size > 0:
            ratio = info.file_size / info.compress_size
            if ratio > MAX_COMPRESSION_RATIO:
                raise ValueError('Ratio de compression suspect détecté')
        if '..' in info.filename:
            raise ValueError('Chemin d\'entrée invalide')

    z.extractall('/var/app_extract_sandbox')  # environnement isolé, quotas disque/CPU/timeout
```

## Checklist de vérification post-patch
- [ ] La taille décompressée annoncée et le ratio de compression sont vérifiés avant toute extraction complète.
- [ ] Le nombre et la profondeur des fichiers/archives imbriquées sont limités.
- [ ] L'extraction s'exécute dans un environnement isolé avec quotas mémoire/disque/CPU et un timeout.
- [ ] Le traitement est incrémental (streaming) et s'arrête dès dépassement des limites, sans tout charger en mémoire.
- [ ] Un test confirme qu'une archive au ratio de compression anormalement élevé est rejetée avant extraction complète.
