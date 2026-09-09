## Vulnérabilité

Clés/secrets codés en dur dans le code source (CWE-798 — Use of Hard-coded Credentials), illustré par une clé API de paiement et une clé de chiffrement intégrées littéralement dans un service Laravel.

## Impact

Toute personne ayant accès au dépôt Git (y compris à son historique), à une archive de code, ou au binaire compilé, peut lire le secret en clair. Cela permet l'usurpation d'appels API tiers (facturation frauduleuse, exfiltration de données via le compte compromis), ou le déchiffrement de toutes les données protégées par la clé de chiffrement embarquée. Le risque persiste même après suppression du secret du code si l'historique Git n'est pas purgé.

## Cause racine

Constantes ou propriétés de classe initialisées avec une valeur littérale de secret, sans indirection vers une source de configuration externe et sécurisée.

## Correction

- Charger tous les secrets depuis les variables d'environnement (`.env` Laravel, non versionné) ou un gestionnaire de secrets (Vault, AWS/Azure/GCP Secret Manager).
- Lever une exception explicite si un secret requis est absent, plutôt que d'utiliser une valeur par défaut silencieuse.
- Ajouter `.env` à `.gitignore` et vérifier qu'aucun secret n'est déjà présent dans l'historique.
- Révoquer et régénérer immédiatement tout secret ayant été exposé dans le code ou l'historique Git.
- Mettre en place une rotation périodique des secrets critiques.

## Références

- CWE-798: Use of Hard-coded Credentials
- OWASP A02:2021 - Cryptographic Failures
- OWASP Secrets Management Cheat Sheet
