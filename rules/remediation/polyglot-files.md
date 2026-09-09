# Remédiation — Fichiers polyglots

## Principe
Toujours re-encoder/transformer les images après upload pour détruire tout contenu polyglot superflu, ne jamais exécuter/inclure dynamiquement un fichier issu d'un répertoire d'upload utilisateur, et neutraliser strictement scripts et entités externes dans les SVG.

## PHP
```php
// Avant — vulnérable
$dest = 'uploads/' . $_FILES['file']['name'];
move_uploaded_file($_FILES['file']['tmp_name'], $dest);
// ailleurs dans l'app : include($dest); // exécute le fichier "image" comme PHP si polyglot

// Après — sécurisé
$info = @getimagesize($_FILES['file']['tmp_name']);
if ($info === false) {
    throw new RuntimeException('Fichier image invalide');
}

// Ré-encodage complet : élimine tout octet superflu (ex: code PHP en fin de fichier GIF/JPEG)
$src = imagecreatefromstring(file_get_contents($_FILES['file']['tmp_name']));
$safeName = bin2hex(random_bytes(16)) . '.jpg';
imagejpeg($src, '/var/app_uploads/' . $safeName, 85);
imagedestroy($src);

// Ne jamais include()/require() un fichier venant du répertoire d'upload
```

## Node.js (SVG — neutralisation via DOMPurify côté serveur)
```js
// Avant — vulnérable
fs.writeFileSync(`public/uploads/${req.file.originalname}`, req.file.buffer); // SVG brut, script possible

// Après — sécurisé
const createDOMPurify = require('dompurify');
const { JSDOM } = require('jsdom');
const DOMPurify = createDOMPurify(new JSDOM('').window);
const crypto = require('crypto');

const svgContent = req.file.buffer.toString('utf8');
const clean = DOMPurify.sanitize(svgContent, {
  USE_PROFILES: { svg: true, svgFilters: true },
  FORBID_TAGS: ['script'],
  FORBID_ATTR: ['onload', 'onerror'],
});

const safeName = crypto.randomBytes(16).toString('hex') + '.svg';
fs.writeFileSync(`/var/app_uploads/${safeName}`, clean);
```

## Python (Flask — ré-encodage raster forcé au lieu de conserver le SVG/format brut)
```python
# Avant — vulnérable
file.save('static/uploads/' + file.filename)  # fichier stocké tel quel, potentiellement polyglot

# Après — sécurisé
from PIL import Image
import secrets

with Image.open(file.stream) as img:
    img.verify()

file.stream.seek(0)
with Image.open(file.stream) as img:
    safe_name = secrets.token_hex(16) + '.jpg'
    # Conversion vers un format raster détruit tout contenu polyglot (script, archive imbriquée)
    img.convert('RGB').save('/var/app_uploads/' + safe_name, format='JPEG')
```

## Checklist de vérification post-patch
- [ ] Chaque image uploadée est ré-encodée (pas de copie octet-pour-octet) avant stockage définitif.
- [ ] Aucune fonction d'inclusion ou d'exécution dynamique (`include`, `require`, désérialisation) n'est appliquée à un fichier issu d'un répertoire d'upload utilisateur.
- [ ] Les fichiers SVG sont soit convertis en raster, soit passés par un parseur strict neutralisant scripts et entités externes (XXE).
- [ ] Un test confirme qu'un fichier image contenant un fragment de script en fin de fichier est neutralisé après traitement.
