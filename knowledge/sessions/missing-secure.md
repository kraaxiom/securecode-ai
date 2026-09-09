---
id: missing-secure
category: sessions
cwe: CWE-614
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: [php, js, python, java, csharp, go]
---

# Cookie sans attribut Secure

## Description
L'attribut `Secure` d'un cookie indique au navigateur de ne l'envoyer que sur des connexions HTTPS. Quand un cookie de session en est dépourvu, il peut être transmis en clair sur une connexion HTTP, par exemple lors d'une redirection non forcée ou d'un accès accidentel via une URL `http://`, exposant l'identifiant de session à toute personne capable d'observer le trafic réseau (attaque de type man-in-the-middle sur un réseau non fiable). Le risque est particulièrement élevé pour les applications qui ne forcent pas HTTPS de bout en bout.

## Où ça apparaît typiquement
- Création manuelle de cookies de session (`setcookie`, `res.cookie`, `response.set_cookie`, en-tête `Set-Cookie`) sans option `Secure`.
- Environnements de développement où le flag `Secure` est désactivé pour tester en HTTP, et qui se retrouve oublié en production.
- Configuration de session du framework non durcie pour l'environnement de production.
- Applications servies à la fois en HTTP et HTTPS sans redirection stricte vers HTTPS.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel de création de cookie sans l'option `secure: true` / `Secure` explicite.
- En-tête `Set-Cookie` construit manuellement sans le mot-clé `Secure`.
- Configuration de session du framework (`session.cookie_secure`, `SESSION_COOKIE_SECURE`) définie à `false` ou absente en configuration de production.
- Absence de mécanisme de redirection HTTP vers HTTPS (HSTS) combinée à un cookie non `Secure`.

## Remédiation
- Définir systématiquement `Secure` sur tout cookie de session en production.
- Forcer HTTPS sur l'ensemble de l'application (redirection + en-tête HSTS) pour rendre le flag `Secure` pleinement efficace.
- Séparer les configurations d'environnement pour ne désactiver `Secure` qu'en développement local, jamais en production.
- Combiner avec `HttpOnly` et `SameSite` pour une protection complète des cookies.
- Voir `rules/remediation/missing-secure.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js/missing-secure/`.

## Références
- OWASP Cheat Sheet: Session Management Cheat Sheet
- CWE-614: Sensitive Cookie in HTTPS Session Without 'Secure' Attribute
