# Remédiation — Upload de web shell

## Principe
Séparer strictement le stockage des fichiers uploadés du répertoire servi comme code exécutable, désactiver l'exécution de scripts dans tout répertoire accessible en écriture par les utilisateurs, et appliquer une liste blanche stricte validée sur le contenu réel.

## PHP
```php
// Avant — vulnérable
// Le fichier atterrit directement dans le webroot, accessible et exécutable
move_uploaded_file($_FILES['file']['tmp_name'], __DIR__ . '/../public/uploads/' . $_FILES['file']['name']);

// Après — sécurisé
$allowedExt = ['jpg', 'jpeg', 'png', 'pdf'];
$ext = strtolower(pathinfo($_FILES['file']['name'], PATHINFO_EXTENSION));
$finfo = finfo_open(FILEINFO_MIME_TYPE);
$realMime = finfo_file($finfo, $_FILES['file']['tmp_name']);
$allowedMime = ['image/jpeg' => 'jpg', 'image/png' => 'png', 'application/pdf' => 'pdf'];

if (!in_array($ext, $allowedExt, true) || !isset($allowedMime[$realMime]) || $allowedMime[$realMime] !== $ext) {
    throw new RuntimeException('Type de fichier non autorisé');
}

$safeName = bin2hex(random_bytes(16)) . '.' . $ext;
// Stockage hors du webroot, dans un répertoire sans exécution de scripts (config serveur)
move_uploaded_file($_FILES['file']['tmp_name'], '/var/app_uploads/' . $safeName);
// Le fichier est ensuite servi via un endpoint dédié (lecture seule, Content-Disposition: attachment)
```

## Nginx — empêcher l'exécution de scripts dans le répertoire d'upload
```nginx
location /uploads/ {
    location ~ \.(php|php\d+|phtml|pl|py|jsp|asp|aspx|sh|cgi)$ {
        deny all;
        return 403;
    }
}
```

## Node.js (Express / multer)
```js
// Avant — vulnérable
const upload = multer({ dest: 'public/uploads/' }); // servi statiquement, potentiellement via un handler externe

// Après — sécurisé
const path = require('path');
const crypto = require('crypto');

const allowedMime = { 'image/jpeg': '.jpg', 'image/png': '.png', 'application/pdf': '.pdf' };

const storage = multer.diskStorage({
  destination: '/var/app_uploads', // hors de tout dossier servi par express.static
  filename: (req, file, cb) => {
    const ext = allowedMime[file.mimetype];
    if (!ext) return cb(new Error('Type de fichier non autorisé'));
    cb(null, crypto.randomBytes(16).toString('hex') + ext);
  },
});

const upload = multer({
  storage,
  fileFilter: (req, file, cb) => cb(null, Boolean(allowedMime[file.mimetype])),
  limits: { fileSize: 10 * 1024 * 1024 },
});
// Aucune route express.static ne pointe vers /var/app_uploads
```

## Python (Django)
```python
# Avant — vulnérable
# MEDIA_ROOT pointe vers un dossier servi par le serveur applicatif avec exécution possible
handle_uploaded_file(request.FILES['file'], settings.MEDIA_ROOT / request.FILES['file'].name)

# Après — sécurisé
import secrets
from pathlib import Path

ALLOWED_MIME = {'image/jpeg': '.jpg', 'image/png': '.png', 'application/pdf': '.pdf'}

f = request.FILES['file']
ext = ALLOWED_MIME.get(f.content_type)
if ext is None:
    raise ValidationError('Type de fichier non autorisé')

safe_name = secrets.token_hex(16) + ext
# UPLOAD_STORAGE_ROOT est hors de STATIC_ROOT/MEDIA servi publiquement, exécution désactivée
dest = Path(settings.UPLOAD_STORAGE_ROOT) / safe_name
with open(dest, 'wb+') as out:
    for chunk in f.chunks():
        out.write(chunk)
```

## Checklist de vérification post-patch
- [ ] Le répertoire de stockage des fichiers uploadés est physiquement séparé du webroot/de l'arborescence exécutable.
- [ ] L'exécution de scripts est explicitement désactivée par configuration serveur sur tout répertoire accessible en écriture par les utilisateurs.
- [ ] La validation combine extension, MIME détecté côté serveur et cohérence entre les deux.
- [ ] Les fichiers sont servis via un endpoint applicatif dédié (pas d'accès direct au chemin de stockage).
- [ ] Une revue périodique/détection d'anomalies existe sur les répertoires d'upload (nouveaux fichiers avec extension inattendue).
- [ ] Un test confirme qu'un script (`.php`, `.jsp`, `.aspx`) déposé n'est ni accessible ni exécutable après upload.
