# Remédiation — Account Takeover (prise de contrôle de compte)

## Principe
Générer les jetons de récupération/vérification avec un CSPRNG, à usage unique et à expiration courte ; exiger une confirmation sur l'ancien ET le nouveau contact lors d'un changement d'e-mail/téléphone ; ne fusionner un compte OAuth que si le fournisseur tiers certifie l'e-mail vérifié ; notifier l'utilisateur par un canal indépendant à chaque changement sensible.

## PHP
```php
// Avant — vulnérable : jeton prévisible, pas d'expiration, pas de vérification du nouvel e-mail
function resetToken($userId) {
    return md5($userId . time());
}
function changeEmail($userId, $newEmail) {
    $pdo->prepare("UPDATE users SET email = ? WHERE id = ?")->execute([$newEmail, $userId]);
}

// Après — sécurisé
function resetToken(): string {
    return bin2hex(random_bytes(32)); // CSPRNG, non dérivé de données prévisibles
}
function storeResetToken(PDO $pdo, int $userId, string $token): void {
    $hash = hash('sha256', $token);
    $stmt = $pdo->prepare(
        "INSERT INTO password_resets (user_id, token_hash, expires_at, used) VALUES (?, ?, ?, 0)"
    );
    $stmt->execute([$userId, $hash, date('Y-m-d H:i:s', time() + 900)]); // 15 min
}
function requestEmailChange(PDO $pdo, int $userId, string $newEmail): void {
    $token = bin2hex(random_bytes(32));
    $pdo->prepare(
        "INSERT INTO email_change_requests (user_id, new_email, token_hash, expires_at) VALUES (?, ?, ?, ?)"
    )->execute([$userId, $newEmail, hash('sha256', $token), date('Y-m-d H:i:s', time() + 3600)]);
    // Envoi d'une confirmation sur l'ANCIEN email (alerte) ET sur le NOUVEL email (lien de confirmation)
    notify_old_email($userId, 'Un changement d\'e-mail a été demandé.');
    send_confirmation_link($newEmail, $token);
}
```

## JS / Node
```js
// Avant — vulnérable
function makeToken(userId) {
  return Buffer.from(`${userId}-${Date.now()}`).toString('base64');
}

// Après — sécurisé
const crypto = require('crypto');

function makeResetToken() {
  return crypto.randomBytes(32).toString('hex'); // CSPRNG
}

async function storeResetToken(db, userId, token) {
  const tokenHash = crypto.createHash('sha256').update(token).digest('hex');
  const expiresAt = new Date(Date.now() + 15 * 60 * 1000);
  await db.query(
    'INSERT INTO password_resets (user_id, token_hash, expires_at, used) VALUES ($1, $2, $3, false)',
    [userId, tokenHash, expiresAt]
  );
}

// OAuth : ne fusionner que si le provider certifie l'email vérifié
function canMergeOAuthAccount(profile) {
  return profile.emailVerified === true;
}
```

## Python
```python
# Avant — vulnérable
import hashlib, time
def make_token(user_id):
    return hashlib.md5(f"{user_id}{int(time.time())}".encode()).hexdigest()

# Après — sécurisé
import secrets, hashlib
from datetime import datetime, timedelta

def make_reset_token() -> str:
    return secrets.token_hex(32)  # CSPRNG

def store_reset_token(db, user_id: int, token: str) -> None:
    token_hash = hashlib.sha256(token.encode()).hexdigest()
    expires_at = datetime.utcnow() + timedelta(minutes=15)
    db.execute(
        "INSERT INTO password_resets (user_id, token_hash, expires_at, used) VALUES (%s, %s, %s, false)",
        [user_id, token_hash, expires_at],
    )

def can_merge_oauth_account(profile: dict) -> bool:
    return profile.get("email_verified") is True
```

## Checklist de vérification post-patch
- [ ] Les jetons de réinitialisation/vérification sont générés via un CSPRNG (`random_bytes`, `crypto.randomBytes`, `secrets.token_hex`), jamais dérivés d'un timestamp ou d'un ID.
- [ ] Chaque jeton a une expiration courte et un flag "utilisé" empêchant sa réutilisation.
- [ ] Le changement d'e-mail/téléphone exige une confirmation sur l'ancien ET le nouveau contact avant activation.
- [ ] La fusion de compte OAuth n'a lieu que si l'e-mail est certifié vérifié par le fournisseur tiers.
- [ ] Une notification indépendante (autre canal que celui modifié) est envoyée à chaque changement sensible.
