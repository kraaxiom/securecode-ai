# Remédiation — NoSQL Injection

## Principe
Ne jamais transmettre un corps de requête HTTP brut (`req.body`) comme filtre de requête à une base NoSQL. Valider strictement le schéma et le type de chaque champ attendu (rejeter tout objet/opérateur là où une valeur scalaire est attendue) et éviter les opérateurs d'évaluation de code côté serveur (`$where`) avec des données non fiables.

## PHP (MongoDB driver)
```php
// Avant — vulnérable
$username = $_POST['username'];
$password = $_POST['password'];
$user = $collection->findOne(['username' => $username, 'password' => $password]);

// Après — sécurisé
$username = (string) $_POST['username'];
$password = (string) $_POST['password'];
if (!is_string($_POST['username']) || !is_string($_POST['password'])) {
    throw new InvalidArgumentException('Format invalide');
}
$user = $collection->findOne(['username' => $username, 'password' => $password]);
```

## JavaScript / Node.js (MongoDB / Mongoose)
```js
// Avant — vulnérable
const { username, password } = req.body;
const user = await User.findOne({ username, password });

// Après — sécurisé
const { username, password } = req.body;
if (typeof username !== 'string' || typeof password !== 'string') {
  return res.status(400).json({ error: 'Format invalide' });
}
// Bibliothèque de sanitation dédiée (ex: express-mongo-sanitize) en middleware global
const user = await User.findOne({ username, password });
```

## Python (PyMongo)
```python
# Avant — vulnérable
username = request.json['username']
password = request.json['password']
user = collection.find_one({"username": username, "password": password})

# Après — sécurisé
username = request.json.get('username')
password = request.json.get('password')
if not isinstance(username, str) or not isinstance(password, str):
    abort(400, "Format invalide")
user = collection.find_one({"username": username, "password": password})
```

## Checklist de vérification post-patch
- [ ] Chaque champ issu de l'utilisateur est validé par type/schéma avant d'être inséré dans une requête NoSQL.
- [ ] Aucun objet ou opérateur (`$ne`, `$gt`, `$where`, `$regex`) ne peut être injecté là où une valeur scalaire est attendue.
- [ ] L'opérateur `$where` (évaluation de code) n'est jamais utilisé avec des données utilisateur.
- [ ] Un middleware ou une validation de schéma (ex: JSON Schema, Zod, Joi) rejette les payloads non conformes en amont de la requête.
- [ ] Un test confirme qu'un payload objet (`{"$ne": null}`) envoyé à la place d'une chaîne est rejeté avant d'atteindre la base.
