# Remédiation — Upload de fichier JSP exécutable

## Principe
Stocker les fichiers téléversés en dehors du répertoire `webapps` du conteneur servlet, valider le contenu réel via une liste blanche stricte, générer un nom de fichier aléatoire, et restreindre/désactiver l'accès au Manager App Tomcat en production.

## PHP
```php
// Avant — vulnérable
$dest = __DIR__ . '/uploads/' . $_FILES['file']['name'];
move_uploaded_file($_FILES['file']['tmp_name'], $dest);

// Après — sécurisé
$allowedMime = ['image/jpeg' => 'jpg', 'image/png' => 'png', 'image/gif' => 'gif'];
$finfo = finfo_open(FILEINFO_MIME_TYPE);
$realMime = finfo_file($finfo, $_FILES['file']['tmp_name']);

if (!isset($allowedMime[$realMime])) {
    throw new RuntimeException('Type de fichier non autorisé');
}

$safeName = bin2hex(random_bytes(16)) . '.' . $allowedMime[$realMime];
move_uploaded_file($_FILES['file']['tmp_name'], '/var/app_uploads/' . $safeName);
```

## Node.js (Express / multer)
```js
// Avant — vulnérable
const upload = multer({ dest: 'webapps/ROOT/uploads/' });

// Après — sécurisé
const allowedMime = { 'image/jpeg': '.jpg', 'image/png': '.png', 'image/gif': '.gif' };

const storage = multer.diskStorage({
  destination: '/var/app_uploads', // hors de tout répertoire déployé/servi
  filename: (req, file, cb) => {
    const ext = allowedMime[file.mimetype];
    if (!ext) return cb(new Error('Type de fichier non autorisé'));
    cb(null, require('crypto').randomBytes(16).toString('hex') + ext);
  },
});
const upload = multer({ storage, fileFilter: (req, file, cb) => cb(null, Boolean(allowedMime[file.mimetype])) });
```

## Python (Flask/Django)
```python
# Avant — vulnérable
file = request.files['file']
file.save(os.path.join('webapps/ROOT/uploads', file.filename))

# Après — sécurisé
import secrets

ALLOWED_MIME = {'image/jpeg': '.jpg', 'image/png': '.png', 'image/gif': '.gif'}
file = request.files['file']
ext = ALLOWED_MIME.get(file.mimetype)
if ext is None:
    abort(400, 'Type de fichier non autorisé')
file.save(os.path.join('/var/app_uploads', secrets.token_hex(16) + ext))
```

## Java (Servlet/Spring)
```java
// Avant — vulnérable
@PostMapping("/upload")
public ResponseEntity<?> upload(@RequestParam("file") MultipartFile file) throws IOException {
    Path dest = Paths.get(request.getServletContext().getRealPath("/uploads"), file.getOriginalFilename());
    file.transferTo(dest);
    return ResponseEntity.ok().build();
}

// Après — sécurisé
private static final Map<String, String> ALLOWED_MIME = Map.of(
    "image/jpeg", ".jpg",
    "image/png", ".png"
);
private static final Path UPLOAD_ROOT = Paths.get("/var/app_uploads"); // hors de webapps

@PostMapping("/upload")
public ResponseEntity<?> upload(@RequestParam("file") MultipartFile file) throws IOException {
    Tika tika = new Tika();
    String realMime = tika.detect(file.getInputStream());
    String ext = ALLOWED_MIME.get(realMime);
    if (ext == null) {
        return ResponseEntity.badRequest().body("Type de fichier non autorisé");
    }
    String safeName = UUID.randomUUID() + ext;
    file.transferTo(UPLOAD_ROOT.resolve(safeName));
    return ResponseEntity.ok().build();
}
```

## Checklist de vérification post-patch
- [ ] Le répertoire de destination n'est plus sous `webapps/` du conteneur servlet.
- [ ] Le contenu réel du fichier est vérifié (détection de type via bibliothèque dédiée), pas le `Content-Type` déclaré.
- [ ] Le nom du fichier stocké est généré côté serveur (UUID).
- [ ] Le Manager App Tomcat est désactivé ou restreint par authentification forte et filtrage IP en production.
