# Remédiation — Double Spending (double dépense)

## Principe
Encapsuler toute opération vérifier-puis-débiter dans une transaction atomique avec verrouillage (`SELECT ... FOR UPDATE`), exiger une clé d'idempotence sur les endpoints financiers et les webhooks, et contraindre le solde à ne jamais devenir négatif au niveau de la base.

## PHP
```php
// Avant — vulnérable : vérification et débit non atomiques
$balance = $pdo->query("SELECT balance FROM wallets WHERE id = $walletId")->fetchColumn();
if ($balance >= $amount) {
    $pdo->exec("UPDATE wallets SET balance = balance - $amount WHERE id = $walletId");
}

// Après — sécurisé
function withdraw(PDO $pdo, int $walletId, float $amount, string $idempotencyKey): bool {
    $pdo->beginTransaction();
    // Rejet des rejeux via clé d'idempotence
    $check = $pdo->prepare("SELECT 1 FROM transactions WHERE idempotency_key = ?");
    $check->execute([$idempotencyKey]);
    if ($check->fetch()) { $pdo->rollBack(); return false; }

    $stmt = $pdo->prepare("SELECT balance FROM wallets WHERE id = ? FOR UPDATE");
    $stmt->execute([$walletId]);
    $balance = (float) $stmt->fetchColumn();
    if ($balance < $amount) { $pdo->rollBack(); return false; }

    $pdo->prepare("UPDATE wallets SET balance = balance - ? WHERE id = ?")->execute([$amount, $walletId]);
    $pdo->prepare("INSERT INTO transactions (wallet_id, amount, idempotency_key) VALUES (?, ?, ?)")
        ->execute([$walletId, -$amount, $idempotencyKey]);
    $pdo->commit();
    return true;
}
// + contrainte SQL : ALTER TABLE wallets ADD CONSTRAINT chk_balance CHECK (balance >= 0);
```

## JS / Node
```js
// Avant — vulnérable
const wallet = await db.query('SELECT balance FROM wallets WHERE id = $1', [walletId]);
if (wallet.rows[0].balance >= amount) {
  await db.query('UPDATE wallets SET balance = balance - $1 WHERE id = $2', [amount, walletId]);
}

// Après — sécurisé
async function withdraw(db, walletId, amount, idempotencyKey) {
  return db.transaction(async (tx) => {
    const existing = await tx.query(
      'SELECT 1 FROM transactions WHERE idempotency_key = $1', [idempotencyKey]
    );
    if (existing.rows.length) return false; // déjà traité

    const { rows } = await tx.query(
      'SELECT balance FROM wallets WHERE id = $1 FOR UPDATE', [walletId]
    );
    if (rows[0].balance < amount) return false;

    await tx.query('UPDATE wallets SET balance = balance - $1 WHERE id = $2', [amount, walletId]);
    await tx.query(
      'INSERT INTO transactions (wallet_id, amount, idempotency_key) VALUES ($1, $2, $3)',
      [walletId, -amount, idempotencyKey]
    );
    return true;
  });
}
```

## Python
```python
# Avant — vulnérable
balance = db.fetch_one("SELECT balance FROM wallets WHERE id = %s", [wallet_id])["balance"]
if balance >= amount:
    db.execute("UPDATE wallets SET balance = balance - %s WHERE id = %s", [amount, wallet_id])

# Après — sécurisé
def withdraw(db, wallet_id: int, amount: float, idempotency_key: str) -> bool:
    with db.transaction():
        if db.fetch_one(
            "SELECT 1 FROM transactions WHERE idempotency_key = %s", [idempotency_key]
        ):
            return False
        row = db.fetch_one(
            "SELECT balance FROM wallets WHERE id = %s FOR UPDATE", [wallet_id]
        )
        if row["balance"] < amount:
            return False
        db.execute(
            "UPDATE wallets SET balance = balance - %s WHERE id = %s", [amount, wallet_id]
        )
        db.execute(
            "INSERT INTO transactions (wallet_id, amount, idempotency_key) VALUES (%s, %s, %s)",
            [wallet_id, -amount, idempotency_key],
        )
    return True
```

## Checklist de vérification post-patch
- [ ] La séquence vérifier-puis-débiter est encapsulée dans une transaction atomique avec `SELECT ... FOR UPDATE`.
- [ ] Chaque transaction/webhook porte une clé d'idempotence unique rejetant les traitements dupliqués.
- [ ] Une contrainte `CHECK (balance >= 0)` (ou équivalent) empêche un solde négatif au niveau base de données.
- [ ] Un test de charge avec requêtes concurrentes confirme l'absence de double débit.
