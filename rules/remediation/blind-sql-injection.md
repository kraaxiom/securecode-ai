# Remédiation — Blind SQL Injection

## Principe
Remplacer toute concaténation de valeur utilisateur dans une requête SQL par une requête préparée avec paramètre lié, même lorsque le résultat de la requête n'est pas directement affiché à l'utilisateur. Uniformiser les temps de réponse et les messages retournés pour empêcher toute inférence de comportement.

## PHP (PDO)
```php
// Avant — vulnérable
$user = $_GET['user'];
$stmt = $pdo->query("SELECT 1 FROM users WHERE username = '$user' AND active = 1");
$exists = $stmt->fetch() !== false;

// Après — sécurisé
$stmt = $pdo->prepare("SELECT 1 FROM users WHERE username = :user AND active = 1");
$stmt->execute(['user' => $_GET['user']]);
$exists = $stmt->fetch() !== false;
```

## Node.js (mysql2)
```js
// Avant — vulnérable
const user = req.query.user;
connection.query(`SELECT 1 FROM users WHERE username = '${user}' AND active = 1`, (err, rows) => {
  res.json({ exists: rows.length > 0 });
});

// Après — sécurisé
connection.query('SELECT 1 FROM users WHERE username = ? AND active = 1', [req.query.user], (err, rows) => {
  res.json({ exists: rows.length > 0 });
});
```

## Python (psycopg2)
```python
# Avant — vulnérable
user = request.args.get("user")
cursor.execute(f"SELECT 1 FROM users WHERE username = '{user}' AND active = true")
exists = cursor.fetchone() is not None

# Après — sécurisé
cursor.execute("SELECT 1 FROM users WHERE username = %s AND active = true", [request.args.get("user")])
exists = cursor.fetchone() is not None
```

## Checklist de vérification post-patch
- [ ] Aucune concaténation de variable restante dans une requête SQL conditionnelle, même si le résultat n'est pas affiché directement.
- [ ] Un test de non-régression confirme que la recherche/authentification légitime fonctionne toujours.
- [ ] Un test confirme qu'une entrée contenant un caractère `'` ou un mot-clé SQL ne modifie ni le résultat ni le temps de réponse observable.
- [ ] Les temps de réponse entre cas "trouvé" et "non trouvé" sont comparables (pas de canal temporel exploitable).
- [ ] Aucun message d'erreur SQL détaillé n'est renvoyé au client en cas d'échec de la requête.
