# Remédiation — UNION-based SQL Injection

## Principe
Remplacer toute concaténation de valeur dans une requête SQL par une requête préparée avec paramètre lié. Valider/typer l'entrée en amont quand c'est possible (ex: caster un ID en entier).

## PHP (PDO)
```php
// Avant — vulnérable
$id = $_GET['id'];
$result = $pdo->query("SELECT * FROM users WHERE id = $id");

// Après — sécurisé
$id = (int) $_GET['id'];
$stmt = $pdo->prepare("SELECT * FROM users WHERE id = :id");
$stmt->execute(['id' => $id]);
$result = $stmt->fetchAll();
```

## PHP (Laravel / Eloquent)
```php
// Avant — vulnérable
$users = DB::select("SELECT * FROM users WHERE id = " . $request->id);

// Après — sécurisé
$users = DB::select("SELECT * FROM users WHERE id = ?", [$request->integer('id')]);
```

## Node.js (mysql2)
```js
// Avant — vulnérable
connection.query(`SELECT * FROM users WHERE id = ${req.query.id}`);

// Après — sécurisé
connection.query('SELECT * FROM users WHERE id = ?', [Number(req.query.id)]);
```

## Python (psycopg2 / Django ORM)
```python
# Avant — vulnérable
cursor.execute(f"SELECT * FROM users WHERE id = {user_id}")

# Après — sécurisé
cursor.execute("SELECT * FROM users WHERE id = %s", [int(user_id)])
```

## Java (JDBC)
```java
// Avant — vulnérable
Statement stmt = conn.createStatement();
stmt.executeQuery("SELECT * FROM users WHERE id = " + id);

// Après — sécurisé
PreparedStatement stmt = conn.prepareStatement("SELECT * FROM users WHERE id = ?");
stmt.setInt(1, Integer.parseInt(id));
stmt.executeQuery();
```

## Checklist de vérification post-patch
- [ ] Aucune concaténation de variable restante dans une requête SQL du fichier corrigé.
- [ ] Un test de non-régression confirme que la requête légitime fonctionne toujours (ex: recherche par ID valide).
- [ ] Un test confirme qu'une entrée contenant un caractère `'` ou un mot-clé SQL ne modifie pas le comportement de la requête.
- [ ] Le compte de connexion SQL utilisé n'a que les privilèges nécessaires (pas de droits DDL si non requis).
