---
id: api-key-leakage
category: api
cwe: CWE-798
owasp: A02:2021-Cryptographic Failures
severity_default: high
languages: [php, js, python, java, csharp, go]
---

# Fuite de clé API

## Description
La fuite de clé API se produit lorsqu'un secret d'authentification (clé API, token, secret client) est codé en dur dans le code source, un fichier de configuration versionné, ou exposé côté client (JavaScript, mobile). Une fois publié dans un dépôt Git, une image Docker ou un bundle frontend, ce secret peut être extrait par toute personne y ayant accès, y compris via l'historique Git ou des outils de scan automatisé. Les conséquences vont de l'usurpation d'identité applicative à l'accès complet à des services tiers facturés (cloud, paiement, messagerie).

## Où ça apparaît typiquement
- Constantes ou variables codées en dur dans le code source (`API_KEY = "sk-..."`).
- Fichiers de configuration commités (`.env`, `config.php`, `appsettings.json`) sans exclusion `.gitignore`.
- Clés injectées côté client dans du JavaScript, des apps mobiles ou des fichiers de configuration front exposés publiquement.
- Logs applicatifs affichant des en-têtes de requête ou des payloads contenant des clés.
- Images Docker ou artefacts CI/CD embarquant des secrets en variable d'environnement figée dans le layer.
- Dépôts publics ou forks accidentels contenant l'historique complet des commits.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Chaînes littérales ressemblant à des clés (préfixes connus type `sk_`, `AKIA`, `ghp_`, longueur/entropie élevée) assignées directement dans le code.
- Fichiers `.env`, `*.pem`, `*.key` présents dans l'arborescence versionnée sans entrée correspondante dans `.gitignore`.
- Absence d'appel à un gestionnaire de secrets (vault, variables d'environnement injectées à l'exécution, secret manager cloud).
- Clés API visibles dans le bundle JavaScript compilé livré au navigateur.
- Commentaires ou messages de commit mentionnant explicitement des clés ou tokens.

## Remédiation
- Ne jamais coder en dur de secrets : les injecter via variables d'environnement ou un gestionnaire de secrets dédié (Vault, AWS Secrets Manager, Azure Key Vault).
- Ajouter systématiquement les fichiers sensibles au `.gitignore` et purger l'historique Git en cas de fuite déjà commise (rotation immédiate de la clé concernée).
- Séparer strictement les clés utilisables côté client (scope minimal, restrictions de domaine/IP) des clés serveur à privilège élevé.
- Mettre en place une rotation régulière des clés et une détection automatisée de secrets (pre-commit hooks, scanners CI).
- Voir `rules/remediation/api-key-leakage.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/api-key-leakage/` (et les répertoires équivalents pour js, python, java, csharp, go).

## Références
- OWASP Cheat Sheet: Secrets Management
- CWE-798: Use of Hard-coded Credentials
