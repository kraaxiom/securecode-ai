# Remédiation — Upload de fichier PHP exécutable

## Principe
Ne jamais faire confiance au nom de fichier ni au type MIME déclaré par le client. Stocker les fichiers téléversés hors du webroot (ou dans un répertoire où l'exécution de scripts est désactivée), valider le contenu réel via une liste blanche stricte, et générer un nom de fichier aléatoire côté serveur.

## PHP
```php
// Avant — vulnérable
$dest = __DIR__ . '/uploads/' . $_FILES['file']['name'];
move_uploaded_file($_FILES['file']['tmp_name'], $dest);

// Après — sécurisé
$allowedExt = ['jpg', 'jpeg', 'png', 'gif'];
$ext = strtolower(pathinfo($_FILES['file']['name'], PATHINFO_EXTENSION));
$finfo = finfo_open(FILEINFO_MIME_TYPE);
$realMime = finfo_file($finfo, $_FILES['file']['tmp_name']);
$allowedMime = ['image/jpeg' => 'jpg', 'image/png' => 'png', 'image/gif' => 'gif'];

if (!in_array($ext, $allowedExt, true) || !isset($allowedMime[$realMime])) {
    throw new RuntimeException('Type de fichier non autorisé');
}

$safeName = bin2hex(random_bytes(16)) . '.' . $allowedMime[$realMime];
$destDir = '/var/app_uploads'; // hors du webroot
move_uploaded_file($_FILES['file']['tmp_name'], $destDir . '/' . $safeName);
```

## Node.js (Express / multer)
```js
// Avant — vulnérable
const upload = multer({ dest: 'public/uploads/' });
app.post('/upload', upload.single('file'), (req, res) => {
  res.send('ok');
});

// Après — sécurisé
const path = require('path');
const crypto = require('crypto');

const allowedMime = { 'image/jpeg': '.jpg', 'image/png': '.png', 'image/gif': '.gif' };

const storage = multer.diskStorage({
  destination: '/var/app_uploads', // hors du webroot
  filename: (req, file, cb) => {
    const ext = allowedMime[file.mimetype];
    if (!ext) return cb(new Error('Type de fichier non autorisé'));
    cb(null, crypto.randomBytes(16).toString('hex') + ext);
  },
});

const upload = multer({
  storage,
  fileFilter: (req, file, cb) => cb(null, Boolean(allowedMime[file.mimetype])),
  limits: { fileSize: 5 * 1024 * 1024 },
});

app.post('/upload', upload.single('file'), (req, res) => res.send('ok'));
```

## Python (Flask/Django)
```python
# Avant — vulnérable
file = request.files['file']
file.save(os.path.join('static/uploads', file.filename))

# Après — sécurisé
import os
import secrets
from werkzeug.utils import secure_filename

ALLOWED_MIME = {'image/jpeg': '.jpg', 'image/png': '.png', 'image/gif': '.gif'}

file = request.files['file']
mime = file.mimetype  # à corroborer avec une détection par contenu (python-magic)
ext = ALLOWED_MIME.get(mime)
if ext is None:
    abort(400, 'Type de fichier non autorisé')

safe_name = secrets.token_hex(16) + ext
dest_dir = '/var/app_uploads'  # hors du webroot
file.save(os.path.join(dest_dir, safe_name))
```

## Checklist de vérification post-patch
- [ ] Le répertoire de destination n'est plus dans l'arborescence servie par le serveur web (ou l'exécution de scripts y est explicitement désactivée).
- [ ] Le nom du fichier stocké est généré côté serveur, jamais dérivé du nom fourni par le client.
- [ ] La validation combine extension, type MIME détecté côté serveur et taille maximale.
- [ ] Un test confirme qu'un fichier `.php` déguisé (renommé, mauvais MIME) est bien rejeté.
