# Remédiation — Secrets exposés dans le code ou la configuration cloud

## Principe
Supprimer tout secret codé en dur, le remplacer par une référence à un gestionnaire de secrets dédié, révoquer/régénérer immédiatement le secret ayant fuité (même après suppression du code), et mettre en place un scan automatique en pre-commit et en CI.

## PHP
```php
// Avant — vulnérable : clé API codée en dur
$stripeClient = new \Stripe\StripeClient('sk_live_EXAMPLE_NOT_A_REAL_KEY');

// Après — sécurisé : secret injecté via un gestionnaire de secrets
$stripeClient = new \Stripe\StripeClient(
    getenv('STRIPE_SECRET_KEY') // valeur injectée depuis AWS Secrets Manager / Vault au démarrage
);
```

## Node.js
```js
// Avant — vulnérable
const dbConnection = mysql.createConnection({
  host: 'db.example.com',
  user: 'admin',
  password: 'Sup3rS3cret!2024', // en dur
});

// Après — sécurisé
const dbConnection = mysql.createConnection({
  host: process.env.DB_HOST,
  user: process.env.DB_USER,
  password: await secretsManager.getSecretValue('prod/db/password'),
});
```

## Fichier de configuration (.env)
```bash
# Avant — vulnérable : .env committé et accessible via le webroot
AWS_SECRET_ACCESS_KEY=AKIAIOSFODNN7EXAMPLEKEY

# Après — sécurisé : .env exclu du dépôt et hors webroot, secret dans un coffre
# .gitignore
.env
.env.*

# La valeur réelle est récupérée à l'exécution depuis le gestionnaire de secrets,
# jamais stockée en clair dans un fichier versionné.
```

## Étapes de remédiation obligatoires en cas de fuite confirmée
1. Révoquer/régénérer immédiatement le secret exposé auprès du fournisseur (ne pas se contenter de le supprimer du code).
2. Purger l'historique Git si le secret a été committé (`git filter-repo` ou BFG Repo-Cleaner), puis forcer la rotation même après purge.
3. Migrer le secret vers un gestionnaire dédié (Vault, AWS Secrets Manager, Azure Key Vault, GCP Secret Manager).
4. Auditer les logs d'accès du service concerné pour détecter un usage abusif pendant la période d'exposition.

## Checklist de vérification post-patch
- [ ] Aucun secret en clair ne subsiste dans le code source ou les fichiers de configuration versionnés.
- [ ] Le secret exposé a été révoqué et régénéré, pas seulement supprimé du code.
- [ ] Les fichiers sensibles (`.env`, `config.php` avec identifiants) sont ajoutés au `.gitignore`.
- [ ] Si le secret a été committé, l'historique Git a été purgé (filter-repo/BFG).
- [ ] L'application récupère désormais les secrets via un gestionnaire dédié, pas via un fichier en clair.
- [ ] Un scan automatique de secrets (gitleaks/trufflehog) est actif en pre-commit et en CI.
