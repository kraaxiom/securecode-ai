# Remédiation — Utilisation de MD5 (hachage faible)

## Principe
Ne jamais utiliser MD5 pour stocker un mot de passe (collisions et cassage par force brute trop rapides). Utiliser un algorithme de hachage de mot de passe dédié avec sel et coût adaptatif (Argon2id de préférence, sinon bcrypt). Pour l'intégrité de données non sensibles (checksums), préférer SHA-256.

## PHP
```php
// Avant — vulnérable
$hash = md5($password);

// Après — sécurisé
$hash = password_hash($password, PASSWORD_ARGON2ID);
// Vérification :
if (password_verify($inputPassword, $hash)) { /* ok */ }
```

## Node.js
```js
// Avant — vulnérable
const hash = crypto.createHash('md5').update(password).digest('hex');

// Après — sécurisé
import argon2 from 'argon2';
const hash = await argon2.hash(password, { type: argon2.argon2id });
const valid = await argon2.verify(hash, password);
```

## Python
```python
# Avant — vulnérable
import hashlib
hash = hashlib.md5(password.encode()).hexdigest()

# Après — sécurisé
from argon2 import PasswordHasher
ph = PasswordHasher()
hash = ph.hash(password)
ph.verify(hash, password)  # lève une exception si invalide
```

## Checklist de vérification post-patch
- [ ] Plus aucun appel à `md5()` pour le hachage de mots de passe ou de secrets dans le code corrigé.
- [ ] L'algorithme utilisé est Argon2id (ou bcrypt à défaut), avec sel automatique et coût configuré.
- [ ] Un plan de migration des hachages existants est en place (re-hachage progressif à la prochaine connexion réussie).
- [ ] MD5 reste éventuellement utilisé uniquement pour des cas non-sensibles (checksum non cryptographique), jamais pour mots de passe ou signatures.
