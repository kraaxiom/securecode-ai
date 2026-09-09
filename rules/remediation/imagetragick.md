# Remédiation — ImageTragick (injection de commande via bibliothèque d'image)

## Principe
Maintenir la bibliothèque de traitement d'image à jour, restreindre strictement les coders/delegates actifs via une policy, isoler le traitement dans un environnement sans privilèges/réseau, et ne jamais propager de données contrôlées par l'utilisateur (nom de fichier, métadonnées) dans un appel de bibliothèque susceptible de les transmettre à un programme externe.

## PHP (Imagick)
```php
// Avant — vulnérable
$imagick = new Imagick();
$imagick->readImage($_FILES['file']['tmp_name']); // délègue selon le format détecté, coders non restreints
$imagick->writeImage('uploads/' . $_FILES['file']['name']);

// Après — sécurisé
// 1. policy.xml ImageMagick restrictive (hors code applicatif) désactivant MVG, MSL, URL, HTTPS, EPHEMERAL
// 2. Ne décoder que des formats raster attendus, dans un processus isolé (conteneur sans accès réseau)
$allowedFormats = ['JPEG', 'PNG', 'GIF'];

$imagick = new Imagick();
$imagick->pingImage($_FILES['file']['tmp_name']); // lecture des métadonnées seule, sans décodage complet
if (!in_array($imagick->getImageFormat(), $allowedFormats, true)) {
    throw new RuntimeException('Format image non autorisé');
}

$imagick->readImage($_FILES['file']['tmp_name']);
$imagick->stripImage(); // supprime métadonnées potentiellement dangereuses
$safeName = bin2hex(random_bytes(16)) . '.jpg';
$imagick->setImageFormat('jpeg'); // ré-encodage forcé, détruit tout contenu delegate
$imagick->writeImage('/var/app_uploads/' . $safeName);
```

## Node.js (sharp au lieu d'ImageMagick/GraphicsMagick)
```js
// Avant — vulnérable
const gm = require('gm');
gm(req.file.path).write(`public/uploads/${req.file.originalname}`, cb);

// Après — sécurisé
// Préférer une bibliothèque sans délégation à des programmes externes (sharp/libvips)
const sharp = require('sharp');
const crypto = require('crypto');

const safeName = crypto.randomBytes(16).toString('hex') + '.jpg';
await sharp(req.file.path)
  .jpeg() // ré-encodage forcé
  .toFile(`/var/app_uploads/${safeName}`);
```

## Python (Pillow avec ré-encodage forcé)
```python
# Avant — vulnérable
from wand.image import Image  # wrapper ImageMagick, delegates actifs par défaut
with Image(filename=uploaded_path) as img:
    img.save(filename='static/uploads/' + original_name)

# Après — sécurisé
from PIL import Image  # Pillow, pas de délégation shell externe
import secrets

ALLOWED_FORMATS = {'JPEG', 'PNG', 'GIF'}

with Image.open(uploaded_path) as img:
    img.verify()  # échoue sur fichier corrompu/malformé

with Image.open(uploaded_path) as img:
    if img.format not in ALLOWED_FORMATS:
        raise ValueError('Format image non autorisé')
    safe_name = secrets.token_hex(16) + '.jpg'
    img.convert('RGB').save('/var/app_uploads/' + safe_name, format='JPEG')
```

## Checklist de vérification post-patch
- [ ] La bibliothèque de traitement d'image (ImageMagick/GraphicsMagick ou wrapper) est à jour et patchée contre les CVE connues de cette classe.
- [ ] Une policy restrictive désactive les coders/delegates non nécessaires (MVG, MSL, URL, HTTPS, EPHEMERAL).
- [ ] Le traitement d'image s'exécute dans un environnement isolé (conteneur sans accès réseau/fichier sensible, utilisateur à privilèges minimaux).
- [ ] L'image est systématiquement ré-encodée (pas de copie octet-pour-octet du fichier reçu).
- [ ] Aucune valeur utilisateur (nom de fichier, métadonnées) n'est transmise sans neutralisation à un appel de bibliothèque.
