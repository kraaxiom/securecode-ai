# Remédiation — Race Condition (condition de concurrence métier)

## Principe
Utiliser des transactions avec verrouillage approprié (`SELECT ... FOR UPDATE`, verrou optimiste) pour toute opération lire-modifier-écrire critique, exiger une clé d'idempotence sur les endpoints sensibles, et déplacer les contraintes critiques au niveau de la base de données.

## PHP
```php
// Avant — vulnérable : lecture puis écriture non atomique sur le stock
$stock = $pdo->query("SELECT stock FROM products WHERE id = $productId")->fetchColumn();
if ($stock > 0) {
    $pdo->exec("UPDATE products SET stock = stock - 1 WHERE id = $productId");
}

// Après — sécurisé
function reserveStock(PDO $pdo, int $productId): bool {
    $pdo->beginTransaction();
    $stmt = $pdo->prepare("SELECT stock FROM products WHERE id = ? FOR UPDATE");
    $stmt->execute([$productId]);
    if ((int) $stmt->fetchColumn() <= 0) { $pdo->rollBack(); return false; }
    $pdo->prepare("UPDATE products SET stock = stock - 1 WHERE id = ? AND stock > 0")->execute([$productId]);
    $pdo->commit();
    return true;
}
// + contrainte SQL : ALTER TABLE products ADD CONSTRAINT chk_stock CHECK (stock >= 0);
```

## JS / Node
```js
// Avant — vulnérable
const p = await db.query('SELECT stock FROM products WHERE id = $1', [productId]);
if (p.rows[0].stock > 0) {
  await db.query('UPDATE products SET stock = stock - 1 WHERE id = $1', [productId]);
}

// Après — sécurisé
async function reserveStock(db, productId) {
  return db.transaction(async (tx) => {
    const { rows } = await tx.query(
      'SELECT stock FROM products WHERE id = $1 FOR UPDATE', [productId]
    );
    if (rows[0].stock <= 0) return false;
    await tx.query(
      'UPDATE products SET stock = stock - 1 WHERE id = $1 AND stock > 0', [productId]
    );
    return true;
  });
}
```

## Python
```python
# Avant — vulnérable
stock = db.fetch_one("SELECT stock FROM products WHERE id = %s", [product_id])["stock"]
if stock > 0:
    db.execute("UPDATE products SET stock = stock - 1 WHERE id = %s", [product_id])

# Après — sécurisé
def reserve_stock(db, product_id: int) -> bool:
    with db.transaction():
        row = db.fetch_one(
            "SELECT stock FROM products WHERE id = %s FOR UPDATE", [product_id]
        )
        if row["stock"] <= 0:
            return False
        db.execute(
            "UPDATE products SET stock = stock - 1 WHERE id = %s AND stock > 0", [product_id]
        )
    return True
```

## Checklist de vérification post-patch
- [ ] Toute séquence lire-modifier-écrire critique (solde, stock, quota) est protégée par une transaction avec verrou (`FOR UPDATE`) ou un verrou optimiste versionné.
- [ ] Les endpoints financiers/de consommation de quota exigent une clé d'idempotence.
- [ ] Une contrainte `CHECK` en base empêche les valeurs négatives ou incohérentes.
- [ ] Un test explicite sous requêtes concurrentes confirme l'absence de dépassement de quota/solde.
