# Remédiation — Dangerous Stored Procedures (procédures stockées dangereuses)

## Principe
Utiliser la liaison de paramètres même à l'intérieur des procédures stockées lors de la construction de SQL dynamique, désactiver les fonctionnalités d'exécution système natives du SGBD, et limiter les privilèges d'exécution au strict nécessaire.

## PHP (appel de procédure)
```php
// Avant — vulnérable : la procédure elle-même concatène ses paramètres en SQL dynamique
// CREATE PROCEDURE search_users(IN p_name VARCHAR(100))
// BEGIN
//   SET @sql = CONCAT('SELECT * FROM users WHERE name = ''', p_name, '''');
//   PREPARE stmt FROM @sql; EXECUTE stmt;
// END

// Après — sécurisé : la procédure lie le paramètre au lieu de le concaténer
// CREATE PROCEDURE search_users(IN p_name VARCHAR(100))
// BEGIN
//   SET @sql = 'SELECT * FROM users WHERE name = ?';
//   PREPARE stmt FROM @sql;
//   SET @p_name = p_name;
//   EXECUTE stmt USING @p_name;
//   DEALLOCATE PREPARE stmt;
// END

$stmt = $pdo->prepare("CALL search_users(?)");
$stmt->execute([$request->input('name')]);
```

## JS / Node (appel de procédure)
```js
// Avant — vulnérable : le code applicatif construit l'appel par concaténation
await db.query(`CALL search_users('${req.query.name}')`);

// Après — sécurisé : paramètre lié à l'appel de procédure
await db.query('CALL search_users(?)', [req.query.name]);
// Et côté SGBD : la procédure elle-même doit lier ses paramètres (voir ci-dessus),
// ne jamais exécuter de fonctionnalités shell natives (ex: xp_cmdshell sur SQL Server désactivé).
```

## Python (appel de procédure)
```python
# Avant — vulnérable
cursor.execute(f"CALL search_users('{name}')")

# Après — sécurisé
cursor.callproc("search_users", [name])  # paramètre lié via l'API du driver

# Restriction complémentaire au niveau du SGBD (exemple SQL Server) :
# EXEC sp_configure 'xp_cmdshell', 0; RECONFIGURE;
# GRANT EXECUTE ON search_users TO app_role;  -- rôle applicatif limité, pas dbowner
```

## Checklist de vérification post-patch
- [ ] Toute construction de SQL dynamique à l'intérieur d'une procédure stockée utilise la liaison de paramètres (`PREPARE`/`USING`), jamais la concaténation.
- [ ] Les fonctionnalités d'exécution système natives du SGBD (ex: `xp_cmdshell`, `COPY PROGRAM`) sont désactivées si non indispensables.
- [ ] Les procédures s'exécutent avec le rôle applicatif minimal requis, pas avec des privilèges d'administration (`SECURITY DEFINER` audité si utilisé).
- [ ] Les procédures stockées sont incluses dans le périmètre des revues de sécurité et des scans de code.
