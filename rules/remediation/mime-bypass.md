# Remédiation — Bypass de validation MIME type

## Principe
Ne jamais utiliser le `Content-Type` ou l'extension déclarée par le client comme critère de validation. Déterminer le type réel du fichier côté serveur à partir de son contenu, et rejeter toute incohérence entre extension, MIME déclaré et type détecté.

## PHP
```php
// Avant — vulnérable
$mime = $_FILES['file']['type']; // entièrement contrôlé par le client
if (in_array($mime, ['image/jpeg', 'image/png'], true)) {
    move_uploaded_file($_FILES['file']['tmp_name'], 'uploads/' . $_FILES['file']['name']);
}

// Après — sécurisé
$finfo = finfo_open(FILEINFO_MIME_TYPE);
$realMime = finfo_file($finfo, $_FILES['file']['tmp_name']); // détection par contenu, côté serveur
finfo_close($finfo);

$allowedMime = ['image/jpeg' => 'jpg', 'image/png' => 'png', 'image/gif' => 'gif'];
if (!isset($allowedMime[$realMime])) {
    throw new RuntimeException('Type de fichier non autorisé');
}

$safeName = bin2hex(random_bytes(16)) . '.' . $allowedMime[$realMime];
move_uploaded_file($_FILES['file']['tmp_name'], '/var/app_uploads/' . $safeName);
```

## Node.js (Express / multer + file-type)
```js
// Avant — vulnérable
const upload = multer({
  fileFilter: (req, file, cb) => cb(null, ['image/jpeg', 'image/png'].includes(file.mimetype)), // valeur client
});

// Après — sécurisé
const { fileTypeFromFile } = require('file-type'); // détection par signature/contenu
const crypto = require('crypto');

app.post('/upload', multer({ dest: '/tmp' }).single('file'), async (req, res) => {
  const detected = await fileTypeFromFile(req.file.path);
  const allowedMime = { 'image/jpeg': '.jpg', 'image/png': '.png', 'image/gif': '.gif' };

  if (!detected || !allowedMime[detected.mime]) {
    return res.status(400).send('Type de fichier non autorisé');
  }

  const safeName = crypto.randomBytes(16).toString('hex') + allowedMime[detected.mime];
  fs.renameSync(req.file.path, `/var/app_uploads/${safeName}`);
  res.send('ok');
});
```

## Python (Flask + python-magic)
```python
# Avant — vulnérable
mime = file.mimetype  # valeur déclarée par le client, falsifiable
if mime in ('image/jpeg', 'image/png'):
    file.save(os.path.join('static/uploads', file.filename))

# Après — sécurisé
import magic  # python-magic, détection par contenu
import secrets

ALLOWED_MIME = {'image/jpeg': '.jpg', 'image/png': '.png', 'image/gif': '.gif'}

file.save('/tmp/upload_tmp')
real_mime = magic.from_file('/tmp/upload_tmp', mime=True)
ext = ALLOWED_MIME.get(real_mime)
if ext is None:
    abort(400, 'Type de fichier non autorisé')

safe_name = secrets.token_hex(16) + ext
os.rename('/tmp/upload_tmp', os.path.join('/var/app_uploads', safe_name))
```

## Checklist de vérification post-patch
- [ ] Le `Content-Type`/MIME déclaré par le client n'est plus utilisé seul comme critère de validation.
- [ ] Le type réel est déterminé côté serveur à partir du contenu (`finfo_file`, `file-type`, `python-magic` ou équivalent).
- [ ] Toute incohérence entre extension déclarée, MIME déclaré et type détecté entraîne un rejet.
- [ ] Un test confirme qu'un fichier renommé avec un `Content-Type` falsifié est bien rejeté.
