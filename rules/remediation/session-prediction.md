# Remédiation — Session Prediction

## Principe
Générer tout identifiant de session exclusivement via un générateur cryptographiquement sûr, avec une entropie d'au moins 128 bits, en utilisant systématiquement les mécanismes fournis par le framework plutôt qu'une implémentation maison.

## PHP
```php
// Avant — vulnérable : ID de session prévisible dérivé de valeurs connues
$sessionId = md5($userId . time());

// Après — sécurisé : générateur cryptographique dédié, entropie suffisante
$sessionId = bin2hex(random_bytes(32)); // 256 bits d'entropie
// Préférer, en Laravel/PHP natif, le gestionnaire de session du framework qui applique déjà cela :
// session_regenerate_id(true); // à chaque changement de niveau de privilège
```

## JS / Node.js
```js
// Avant — vulnérable : ID de session basé sur un compteur/timestamp
let counter = 0;
function generateSessionId() {
  return `sess_${Date.now()}_${counter++}`;
}

// Après — sécurisé : générateur cryptographique dédié via express-session + store sécurisé
const crypto = require('crypto');

app.use(session({
  genid: () => crypto.randomBytes(32).toString('hex'), // 256 bits d'entropie
  secret: process.env.SESSION_SECRET,
  resave: false,
  saveUninitialized: false,
  cookie: { httpOnly: true, secure: true, sameSite: 'strict' },
}));

// Régénération à chaque élévation de privilège :
req.session.regenerate(() => { req.session.userId = user.id; });
```

## Python (Django)
```python
# Avant — vulnérable : identifiant de session maison basé sur un hash prévisible
import hashlib
session_id = hashlib.md5(f"{user.id}{time.time()}".encode()).hexdigest()

# Après — sécurisé : utiliser le framework de session Django (CSPRNG intégré)
from django.contrib.auth import login

def login_view(request, user):
    login(request, user)  # Django génère un session_key via un CSPRNG (get_random_string)
    request.session.cycle_key()  # renouvellement explicite après authentification
```

## Checklist de vérification post-patch
- [ ] Aucun identifiant de session n'est dérivé d'une valeur prévisible (timestamp, compteur, hash d'ID utilisateur).
- [ ] La génération repose exclusivement sur un CSPRNG (générateur cryptographique) avec au moins 128 bits d'entropie.
- [ ] L'identifiant de session est régénéré à chaque changement de niveau de privilège (connexion, élévation).
- [ ] Le mécanisme de session utilisé est celui fourni par le framework, pas une implémentation maison.
- [ ] Les cookies de session sont marqués `HttpOnly`, `Secure` et `SameSite`.
