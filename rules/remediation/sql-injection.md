# Remédiation — SQL Injection (générique)

## Principe
Utiliser systématiquement des requêtes préparées avec paramètres liés — jamais de concaténation ou d'interpolation de valeur dans une chaîne SQL. Pour les identifiants dynamiques (noms de colonnes/tables, clauses `ORDER BY`), utiliser une liste blanche stricte plutôt que la valeur brute. Valider et typer les entrées avant la requête, et appliquer le principe du moindre privilège sur le compte de connexion.

## PHP (PDO)
```php
// Avant — vulnérable
$name = $_GET['name'];
$result = $pdo->query("SELECT * FROM users WHERE name = '$name'");

// Après — sécurisé
$stmt = $pdo->prepare("SELECT * FROM users WHERE name = :name");
$stmt->execute(['name' => $_GET['name']]);
$result = $stmt->fetchAll();

// Cas d'un tri dynamique — utiliser une liste blanche, jamais la valeur brute
$allowedSort = ['name', 'created_at'];
$sort = in_array($_GET['sort'], $allowedSort, true) ? $_GET['sort'] : 'name';
$stmt = $pdo->prepare("SELECT * FROM users ORDER BY $sort");
$stmt->execute();
```

## JavaScript / Node.js (mysql2 / pg)
```js
// Avant — vulnérable
const name = req.query.name;
connection.query(`SELECT * FROM users WHERE name = '${name}'`);

// Après — sécurisé
connection.query('SELECT * FROM users WHERE name = ?', [req.query.name]);

// Cas d'un tri dynamique — liste blanche
const allowedSort = ['name', 'created_at'];
const sort = allowedSort.includes(req.query.sort) ? req.query.sort : 'name';
connection.query(`SELECT * FROM users ORDER BY ${sort}`); // sûr car sort vient de la liste blanche
```

## Python (psycopg2 / Django ORM)
```python
# Avant — vulnérable
name = request.GET['name']
cursor.execute(f"SELECT * FROM users WHERE name = '{name}'")

# Après — sécurisé
cursor.execute("SELECT * FROM users WHERE name = %s", [request.GET['name']])

# Cas d'un tri dynamique — liste blanche
ALLOWED_SORT = {'name', 'created_at'}
sort = request.GET.get('sort', 'name')
sort = sort if sort in ALLOWED_SORT else 'name'
cursor.execute(f"SELECT * FROM users ORDER BY {sort}")  # sûr car sort vient de la liste blanche
```

## Checklist de vérification post-patch
- [ ] Aucune concaténation ou f-string/template literal contenant une variable n'est utilisée pour construire une requête SQL.
- [ ] Toute valeur utilisateur passe par un paramètre lié (placeholder), jamais insérée directement dans le texte de la requête.
- [ ] Les identifiants dynamiques (colonnes, tables, direction de tri) passent par une liste blanche explicite avant insertion dans la requête.
- [ ] Le compte de connexion à la base de données n'a que les privilèges strictement nécessaires à l'application.
- [ ] Un test confirme qu'une entrée contenant un caractère `'` ou un mot-clé SQL ne modifie pas le comportement de la requête.
