# Remédiation — Secret JWT faible ou codé en dur

## Principe
Un secret JWT court, prévisible ou codé en dur permet à un attaquant de forger des tokens valides (bruteforce hors ligne ou lecture du code source). Utiliser un secret aléatoire d'au moins 256 bits stocké hors du code, ou passer à une signature asymétrique (RS256/ES256) pour découpler émission et vérification.

## PHP (firebase/php-jwt)
```php
// Avant — vulnérable
$secret = "secret123";
$jwt = JWT::encode($payload, $secret, 'HS256');

// Après — sécurisé
$secret = getenv('JWT_SECRET'); // généré via random_bytes(32) puis stocké en secret manager
$jwt = JWT::encode($payload, $secret, 'HS256');
// Alternative recommandée : signature asymétrique
$jwt = JWT::encode($payload, $privateKeyRS256, 'RS256');
```

## Node.js (jsonwebtoken)
```js
// Avant — vulnérable
const token = jwt.sign(payload, 'mysecret');

// Après — sécurisé
const token = jwt.sign(payload, process.env.JWT_SECRET, { algorithm: 'HS256', expiresIn: '15m' });
// Alternative recommandée : RS256 avec clé privée dédiée
const token = jwt.sign(payload, privateKey, { algorithm: 'RS256', expiresIn: '15m' });
```

## Python (PyJWT)
```python
# Avant — vulnérable
token = jwt.encode(payload, "secret", algorithm="HS256")

# Après — sécurisé
import os
secret = os.environ["JWT_SECRET"]  # >= 32 octets aléatoires
token = jwt.encode(payload, secret, algorithm="HS256")
```

## Checklist de vérification post-patch
- [ ] Le secret JWT n'est plus codé en dur dans le code source.
- [ ] Le secret fait au moins 256 bits d'entropie et est stocké dans un gestionnaire de secrets/variable d'environnement.
- [ ] L'algorithme de signature est explicitement vérifié côté serveur (pas d'acceptation de `alg: none`).
- [ ] Les tokens émis avec l'ancien secret sont invalidés (rotation de secret + courte durée d'expiration).
- [ ] Une expiration courte (`exp`) est appliquée sur les tokens émis.
