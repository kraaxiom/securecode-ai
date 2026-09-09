---
id: missing-x-content-type-options
category: headers
cwe: CWE-436
owasp: A05:2021-Security Misconfiguration
severity_default: low
languages: [php, js, python, java, csharp, go, rust]
---

# X-Content-Type-Options manquant

## Description
L'absence de l'en-tête `X-Content-Type-Options: nosniff` permet à certains navigateurs de faire du "MIME sniffing" : ils ignorent le `Content-Type` déclaré par le serveur et devinent le type réel du contenu en l'inspectant. Un fichier uploadé par un utilisateur et déclaré comme `text/plain` ou `image/*` peut ainsi être réinterprété comme du HTML/JavaScript exécutable par le navigateur, ouvrant la voie à du XSS stocké via upload de fichier.

## Où ça apparaît typiquement
- Endpoints servant des fichiers uploadés par les utilisateurs (avatars, pièces jointes, documents) sans cet en-tête.
- Configuration serveur web ne forçant pas `nosniff` globalement.
- API renvoyant du JSON/texte sans en-tête de sécurité, potentiellement interprété différemment selon le navigateur.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- En-tête `X-Content-Type-Options` absent des réponses HTTP, en particulier sur les routes de téléchargement de fichiers.
- `Content-Type` des réponses de fichiers uploadés non strictement défini ou dérivé de l'extension fournie par l'utilisateur.

## Remédiation
- Ajouter `X-Content-Type-Options: nosniff` sur toutes les réponses HTTP, en particulier celles servant du contenu utilisateur.
- Définir un `Content-Type` explicite et correct pour chaque réponse, notamment pour les fichiers uploadés (validation par contenu réel, pas par extension).
- Servir les fichiers uploadés potentiellement dangereux depuis un domaine séparé sans cookies de session (isolation d'origine).
- Voir `rules/remediation/missing-x-content-type-options.md`.

## Exemple avant/après
Voir `examples/js/missing-x-content-type-options/`.

## Références
- OWASP Secure Headers Project
- MDN Web Docs: X-Content-Type-Options
- CWE-436: Interpretation Conflict
