# Remédiation — JWT signé avec un secret faible

## Principe
Générer le secret JWT avec une entropie cryptographique suffisante (256 bits minimum pour HS256), le stocker dans un gestionnaire de secrets, et envisager un algorithme asymétrique pour séparer signature et vérification.

## PHP
```php
// Avant — vulnérable : secret court, en dur dans le code
$secret = 'monsecret123';
$token = JWT::encode($payload, $secret, 'HS256');

// Après — sécurisé : secret aléatoire fort, chargé depuis un gestionnaire de secrets
$secret = env('JWT_SECRET'); // généré une fois via random_bytes(32) puis stocké dans le vault
if (strlen($secret) < 32) {
    throw new RuntimeException('JWT_SECRET insuffisant (256 bits minimum requis).');
}
$token = JWT::encode($payload, $secret, 'HS256');
```

## JS / Node.js
```js
// Avant — vulnérable : secret par défaut de la documentation, jamais changé
const JWT_SECRET = 'your-256-bit-secret';
const token = jwt.sign(payload, JWT_SECRET, { algorithm: 'HS256' });

// Après — sécurisé : secret chargé depuis un gestionnaire de secrets, entropie vérifiée
const JWT_SECRET = process.env.JWT_SECRET;
if (!JWT_SECRET || Buffer.byteLength(JWT_SECRET, 'utf8') < 32) {
  throw new Error('JWT_SECRET manquant ou insuffisant (256 bits minimum).');
}
const token = jwt.sign(payload, JWT_SECRET, { algorithm: 'HS256' });
```

## Python
```python
# Avant — vulnérable : secret faible codé en dur
JWT_SECRET = "secret"
token = jwt.encode(payload, JWT_SECRET, algorithm="HS256")

# Après — sécurisé : secret fort depuis variable d'environnement, longueur vérifiée
import os

JWT_SECRET = os.environ["JWT_SECRET"]  # généré via secrets.token_bytes(32), stocké dans un vault
if len(JWT_SECRET.encode()) < 32:
    raise RuntimeError("JWT_SECRET insuffisant (256 bits minimum requis).")
token = jwt.encode(payload, JWT_SECRET, algorithm="HS256")
```

## Checklist de vérification post-patch
- [ ] Le secret JWT n'apparaît en dur nulle part dans le code source ni dans un fichier versionné.
- [ ] Le secret provient d'un gestionnaire de secrets dédié et a une entropie d'au moins 256 bits pour HS256.
- [ ] Une procédure de rotation du secret existe et est testée.
- [ ] Un algorithme asymétrique (RS256/ES256) est envisagé pour les cas où la clé de vérification doit être partagée avec des tiers.
- [ ] Le code vérifie explicitement une longueur minimale du secret au démarrage de l'application.
