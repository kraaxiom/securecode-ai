# Remédiation — Header Injection

## Principe
Ne jamais construire un en-tête HTTP (nom ou valeur) par concaténation d'une entrée utilisateur brute. Rejeter ou supprimer les caractères de contrôle `\r` et `\n` avant toute écriture d'en-tête, et privilégier les API du framework qui encodent/valident automatiquement.

## PHP
```php
// Avant — vulnérable
$filename = $_GET['filename'];
header("Content-Disposition: attachment; filename=$filename");

// Après — sécurisé
$filename = $_GET['filename'];
$filename = preg_replace('/[\r\n]/', '', $filename);
$filename = basename($filename); // liste blanche de nom de fichier
header('Content-Disposition: attachment; filename="' . $filename . '"');
```

## JavaScript / Node.js (Express)
```js
// Avant — vulnérable
res.setHeader('X-Custom-User', req.query.username);

// Après — sécurisé
// Express rejette nativement \r\n depuis Node.js récent, mais on valide
// explicitement en défense en profondeur.
const username = String(req.query.username).replace(/[\r\n]/g, '');
if (!/^[\w.-]+$/.test(username)) {
  return res.status(400).send('Valeur invalide');
}
res.setHeader('X-Custom-User', username);
```

## Python (Flask)
```python
# Avant — vulnérable
filename = request.args.get("filename")
response.headers["Content-Disposition"] = f"attachment; filename={filename}"

# Après — sécurisé
import re
filename = request.args.get("filename", "")
filename = re.sub(r"[\r\n]", "", filename)
filename = os.path.basename(filename)
response.headers["Content-Disposition"] = f'attachment; filename="{filename}"'
```

## Checklist de vérification post-patch
- [ ] Aucune valeur d'en-tête n'est écrite sans suppression préalable des caractères `\r` et `\n`.
- [ ] Les en-têtes sont définis via les API du framework (pas d'écriture brute sur le flux socket).
- [ ] Un test confirme qu'une valeur contenant `\r\n` ne crée pas d'en-tête supplémentaire dans la réponse.
- [ ] Un test de non-régression confirme que les valeurs légitimes d'en-tête fonctionnent toujours.
- [ ] Les noms de fichiers ou identifiants réinjectés dans un en-tête sont validés par liste blanche de caractères.
