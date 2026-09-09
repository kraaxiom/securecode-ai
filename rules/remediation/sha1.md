# Remédiation — Utilisation de SHA-1 (hachage faible)

## Principe
SHA-1 est cryptographiquement cassé pour la résistance aux collisions (attaque SHAttered). Ne jamais l'utiliser pour signer, pour vérifier une intégrité sensible, ou pour hacher un mot de passe. Utiliser SHA-256 (ou supérieur) pour l'intégrité générale, et Argon2id pour les mots de passe.

## PHP
```php
// Avant — vulnérable
$hash = sha1($password);

// Après — sécurisé (mot de passe)
$hash = password_hash($password, PASSWORD_ARGON2ID);

// Après — sécurisé (intégrité de fichier / signature)
$hash = hash('sha256', $fileContent);
```

## Node.js
```js
// Avant — vulnérable
const hash = crypto.createHash('sha1').update(password).digest('hex');

// Après — sécurisé (mot de passe)
import argon2 from 'argon2';
const hash = await argon2.hash(password, { type: argon2.argon2id });

// Après — sécurisé (intégrité)
const hash = crypto.createHash('sha256').update(fileBuffer).digest('hex');
```

## Python
```python
# Avant — vulnérable
import hashlib
hash = hashlib.sha1(password.encode()).hexdigest()

# Après — sécurisé (mot de passe)
from argon2 import PasswordHasher
hash = PasswordHasher().hash(password)

# Après — sécurisé (intégrité)
hash = hashlib.sha256(file_content).hexdigest()
```

## Checklist de vérification post-patch
- [ ] Plus aucun appel à `sha1()` pour un mot de passe, une signature ou un token de sécurité.
- [ ] Les mots de passe utilisent Argon2id (ou bcrypt), les usages d'intégrité générale utilisent SHA-256+.
- [ ] Les certificats/signatures encore basés sur SHA-1 côté infrastructure sont recensés pour migration.
- [ ] Un plan de migration des hachages existants (mots de passe) est en place si applicable.
