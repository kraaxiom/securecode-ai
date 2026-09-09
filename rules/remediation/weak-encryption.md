# Remédiation — Weak Encryption (chiffrement faible des données)

## Principe
Hacher les mots de passe avec argon2id, chiffrer les données sensibles au niveau colonne avec un algorithme moderne authentifié (AES-GCM), et gérer les clés via un service dédié (KMS/HSM) séparé du stockage des données.

## PHP
```php
// Avant — vulnérable : MD5 sans sel pour les mots de passe
$hash = md5($password);

// Après — sécurisé : argon2id
$hash = password_hash($password, PASSWORD_ARGON2ID, ['memory_cost' => 65536, 'time_cost' => 4, 'threads' => 2]);
password_verify($inputPassword, $hash);

// Chiffrement de colonne sensible — AES-256-GCM, clé issue d'un KMS (jamais en dur)
function encryptField(string $plaintext, string $key): string {
    $iv = random_bytes(12);
    $ciphertext = openssl_encrypt($plaintext, 'aes-256-gcm', $key, OPENSSL_RAW_DATA, $iv, $tag);
    return base64_encode($iv . $tag . $ciphertext);
}
```

## JS / Node
```js
// Avant — vulnérable
const crypto = require('crypto');
const hash = crypto.createHash('sha1').update(password).digest('hex');

// Après — sécurisé : argon2id
const argon2 = require('argon2');
const hash = await argon2.hash(password, { type: argon2.argon2id, memoryCost: 65536, timeCost: 4 });
await argon2.verify(hash, inputPassword);

// Chiffrement de colonne — AES-256-GCM, clé récupérée depuis un KMS
function encryptField(plaintext, key) {
  const iv = crypto.randomBytes(12);
  const cipher = crypto.createCipheriv('aes-256-gcm', key, iv);
  const ciphertext = Buffer.concat([cipher.update(plaintext, 'utf8'), cipher.final()]);
  return Buffer.concat([iv, cipher.getAuthTag(), ciphertext]).toString('base64');
}
```

## Python
```python
# Avant — vulnérable
import hashlib
hashed = hashlib.md5(password.encode()).hexdigest()

# Après — sécurisé : argon2id
from argon2 import PasswordHasher
ph = PasswordHasher()  # argon2id par défaut
hashed = ph.hash(password)
ph.verify(hashed, input_password)

# Chiffrement de colonne — AES-256-GCM, clé issue d'un KMS externe
from cryptography.hazmat.primitives.ciphers.aead import AESGCM
import os

def encrypt_field(plaintext: bytes, key: bytes) -> bytes:
    nonce = os.urandom(12)
    aesgcm = AESGCM(key)  # key provient du KMS, jamais codée en dur
    return nonce + aesgcm.encrypt(nonce, plaintext, None)
```

## Checklist de vérification post-patch
- [ ] Les mots de passe sont hachés avec argon2id (ou bcrypt à défaut), jamais MD5/SHA-1 sans sel.
- [ ] Les données sensibles en base sont chiffrées avec un algorithme moderne authentifié (AES-GCM), pas ECB ni CBC sans MAC.
- [ ] Aucune clé de chiffrement n'est codée en dur dans le code source ou stockée avec les données qu'elle protège.
- [ ] Les clés sont gérées via un service dédié (KMS/HSM) avec rotation régulière.
