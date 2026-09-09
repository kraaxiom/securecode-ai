# Remédiation — Bypass par falsification des magic bytes

## Principe
Ne jamais se fier à la seule signature binaire des premiers octets. Décoder complètement le fichier avec une bibliothèque d'image robuste et le ré-encoder systématiquement dans un nouveau fichier propre ; rejeter tout fichier dont le décodage complet échoue.

## PHP
```php
// Avant — vulnérable
$handle = fopen($_FILES['file']['tmp_name'], 'rb');
$header = fread($handle, 4);
fclose($handle);
if (bin2hex($header) === 'ffd8ffe0') { // signature JPEG uniquement
    move_uploaded_file($_FILES['file']['tmp_name'], 'uploads/' . $_FILES['file']['name']);
}

// Après — sécurisé
// getimagesize() décode réellement l'en-tête structurel de l'image, pas seulement 4 octets
$info = @getimagesize($_FILES['file']['tmp_name']);
if ($info === false || !in_array($info[2], [IMAGETYPE_JPEG, IMAGETYPE_PNG, IMAGETYPE_GIF], true)) {
    throw new RuntimeException('Fichier image invalide');
}

// Ré-encodage complet : détruit tout contenu superflu après la signature
$src = imagecreatefromstring(file_get_contents($_FILES['file']['tmp_name']));
if ($src === false) {
    throw new RuntimeException('Décodage image échoué');
}
$safeName = bin2hex(random_bytes(16)) . '.jpg';
imagejpeg($src, '/var/app_uploads/' . $safeName, 85);
imagedestroy($src);
```

## Node.js (sharp)
```js
// Avant — vulnérable
const buf = fs.readFileSync(req.file.path).slice(0, 4);
if (buf.toString('hex') === 'ffd8ffe0') {
  fs.renameSync(req.file.path, `public/uploads/${req.file.originalname}`);
}

// Après — sécurisé
const sharp = require('sharp');
const crypto = require('crypto');

try {
  const meta = await sharp(req.file.path).metadata(); // décodage structurel complet
  if (!['jpeg', 'png', 'gif'].includes(meta.format)) {
    throw new Error('Format non autorisé');
  }
  const safeName = crypto.randomBytes(16).toString('hex') + '.jpg';
  await sharp(req.file.path).jpeg().toFile(`/var/app_uploads/${safeName}`); // ré-encodage forcé
} catch {
  throw new Error('Fichier image invalide');
}
```

## Python (Pillow)
```python
# Avant — vulnérable
with open(uploaded_path, 'rb') as f:
    header = f.read(4)
if header == b'\xff\xd8\xff\xe0':
    file.save('static/uploads/' + file.filename)

# Après — sécurisé
from PIL import Image
import secrets

try:
    with Image.open(uploaded_path) as img:
        img.verify()  # décodage structurel, échoue sur fichier hybride/corrompu
    with Image.open(uploaded_path) as img:
        if img.format not in ('JPEG', 'PNG', 'GIF'):
            raise ValueError('Format non autorisé')
        safe_name = secrets.token_hex(16) + '.jpg'
        img.convert('RGB').save('/var/app_uploads/' + safe_name, format='JPEG')  # ré-encodage forcé
except Exception:
    abort(400, 'Fichier image invalide')
```

## Checklist de vérification post-patch
- [ ] La validation ne se base plus sur la seule lecture des premiers octets.
- [ ] Le fichier est décodé intégralement par une bibliothèque d'image dédiée, avec échec explicite sur contenu malformé.
- [ ] Le fichier stocké est le résultat d'un ré-encodage, jamais une copie octet-pour-octet du fichier reçu.
- [ ] Un test confirme qu'un fichier préfixé de magic bytes valides mais suivi de contenu non-image est rejeté.
