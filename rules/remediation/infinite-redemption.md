# Remédiation — Infinite Redemption (rachat/échange illimité)

## Principe
Marquer l'octroi d'un avantage de façon atomique et transactionnelle côté serveur avant toute confirmation, protéger l'endpoint de réclamation contre les requêtes concurrentes via verrouillage ou contrainte d'unicité, et vérifier l'éligibilité via des signaux résistants à la duplication triviale.

## PHP
```php
// Avant — vulnérable : pas de verrou, réclamation concurrente possible
$claimed = $pdo->query("SELECT reward_claimed FROM users WHERE id = $userId")->fetchColumn();
if (!$claimed) {
    $pdo->exec("UPDATE users SET reward_claimed = 1, credits = credits + 10 WHERE id = $userId");
}

// Après — sécurisé
function claimReferralReward(PDO $pdo, int $userId): bool {
    $pdo->beginTransaction();
    $stmt = $pdo->prepare("SELECT reward_claimed FROM users WHERE id = ? FOR UPDATE");
    $stmt->execute([$userId]);
    if ((bool) $stmt->fetchColumn()) {
        $pdo->rollBack();
        return false;
    }
    $pdo->prepare("UPDATE users SET reward_claimed = 1, credits = credits + 10 WHERE id = ?")
        ->execute([$userId]);
    // Registre d'événements immuable pour audit et détection de doublons
    $pdo->prepare("INSERT INTO reward_events (user_id, type, amount) VALUES (?, 'referral', 10)")
        ->execute([$userId]);
    $pdo->commit();
    return true;
}
```

## JS / Node
```js
// Avant — vulnérable
const user = await db.query('SELECT reward_claimed FROM users WHERE id = $1', [userId]);
if (!user.rows[0].reward_claimed) {
  await db.query('UPDATE users SET reward_claimed = true, credits = credits + 10 WHERE id = $1', [userId]);
}

// Après — sécurisé
async function claimReward(db, userId) {
  return db.transaction(async (tx) => {
    const { rows } = await tx.query(
      'SELECT reward_claimed FROM users WHERE id = $1 FOR UPDATE', [userId]
    );
    if (rows[0].reward_claimed) return false;
    await tx.query(
      'UPDATE users SET reward_claimed = true, credits = credits + 10 WHERE id = $1', [userId]
    );
    await tx.query(
      "INSERT INTO reward_events (user_id, type, amount) VALUES ($1, 'referral', 10)", [userId]
    );
    return true;
  });
}
```

## Python
```python
# Avant — vulnérable
claimed = db.fetch_one("SELECT reward_claimed FROM users WHERE id = %s", [user_id])["reward_claimed"]
if not claimed:
    db.execute("UPDATE users SET reward_claimed = true, credits = credits + 10 WHERE id = %s", [user_id])

# Après — sécurisé
def claim_reward(db, user_id: int) -> bool:
    with db.transaction():
        row = db.fetch_one(
            "SELECT reward_claimed FROM users WHERE id = %s FOR UPDATE", [user_id]
        )
        if row["reward_claimed"]:
            return False
        db.execute(
            "UPDATE users SET reward_claimed = true, credits = credits + 10 WHERE id = %s",
            [user_id],
        )
        db.execute(
            "INSERT INTO reward_events (user_id, type, amount) VALUES (%s, 'referral', 10)",
            [user_id],
        )
    return True
```

## Checklist de vérification post-patch
- [ ] L'octroi de l'avantage est marqué de façon atomique et transactionnelle (`FOR UPDATE` ou équivalent) avant confirmation.
- [ ] L'endpoint de réclamation est protégé contre les requêtes concurrentes (verrou ou contrainte d'unicité).
- [ ] L'éligibilité repose sur plusieurs signaux (pas seulement e-mail/IP, facilement recréables).
- [ ] Un registre d'événements immuable trace chaque octroi pour audit et détection de doublons.
