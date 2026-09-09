# Remédiation — Identifiant de session faible

## Principe
Générer tout identifiant de session ou jeton exclusivement avec un générateur aléatoire cryptographiquement sûr, et s'appuyer sur le mécanisme de gestion de session natif du framework plutôt que de le réimplémenter, en garantissant au moins 128 bits d'entropie.

## PHP
```php
// Avant — vulnérable
$sessionId = md5(time() . rand());

// Après — sécurisé
$sessionId = bin2hex(random_bytes(32)); // CSPRNG, 256 bits d'entropie
// Ou mieux : s'appuyer sur le générateur natif de PHP
session_start(); // PHP génère déjà un ID cryptographiquement sûr
```

## Node.js (Express)
```js
// Avant — vulnérable
const sessionId = Buffer.from(`${Date.now()}-${Math.random()}`).toString('hex');

// Après — sécurisé
const crypto = require('crypto');
const sessionId = crypto.randomBytes(32).toString('hex'); // CSPRNG

// Ou déléguer à express-session qui utilise déjà un CSPRNG en interne
app.use(session({ secret: process.env.SESSION_SECRET, genid: () => crypto.randomBytes(32).toString('hex') }));
```

## Python (Flask/Django)
```python
# Avant — vulnérable
import random
session_id = str(random.random()) + str(int(time.time()))

# Après — sécurisé
import secrets
session_id = secrets.token_urlsafe(32)  # CSPRNG, ~256 bits d'entropie

# Flask/Django génèrent déjà un identifiant de session sûr nativement — préférer ces mécanismes
```

## Checklist de vérification post-patch
- [ ] Tout identifiant de session ou jeton est généré via un CSPRNG (`random_bytes`, `crypto.randomBytes`, `secrets`).
- [ ] Aucune concaténation de valeurs prévisibles (timestamp, ID incrémental) n'entre dans la génération.
- [ ] L'entropie de l'identifiant atteint au moins 128 bits.
- [ ] Le mécanisme de session natif du framework est utilisé plutôt qu'une implémentation maison.
