# Remédiation — Clés/secrets codés en dur

## Principe
Retirer toute clé API, mot de passe, clé privée ou secret cryptographique du code source et le déplacer vers un gestionnaire de secrets ou des variables d'environnement non versionnées. Révoquer et régénérer immédiatement tout secret ayant été exposé dans l'historique git.

## PHP
```php
// Avant — vulnérable
$apiKey = "sk_live_EXAMPLE_NOT_A_REAL_KEY";
$stripe = new \Stripe\StripeClient($apiKey);

// Après — sécurisé
$apiKey = getenv('STRIPE_SECRET_KEY');
if (!$apiKey) {
    throw new \RuntimeException('STRIPE_SECRET_KEY manquant dans l\'environnement');
}
$stripe = new \Stripe\StripeClient($apiKey);
```

## Node.js
```js
// Avant — vulnérable
const apiKey = "AKIAIOSFODNN7EXAMPLE";
const s3 = new AWS.S3({ accessKeyId: apiKey, secretAccessKey: "wJalrXUtnFEMI/K7MDENG..." });

// Après — sécurisé
import 'dotenv/config';
const s3 = new AWS.S3({
  accessKeyId: process.env.AWS_ACCESS_KEY_ID,
  secretAccessKey: process.env.AWS_SECRET_ACCESS_KEY,
});
```

## Python
```python
# Avant — vulnérable
API_KEY = "AIzaSyD-EXAMPLE-KEY-1234567890"

# Après — sécurisé
import os
API_KEY = os.environ["API_KEY"]  # lève une exception explicite si absent
```

## Checklist de vérification post-patch
- [ ] Aucun littéral ressemblant à une clé/token/mot de passe ne reste dans le code corrigé.
- [ ] Le secret est chargé via variable d'environnement ou gestionnaire de secrets (Vault, AWS Secrets Manager, etc.).
- [ ] Le fichier `.env` (ou équivalent) est bien listé dans `.gitignore`.
- [ ] Le secret exposé a été révoqué et régénéré côté fournisseur — pas seulement retiré du code.
- [ ] L'historique git a été purgé du secret si celui-ci a été committé (BFG Repo-Cleaner ou `git filter-repo`).
