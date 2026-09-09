# Remédiation — Coupon Abuse (abus de codes promotionnels)

## Principe
Enregistrer et vérifier l'usage des coupons côté serveur avec une contrainte d'unicité en base liée au compte, générer les codes avec une entropie suffisante, et revalider systématiquement le périmètre d'éligibilité (dates, montant, cumul) à chaque application.

## PHP
```php
// Avant — vulnérable : marquage "utilisé" côté client (cookie), code prévisible
$code = 'PROMO' . str_pad($userId, 6, '0', STR_PAD_LEFT); // énumérable
setcookie('coupon_used_' . $code, '1');

// Après — sécurisé
function generateCouponCode(): string {
    return strtoupper(bin2hex(random_bytes(6))); // haute entropie, non énumérable
}

function applyCoupon(PDO $pdo, int $userId, string $code): bool {
    $pdo->beginTransaction();
    $stmt = $pdo->prepare(
        "SELECT id, expires_at, min_amount FROM coupons WHERE code = ? FOR UPDATE"
    );
    $stmt->execute([$code]);
    $coupon = $stmt->fetch();
    if (!$coupon || strtotime($coupon['expires_at']) < time()) {
        $pdo->rollBack();
        return false;
    }
    // Contrainte UNIQUE(coupon_id, user_id) empêche toute réapplication
    $insert = $pdo->prepare("INSERT INTO coupon_redemptions (coupon_id, user_id) VALUES (?, ?)");
    try {
        $insert->execute([$coupon['id'], $userId]);
    } catch (PDOException $e) {
        $pdo->rollBack();
        return false; // déjà utilisé
    }
    $pdo->commit();
    return true;
}
```

## JS / Node
```js
// Avant — vulnérable
if (!localStorage.getItem(`coupon_${code}`)) {
  applyDiscount(code);
  localStorage.setItem(`coupon_${code}`, '1'); // contournable
}

// Après — sécurisé (contrainte unique en base : UNIQUE(coupon_id, user_id))
async function applyCoupon(db, userId, code) {
  const coupon = await db.query(
    'SELECT id, expires_at, min_amount FROM coupons WHERE code = $1', [code]
  );
  if (!coupon.rows[0] || new Date(coupon.rows[0].expires_at) < new Date()) {
    throw new Error('Coupon invalide ou expiré');
  }
  try {
    await db.query(
      'INSERT INTO coupon_redemptions (coupon_id, user_id) VALUES ($1, $2)',
      [coupon.rows[0].id, userId]
    );
  } catch (e) {
    if (e.code === '23505') throw new Error('Coupon déjà utilisé'); // violation contrainte unique
    throw e;
  }
}
```

## Python
```python
# Avant — vulnérable : code séquentiel, pas de contrainte serveur
def generate_code(user_id):
    return f"REF{user_id:06d}"

# Après — sécurisé
import secrets

def generate_coupon_code() -> str:
    return secrets.token_urlsafe(9).upper()

def apply_coupon(db, user_id: int, code: str) -> bool:
    coupon = db.fetch_one(
        "SELECT id, expires_at FROM coupons WHERE code = %s FOR UPDATE", [code]
    )
    if not coupon or coupon["expires_at"] < now():
        return False
    try:
        db.execute(
            "INSERT INTO coupon_redemptions (coupon_id, user_id) VALUES (%s, %s)",
            [coupon["id"], user_id],
        )  # contrainte UNIQUE(coupon_id, user_id) en base
    except UniqueViolation:
        return False
    return True
```

## Checklist de vérification post-patch
- [ ] Le statut "coupon utilisé" est enregistré côté serveur avec une contrainte d'unicité `(coupon_id, user_id)`.
- [ ] Les codes promotionnels sont générés avec une entropie suffisante (CSPRNG), non séquentiels ni dérivés d'un ID utilisateur.
- [ ] Le périmètre d'éligibilité (dates, montant minimum, cumul) est revalidé côté serveur à chaque application.
- [ ] Une limitation de tentatives (rate limiting) est en place sur l'endpoint d'application de coupon.
