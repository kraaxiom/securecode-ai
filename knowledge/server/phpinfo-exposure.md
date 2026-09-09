---
id: phpinfo-exposure
category: server
cwe: CWE-215
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: [php]
---

# Page phpinfo() exposée

## Description
Une page appelant `phpinfo()` accessible publiquement révèle une quantité importante d'informations sur l'environnement serveur : version exacte de PHP, modules chargés, chemins absolus, variables d'environnement (parfois y compris des identifiants injectés via des variables serveur), configuration de sécurité (`disable_functions`, `open_basedir`), et parfois des en-têtes de requêtes précédentes. Ces informations facilitent grandement la recherche de vulnérabilités connues correspondant à la stack exacte.

## Où ça apparaît typiquement
- Fichier `phpinfo.php`/`info.php` créé pour du débogage et oublié en production.
- Endpoint de diagnostic/healthcheck appelant `phpinfo()` sans authentification.
- Page d'installation d'un CMS/framework encore accessible après le déploiement initial.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fichier source contenant un appel direct à `phpinfo()` accessible sans authentification.
- Réponse HTTP correspondant au format de sortie caractéristique de `phpinfo()` (tableaux de configuration PHP).

## Remédiation
- Supprimer tout fichier de test appelant `phpinfo()` avant le déploiement en production.
- Si un diagnostic serveur est nécessaire, l'exposer derrière une authentification stricte et le restreindre par IP.
- Automatiser un contrôle en CI/CD interdisant la présence de fichiers de debug connus dans l'artefact de déploiement.
- Voir `rules/remediation/phpinfo-exposure.md`.

## Exemple avant/après
Voir `examples/php/phpinfo-exposure/`.

## Références
- OWASP Top 10 2021 — A05: Security Misconfiguration
- OWASP Testing Guide: Test for Default Credentials and Exposed Admin Interfaces
- CWE-215: Insertion of Sensitive Information Into Debugging Code
