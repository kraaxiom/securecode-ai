# Remédiation — Bypass par double extension

## Principe
Extraire explicitement la dernière extension réelle du nom de fichier (jamais par recherche de sous-chaîne), la comparer à une liste blanche exacte, rejeter tout nom contenant plusieurs extensions, et renommer systématiquement le fichier côté serveur avant stockage.

## PHP
```php
// Avant — vulnérable
$filename = $_FILES['file']['name'];
if (strpos($filename, '.jpg') !== false || strpos($filename, '.png') !== false) {
    move_uploaded_file($_FILES['file']['tmp_name'], 'uploads/' . $filename);
}

// Après — sécurisé
$allowedExt = ['jpg', 'jpeg', 'png', 'gif'];
$filename = $_FILES['file']['name'];

// Extraction de la SEULE extension finale
$ext = strtolower(pathinfo($filename, PATHINFO_EXTENSION));

// Rejet si le nom contient plus d'un point suivi d'une extension suspecte (ex: shell.php.jpg)
if (substr_count($filename, '.') > 1) {
    throw new RuntimeException('Nom de fichier non autorisé (extensions multiples)');
}
if (!in_array($ext, $allowedExt, true)) {
    throw new RuntimeException('Extension non autorisée');
}

$safeName = bin2hex(random_bytes(16)) . '.' . $ext;
move_uploaded_file($_FILES['file']['tmp_name'], '/var/app_uploads/' . $safeName);
```

## Node.js (Express)
```js
// Avant — vulnérable
const filename = req.file.originalname;
if (filename.includes('.jpg') || filename.includes('.png')) {
  fs.renameSync(req.file.path, `public/uploads/${filename}`);
}

// Après — sécurisé
const path = require('path');
const crypto = require('crypto');

const allowedExt = new Set(['.jpg', '.jpeg', '.png', '.gif']);
const filename = req.file.originalname;

// path.extname ne retourne que la dernière extension
const ext = path.extname(filename).toLowerCase();

if ((filename.match(/\./g) || []).length > 1) {
  throw new Error('Nom de fichier non autorisé (extensions multiples)');
}
if (!allowedExt.has(ext)) {
  throw new Error('Extension non autorisée');
}

const safeName = crypto.randomBytes(16).toString('hex') + ext;
fs.renameSync(req.file.path, path.join('/var/app_uploads', safeName));
```

## Python (Flask)
```python
# Avant — vulnérable
filename = file.filename
if '.jpg' in filename or '.png' in filename:
    file.save(os.path.join('static/uploads', filename))

# Après — sécurisé
import os
import secrets

ALLOWED_EXT = {'jpg', 'jpeg', 'png', 'gif'}
filename = file.filename

# Extraction de la seule extension finale
ext = filename.rsplit('.', 1)[-1].lower() if '.' in filename else ''

if filename.count('.') > 1:
    abort(400, 'Nom de fichier non autorisé (extensions multiples)')
if ext not in ALLOWED_EXT:
    abort(400, 'Extension non autorisée')

safe_name = secrets.token_hex(16) + '.' + ext
file.save(os.path.join('/var/app_uploads', safe_name))
```

## Checklist de vérification post-patch
- [ ] L'extension est extraite via une fonction dédiée (`pathinfo`, `path.extname`, `rsplit`) et non par recherche de sous-chaîne.
- [ ] Tout nom de fichier contenant plusieurs points/extensions est explicitement rejeté ou neutralisé par renommage.
- [ ] Le fichier stocké porte un nom généré côté serveur, jamais dérivé du nom client.
- [ ] La configuration du serveur web (Apache `AddHandler`/`AddType`, Nginx `location`) n'interprète pas un fichier sur la base d'une extension intermédiaire.
- [ ] Un test confirme qu'un fichier nommé `shell.php.jpg` est rejeté.
