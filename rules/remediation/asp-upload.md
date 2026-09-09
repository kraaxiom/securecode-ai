# Remédiation — Upload de fichier ASP/ASP.NET exécutable

## Principe
Stocker les fichiers en dehors de l'arborescence servie par IIS ou dans un dossier dont le `web.config` retire tous les handlers d'exécution. Valider le contenu réel du fichier via une liste blanche stricte plutôt que le `ContentType` déclaré, et générer un nom de fichier aléatoire côté serveur.

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
const upload = multer({ dest: 'wwwroot/uploads/' });

// Après — sécurisé
const allowedMime = { 'image/jpeg': '.jpg', 'image/png': '.png', 'image/gif': '.gif' };

const storage = multer.diskStorage({
  destination: '/var/app_uploads', // hors du répertoire servi par IIS/reverse proxy
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
file.save(os.path.join('wwwroot/uploads', file.filename))

# Après — sécurisé
import secrets

ALLOWED_MIME = {'image/jpeg': '.jpg', 'image/png': '.png', 'image/gif': '.gif'}
file = request.files['file']
ext = ALLOWED_MIME.get(file.mimetype)
if ext is None:
    abort(400, 'Type de fichier non autorisé')
file.save(os.path.join('/var/app_uploads', secrets.token_hex(16) + ext))
```

## .NET (ASP.NET Core)
```csharp
// Avant — vulnérable
[HttpPost("upload")]
public async Task<IActionResult> Upload(IFormFile file)
{
    var path = Path.Combine(_env.WebRootPath, "uploads", file.FileName);
    using var stream = new FileStream(path, FileMode.Create);
    await file.CopyToAsync(stream);
    return Ok();
}

// Après — sécurisé
private static readonly Dictionary<string, byte[]> AllowedSignatures = new()
{
    { ".jpg", new byte[] { 0xFF, 0xD8, 0xFF } },
    { ".png", new byte[] { 0x89, 0x50, 0x4E, 0x47 } },
};

[HttpPost("upload")]
public async Task<IActionResult> Upload(IFormFile file)
{
    var ext = Path.GetExtension(file.FileName).ToLowerInvariant();
    if (!AllowedSignatures.TryGetValue(ext, out var signature))
        return BadRequest("Type de fichier non autorisé");

    await using var stream = file.OpenReadStream();
    var header = new byte[signature.Length];
    await stream.ReadAsync(header, 0, signature.Length);
    if (!header.SequenceEqual(signature))
        return BadRequest("Signature de fichier invalide");

    var safeName = $"{Guid.NewGuid():N}{ext}";
    // Répertoire hors de wwwroot, ou protégé par un web.config retirant tous les handlers
    var destPath = Path.Combine(_uploadsRootOutsideWebRoot, safeName);
    stream.Seek(0, SeekOrigin.Begin);
    await using var outStream = new FileStream(destPath, FileMode.Create);
    await stream.CopyToAsync(outStream);
    return Ok();
}
```

## Checklist de vérification post-patch
- [ ] Le dossier d'upload n'est plus servi directement par IIS/Kestrel, ou un `web.config` dédié y retire tous les handlers d'exécution.
- [ ] La validation repose sur le contenu réel (signature binaire) et non sur `ContentType` ni sur l'extension déclarée seule.
- [ ] Le nom de fichier stocké est généré côté serveur (GUID), jamais dérivé du nom client.
- [ ] Les mappings d'extensions exécutables hérités (`.asa`, `.cer`, `.cdx`) sont retirés de la configuration IIS du dossier d'upload.
