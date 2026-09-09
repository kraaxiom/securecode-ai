# Remédiation — Boolean-based SQL Injection

## Principe
Remplacer toute comparaison SQL construite par concaténation d'une entrée utilisateur par une requête préparée avec paramètre lié. Valider strictement le type attendu (ex: identifiant numérique) avant toute utilisation dans une clause `WHERE`.

## PHP (PDO)
```php
// Avant — vulnérable
$name = $_POST['name'];
$result = $pdo->query("SELECT * FROM products WHERE name = '$name'");

// Après — sécurisé
$stmt = $pdo->prepare("SELECT * FROM products WHERE name = :name");
$stmt->execute(['name' => $_POST['name']]);
$result = $stmt->fetchAll();
```

## Node.js (mysql2)
```js
// Avant — vulnérable
const name = req.body.name;
connection.query(`SELECT * FROM products WHERE name = '${name}'`, (err, rows) => {
  res.json(rows);
});

// Après — sécurisé
connection.query('SELECT * FROM products WHERE name = ?', [req.body.name], (err, rows) => {
  res.json(rows);
});
```

## Python (Django ORM)
```python
# Avant — vulnérable
name = request.GET.get("name")
products = Product.objects.raw(f"SELECT * FROM products WHERE name = '{name}'")

# Après — sécurisé
name = request.GET.get("name")
products = Product.objects.filter(name=name)
```

## Checklist de vérification post-patch
- [ ] Aucune comparaison SQL construite par concaténation de chaîne restante dans le fichier corrigé.
- [ ] Un test de non-régression confirme que la recherche/le filtre légitime retourne toujours les bons résultats.
- [ ] Un test confirme qu'une entrée contenant une expression logique SQL (ex: opérateur `OR`) n'altère pas le nombre de résultats retournés.
- [ ] Les identifiants attendus numériques sont validés/castés avant d'être utilisés dans la requête.
- [ ] Une limite de tentatives et une journalisation des motifs de requête anormaux sont en place sur les points d'entrée sensibles (connexion, recherche).
