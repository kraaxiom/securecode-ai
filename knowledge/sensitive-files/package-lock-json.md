---
id: package-lock-json
category: sensitive-files
cwe: CWE-200
owasp: A05:2021-Security Misconfiguration
severity_default: low
languages: []
---

# Exposition du fichier package-lock.json

## Description
`package-lock.json` fige les versions exactes de toutes les dépendances npm d'un projet Node.js, y compris les dépendances transitives. Sa divulgation publique permet à un attaquant de reconstituer précisément l'arbre de dépendances de l'application et de rechercher des vulnérabilités connues (CVE npm) affectant ces versions précises, sans avoir à deviner la stack technique utilisée.

## Où ça apparaît typiquement
- Fichier déployé à la racine du webroot avec le reste du projet Node.js.
- Absence de règle serveur bloquant les fichiers de manifeste/lock npm.
- Artefacts de build incluant les fichiers de métadonnées npm au lieu du seul code compilé.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- `package-lock.json` accessible directement via une requête HTTP sur le webroot déployé.
- Absence de séparation entre artefact de build (dist/) et fichiers source/métadonnées de dépendances.
- Dépendances listées avec des versions non auditées depuis longtemps.

## Remédiation
- Bloquer l'accès direct aux fichiers de manifeste npm (`package.json`, `package-lock.json`) au niveau du serveur web.
- Ne déployer que le résultat du build (dossier `dist`/`build`), jamais les sources ni les métadonnées npm.
- Exécuter `npm audit` régulièrement en CI pour détecter les dépendances vulnérables avant qu'elles ne soient exploitées.
- Voir `rules/remediation/package-lock-json.md` pour les diffs de configuration serveur.

## Exemple avant/après
Voir `examples/config/package-lock-json/` (configuration serveur avant/après blocage de l'accès).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
