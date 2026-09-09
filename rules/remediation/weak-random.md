# Remédiation — Générateur aléatoire non cryptographique

## Principe
`rand()`, `Math.random()` ou `random` (module Python standard) sont des PRNG statistiques prévisibles, inadaptés à la génération de tokens, clés de session, ou identifiants de réinitialisation de mot de passe. Utiliser systématiquement un générateur cryptographiquement sûr (CSPRNG).

## PHP
```php
// Avant — vulnérable
$token = md5(rand());

// Après — sécurisé
$token = bin2hex(random_bytes(32));
```

## Node.js
```js
// Avant — vulnérable
const token = Math.random().toString(36).substring(2);

// Après — sécurisé
const token = crypto.randomBytes(32).toString('hex');
```

## Python
```python
# Avant — vulnérable
import random
token = str(random.random())

# Après — sécurisé
import secrets
token = secrets.token_hex(32)
```

## Checklist de vérification post-patch
- [ ] Plus aucun usage de `rand()`, `Math.random()` ou `random.*` pour un token, une clé, un identifiant de session ou un lien de réinitialisation.
- [ ] Le générateur utilisé est un CSPRNG (`random_bytes` en PHP, `crypto.randomBytes` en Node, `secrets` en Python).
- [ ] La longueur du token généré offre une entropie suffisante (au moins 128 bits, soit 16 octets).
- [ ] Les tokens déjà émis avec l'ancien générateur sont invalidés si le risque le justifie.
