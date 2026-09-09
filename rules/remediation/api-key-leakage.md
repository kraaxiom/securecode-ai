# Remédiation — Fuite de clé API

## Principe
Ne jamais coder en dur un secret dans le code source ou un fichier versionné : charger les clés API depuis des variables d'environnement injectées à l'exécution ou depuis un gestionnaire de secrets dédié, et exclure systématiquement les fichiers sensibles du contrôle de version.

## PHP
```php
// Avant — vulnérable
<?php
$apiKey = "sk_live_EXAMPLE_NOT_A_REAL_KEY";
$client = new PaymentClient($apiKey);

// Après — sécurisé
<?php
$apiKey = getenv('PAYMENT_API_KEY');
if ($apiKey === false) {
    throw new RuntimeException('PAYMENT_API_KEY non configurée');
}
$client = new PaymentClient($apiKey);
```

## Node.js (Express)
```js
// Avant — vulnérable
const stripe = require('stripe')('sk_live_EXAMPLE_NOT_A_REAL_KEY');

// Après — sécurisé
require('dotenv').config();
if (!process.env.STRIPE_SECRET_KEY) {
  throw new Error('STRIPE_SECRET_KEY manquante dans l\'environnement');
}
const stripe = require('stripe')(process.env.STRIPE_SECRET_KEY);
```

## Python (Flask/Django/FastAPI)
```python
# Avant — vulnérable
API_KEY = "sk_live_EXAMPLE_NOT_A_REAL_KEY"
client = PaymentClient(api_key=API_KEY)

# Après — sécurisé
import os

api_key = os.environ.get("PAYMENT_API_KEY")
if not api_key:
    raise RuntimeError("PAYMENT_API_KEY non configurée dans l'environnement")
client = PaymentClient(api_key=api_key)
```

## Checklist de vérification post-patch
- [ ] Aucune chaîne littérale ressemblant à une clé API (préfixe connu, forte entropie) ne subsiste dans le code source.
- [ ] Les fichiers `.env`, `*.pem`, `*.key` sont bien listés dans `.gitignore` et absents de l'historique Git.
- [ ] Les clés précédemment exposées ont été révoquées et régénérées (rotation immédiate).
- [ ] Aucune clé serveur à privilège élevé n'est présente dans le bundle JavaScript livré au navigateur.
