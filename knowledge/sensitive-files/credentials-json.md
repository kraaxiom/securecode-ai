---
id: credentials-json
category: sensitive-files
cwe: CWE-538
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Exposition du fichier credentials.json

## Description
`credentials.json` est un nom de fichier standard utilisé par de nombreux SDK cloud (Google Cloud, Firebase, OAuth) pour stocker des identifiants de service : clés privées, client secrets, jetons d'accès. Lorsqu'il est déployé par erreur dans un répertoire accessible publiquement, il permet à un attaquant d'usurper l'identité de l'application auprès des services cloud concernés, avec les mêmes permissions que le compte de service légitime.

## Où ça apparaît typiquement
- Copie du fichier d'identifiants du SDK cloud laissée dans le répertoire de déploiement ou l'archive de build.
- Scripts d'automatisation (CI/CD) qui écrivent temporairement le fichier dans un répertoire servi par erreur.
- Dépôt de code où le fichier a été commité puis inclus dans l'artefact de déploiement.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- `credentials.json` présent dans l'arborescence déployée en production plutôt que monté via un secret manager.
- Fichier absent du `.gitignore`/`.dockerignore` alors qu'il est utilisé localement.
- Aucune restriction d'accès (IAM, ACL) supplémentaire sur le compte de service associé, aggravant l'impact d'une fuite.

## Remédiation
- Ne jamais inclure `credentials.json` dans l'image/artefact déployé ; l'injecter via un secret manager ou une identité de service managée (Workload Identity, IAM Role).
- Exclure systématiquement ce fichier du contrôle de version et des artefacts de build.
- Révoquer et régénérer les identifiants immédiatement en cas d'exposition, et auditer les accès effectués avec ces identifiants.
- Voir `rules/remediation/credentials-json.md` pour les diffs par configuration serveur/CI.

## Exemple avant/après
Voir `examples/config/credentials-json/` (pipeline CI avant/après suppression du fichier de l'artefact).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-538: Insertion of Sensitive Information into Externally-Accessible File or Directory
