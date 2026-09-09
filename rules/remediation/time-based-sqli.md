# Remédiation — Time-Based Blind SQL Injection

## Principe
La requête préparée avec paramètre lié élimine toute la classe de vulnérabilité, pas seulement la variante "time-based". En complément, configurer des timeouts stricts d'exécution côté base de données limite l'impact d'une éventuelle injection résiduelle.

## PHP (PDO)
```php
// Avant — vulnérable
$id = $_GET['id'];
$pdo->query("SELECT * FROM orders WHERE id = $id");

// Après — sécurisé
$stmt = $pdo->prepare("SELECT * FROM orders WHERE id = :id");
$stmt->setAttribute(PDO::ATTR_TIMEOUT, 5);
$stmt->execute(['id' => (int) $id]);
```

## Node.js (pg)
```js
// Avant — vulnérable
pool.query(`SELECT * FROM orders WHERE id = ${req.query.id}`);

// Après — sécurisé
pool.query('SELECT * FROM orders WHERE id = $1', [Number(req.query.id)]);
// Timeout appliqué au niveau du pool de connexion
const pool = new Pool({ ...config, statement_timeout: 5000 });
```

## Python (Django ORM)
```python
# Avant — vulnérable
Order.objects.raw(f"SELECT * FROM orders WHERE id = {order_id}")

# Après — sécurisé
Order.objects.raw("SELECT * FROM orders WHERE id = %s", [int(order_id)])
# Timeout au niveau de la connexion (settings.py)
# DATABASES['default']['OPTIONS'] = {'options': '-c statement_timeout=5000'}
```

## Checklist de vérification post-patch
- [ ] Toutes les requêtes du chemin corrigé utilisent des paramètres liés, sans concaténation résiduelle.
- [ ] Un timeout d'exécution de requête est configuré côté base de données pour limiter l'impact des injections résiduelles.
- [ ] Un test confirme qu'une entrée conçue pour provoquer un délai (ex: fonction de pause) n'a plus d'effet sur le temps de réponse.
- [ ] Une surveillance des temps de réponse anormaux est en place comme détection complémentaire.
