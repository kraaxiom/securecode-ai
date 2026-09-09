# Remédiation — Stacked Query SQL Injection

## Principe
Utiliser systématiquement des requêtes préparées avec paramètres liés, et désactiver l'exécution multi-instructions (stacked queries) au niveau du driver lorsque l'application n'en a pas besoin fonctionnellement.

## PHP (PDO)
```php
// Avant — vulnérable
$name = $_POST['name'];
$pdo->query("UPDATE users SET name = '$name' WHERE id = 1");

// Après — sécurisé
// Émulation des préparations désactivée pour empêcher le multi-statement côté client
$pdo->setAttribute(PDO::ATTR_EMULATE_PREPARES, false);
$stmt = $pdo->prepare("UPDATE users SET name = :name WHERE id = 1");
$stmt->execute(['name' => $name]);
```

## Node.js (mysql2)
```js
// Avant — vulnérable
const connection = mysql.createConnection({ multipleStatements: true, ...config });
connection.query(`UPDATE users SET name = '${req.body.name}' WHERE id = 1`);

// Après — sécurisé
const connection = mysql.createConnection({ multipleStatements: false, ...config });
connection.query('UPDATE users SET name = ? WHERE id = 1', [req.body.name]);
```

## Python (psycopg2)
```python
# Avant — vulnérable
cursor.execute(f"UPDATE users SET name = '{name}' WHERE id = 1")

# Après — sécurisé
cursor.execute("UPDATE users SET name = %s WHERE id = 1", [name])
```

## Checklist de vérification post-patch
- [ ] L'option multi-statement du driver (ex: `multipleStatements`, `ATTR_EMULATE_PREPARES`) est désactivée si non requise fonctionnellement.
- [ ] Toutes les requêtes du fichier corrigé utilisent des paramètres liés, sans exception.
- [ ] Un test confirme qu'un point-virgule dans une entrée utilisateur ne permet plus d'ajouter une instruction SQL distincte.
- [ ] Le compte de connexion à la base de données applique le principe du moindre privilège (pas de droits DDL/DML superflus).
