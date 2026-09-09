---
id: default-credentials
category: auth
cwe: CWE-1392
owasp: A07:2021-Identification and Authentication Failures
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Default Credentials

## Description
Cette vulnérabilité survient quand une application, un service ou un équipement conserve des identifiants par défaut fournis par l'éditeur (compte admin/admin, mot de passe d'installation, clé API d'exemple) sans forcer leur changement à la mise en production. Ces identifiants étant publiquement documentés, un attaquant peut les essayer directement pour obtenir un accès administrateur sans avoir besoin de deviner ou de casser quoi que ce soit.

## Où ça apparaît typiquement
- Panneaux d'administration de CMS, bases de données ou outils de monitoring déployés avec leur configuration d'origine.
- Comptes de service créés automatiquement par des scripts de provisionnement avec un mot de passe fixe documenté.
- Environnements de démonstration/staging accidentellement exposés en production avec les identifiants d'exemple.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Présence d'identifiants en dur dans des fichiers de configuration, scripts de déploiement ou données de seed, réutilisés tels quels en production.
- Absence de mécanisme forçant le changement du mot de passe par défaut à la première connexion.
- Documentation ou configuration listant des comptes standards (`admin`, `root`, `test`) sans rotation obligatoire.

## Remédiation
- Générer un mot de passe unique et aléatoire à l'installation, et forcer son changement à la première connexion.
- Supprimer ou désactiver systématiquement tout compte de démonstration avant mise en production.
- Auditer périodiquement les comptes de service et équipements pour détecter des identifiants encore par défaut.
- Voir `rules/remediation/default-credentials.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/default-credentials/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Authentication
- CWE-1392: Use of Default Credentials
