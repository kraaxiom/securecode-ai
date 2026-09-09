---
id: missing-hsts
category: headers
cwe: CWE-319
owasp: A02:2021-Cryptographic Failures
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Strict-Transport-Security (HSTS) manquant

## Description
L'absence d'en-tête `Strict-Transport-Security` laisse la porte ouverte à des attaques de type downgrade/SSL stripping : un utilisateur qui tape l'URL sans `https://` ou clique sur un lien HTTP sera d'abord servi en clair, moment que peut exploiter un attaquant en position de man-in-the-middle pour intercepter la connexion avant toute redirection vers HTTPS. HSTS force le navigateur à toujours contacter le site en HTTPS après une première visite réussie.

## Où ça apparaît typiquement
- Sites servant du contenu en HTTPS mais sans jamais renvoyer l'en-tête `Strict-Transport-Security`.
- Redirection HTTP → HTTPS mise en place côté serveur, mais sans HSTS pour empêcher la première requête en clair d'être interceptée.
- Configuration reverse proxy/CDN où l'en-tête est supprimé ou non transmis en amont de l'application.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- En-tête `Strict-Transport-Security` absent des réponses HTTPS de l'application.
- Configuration serveur web sans directive `add_header Strict-Transport-Security`.
- Absence d'inscription du domaine dans la liste de préchargement HSTS (`hstspreload.org`) alors que le site est critique.

## Remédiation
- Ajouter l'en-tête `Strict-Transport-Security: max-age=31536000; includeSubDomains` sur toutes les réponses HTTPS.
- N'activer `preload` qu'après validation que tous les sous-domaines supportent HTTPS de façon fiable.
- S'assurer que la redirection HTTP → HTTPS est en place en complément, pas en remplacement, de HSTS.
- Voir `rules/remediation/missing-hsts.md`.

## Exemple avant/après
Voir `examples/js/missing-hsts/`.

## Références
- OWASP Secure Headers Project
- RFC 6797: HTTP Strict Transport Security (HSTS)
- CWE-319: Cleartext Transmission of Sensitive Information
