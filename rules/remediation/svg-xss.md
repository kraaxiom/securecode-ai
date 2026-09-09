# Remédiation — XSS via SVG

## Principe
Sanitiser le contenu SVG (suppression des balises `script`, gestionnaires d'événements, schémas `javascript:`) avant stockage ou affichage. Servir les fichiers uploadés depuis une origine séparée sans cookies, et forcer le téléchargement plutôt que l'affichage inline lorsque possible.

## PHP (upload + sanitisation avant stockage)
```php
// Avant — vulnérable — SVG stocké et servi tel quel
move_uploaded_file($_FILES['avatar']['tmp_name'], $uploadDir . $filename);
header('Content-Type: image/svg+xml');
readfile($uploadDir . $filename);

// Après — sécurisé — sanitisation XML avant stockage + service en téléchargement forcé
$svgContent = file_get_contents($_FILES['avatar']['tmp_name']);
$purifier = new \enshrined\svgSanitize\Sanitizer();
$clean = $purifier->sanitize($svgContent);
file_put_contents($uploadDir . $filename, $clean);

// Service depuis un sous-domaine dédié, en téléchargement forcé :
header('Content-Type: image/svg+xml');
header('Content-Disposition: attachment; filename="' . basename($filename) . '"');
readfile($uploadDir . $filename);
```

## JavaScript / Node.js (upload + affichage inline)
```js
// Avant — vulnérable — SVG utilisateur inséré inline dans le DOM sans sanitisation
fetch('/svg/' + id)
  .then((r) => r.text())
  .then((svgText) => { container.innerHTML = svgText; });

// Après — sécurisé — sanitisation dédiée SVG/HTML avant insertion inline
import DOMPurify from 'dompurify';

fetch('/svg/' + id)
  .then((r) => r.text())
  .then((svgText) => {
    container.innerHTML = DOMPurify.sanitize(svgText, { USE_PROFILES: { svg: true, svgFilters: true } });
  });
```

```js
// Avant — vulnérable — Express sert les SVG uploadés depuis la même origine sans sanitisation
app.use('/uploads', express.static('uploads'));

// Après — sécurisé — sanitisation à l'upload via une bibliothèque dédiée
const { sanitize } = require('svg-sanitizer'); // ou DOMPurify côté serveur avec jsdom
app.post('/upload', upload.single('file'), (req, res) => {
  const clean = sanitize(req.file.buffer.toString('utf8'));
  fs.writeFileSync(path.join('uploads', req.file.filename), clean);
  res.sendStatus(204);
});
```

## Python (Flask — sanitisation avant stockage)
```python
# Avant — vulnérable
file.save(os.path.join(upload_dir, filename))

# Après — sécurisé — sanitisation XML via bibliothèque dédiée avant stockage
from lxml import etree

def sanitize_svg(svg_bytes):
    parser = etree.XMLParser(resolve_entities=False, no_network=True)
    tree = etree.fromstring(svg_bytes, parser)
    for tag in ('script',):
        for el in tree.iter(tag):
            el.getparent().remove(el)
    for el in tree.iter():
        for attr in list(el.attrib):
            if attr.startswith('on') or 'javascript:' in el.attrib[attr].lower():
                del el.attrib[attr]
    return etree.tostring(tree)

clean = sanitize_svg(file.read())
with open(os.path.join(upload_dir, filename), 'wb') as f:
    f.write(clean)
```

## Checklist de vérification post-patch
- [ ] Tout fichier SVG uploadé est sanitisé (suppression de `script`, gestionnaires `on*`, schémas `javascript:`) avant stockage.
- [ ] Les fichiers uploadés (SVG inclus) sont servis depuis une origine séparée, sans cookies de session.
- [ ] L'affichage inline (`innerHTML`) de SVG utilisateur passe par un sanitiseur dédié plutôt qu'une insertion brute.
- [ ] Les fichiers non indispensables au rendu direct sont servis avec `Content-Disposition: attachment`.
- [ ] Un test de non-régression confirme qu'un SVG légitime (icône, logo) s'affiche toujours correctement après sanitisation.
