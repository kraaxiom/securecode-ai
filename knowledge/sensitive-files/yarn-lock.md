---
id: yarn-lock
category: sensitive-files
cwe: CWE-200
owasp: A05:2021-Security Misconfiguration
severity_default: low
languages: []
---

# Exposition du fichier yarn.lock

## Description
`yarn.lock` joue le même rôle que `package-lock.json` pour les projets utilisant le gestionnaire de paquets Yarn : il fige les versions exactes de toutes les dépendances, y compris transitives. Sa divulgation publique permet à un attaquant de cartographier précisément la stack de dépendances du projet et de croiser ces versions avec des vulnérabilités connues, sans effort de reconnaissance supplémentaire.

## Où ça apparaît typiquement
- Fichier déployé à la racine du webroot avec le reste du projet Node.js.
- Absence de règle serveur bloquant les fichiers de manifeste/lock Yarn.
- Artefacts de déploiement incluant les fichiers source/métadonnées au lieu du seul build compilé.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- `yarn.lock` accessible directement via une requête HTTP sur le webroot déployé.
- Absence de séparation entre artefact de build et fichiers de métadonnées de dépendances.
- Dépendances listées avec des versions anciennes non auditées.

## Remédiation
- Bloquer l'accès direct aux fichiers de manifeste Yarn (`package.json`, `yarn.lock`) au niveau du serveur web.
- Ne déployer que le résultat du build, jamais les sources ni les métadonnées de dépendances.
- Exécuter `yarn audit` régulièrement en CI pour détecter les dépendances vulnérables.
- Voir `rules/remediation/yarn-lock.md` pour les diffs de configuration serveur.

## Exemple avant/après
Voir `examples/config/yarn-lock/` (configuration serveur avant/après blocage de l'accès).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
