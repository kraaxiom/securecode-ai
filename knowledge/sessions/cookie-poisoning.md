---
id: cookie-poisoning
category: sessions
cwe: CWE-565
owasp: A07:2021-Identification and Authentication Failures
severity_default: high
languages: [php, js, python, java, csharp, go]
---

# Cookie Poisoning

## Description
Le "cookie poisoning" désigne la manipulation par un attaquant du contenu d'un cookie (identifiant de session, rôle, prix, préférence) pour altérer le comportement de l'application côté serveur. Le problème survient quand l'application fait confiance à une valeur de cookie sans la valider, la signer ou la chiffrer, alors que le cookie est entièrement sous le contrôle du navigateur (donc de l'utilisateur). Un attaquant peut alors modifier son propre cookie pour s'octroyer des privilèges, contourner une logique métier ou usurper un état applicatif.

## Où ça apparaît typiquement
- Cookies applicatifs stockant directement un rôle, un ID utilisateur ou un flag (`role=admin`, `user_id=42`) en clair.
- Logique métier (panier, prix, quota) reposant sur une valeur de cookie non vérifiée côté serveur.
- Cookies "maison" utilisés comme mécanisme de session sans signature ni chiffrement (au lieu du gestionnaire de session natif du framework).
- Réutilisation d'une valeur de cookie dans une requête SQL, un chemin de fichier ou une commande sans validation.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Lecture directe de `$_COOKIE`, `req.cookies`, `request.COOKIES`, `HttpServletRequest.getCookies()` utilisée pour une décision d'autorisation ou métier, sans vérification de signature/intégrité.
- Cookie contenant des données structurées (JSON, paires clé=valeur) en clair sans HMAC ni chiffrement authentifié.
- Absence d'appel à une fonction de signature (`sign`, `itsdangerous`, `cookie-signature`) ou de framework de session côté serveur.
- Cookie utilisé comme unique source de vérité pour un état sensible, sans re-vérification en base de données à chaque requête sensible.

## Remédiation
- Ne jamais faire confiance à un cookie pour une décision de sécurité sans signature cryptographique (HMAC) ou chiffrement authentifié, avec vérification systématique côté serveur.
- Utiliser le mécanisme de session natif du framework (session store côté serveur, cookie ne contenant qu'un identifiant opaque).
- Revalider en base de données les informations sensibles (rôle, propriétaire de ressource) à chaque requête plutôt que de se fier au cookie.
- Voir `rules/remediation/cookie-poisoning.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/cookie-poisoning/`.

## Références
- OWASP Cheat Sheet: Session Management Cheat Sheet
- CWE-565: Reliance on Cookies without Validation and Integrity Checking
