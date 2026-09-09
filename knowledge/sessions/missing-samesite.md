---
id: missing-samesite
category: sessions
cwe: CWE-1275
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: [php, js, python, java, csharp, go]
---

# Cookie sans attribut SameSite

## Description
L'attribut `SameSite` d'un cookie indique au navigateur s'il doit envoyer le cookie lors de requêtes cross-site. Quand un cookie de session est émis sans cet attribut (ou avec `SameSite=None` non justifié), le navigateur continue de l'attacher automatiquement aux requêtes déclenchées depuis un autre site, ce qui expose l'application aux attaques CSRF et facilite certains scénarios de fuite d'informations cross-site. C'est une protection de défense en profondeur simple à mettre en place mais souvent oubliée dans les configurations par défaut.

## Où ça apparaît typiquement
- Création manuelle de cookies de session (`setcookie`, `res.cookie`, `response.set_cookie`, en-tête `Set-Cookie`) sans option `SameSite`.
- Frameworks web dont la configuration de session par défaut n'impose pas `SameSite=Lax` ou `Strict`.
- Cookies applicatifs sensibles (tokens, préférences d'authentification) hérités d'anciennes configurations antérieures à l'introduction de `SameSite`.
- Usage de `SameSite=None` sans attribut `Secure` associé ni justification fonctionnelle (widget intégré en iframe cross-site, etc.).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel de création de cookie sans option `samesite` / `SameSite` explicite.
- En-tête `Set-Cookie` construit manuellement sans le mot-clé `SameSite`.
- Configuration de session du framework (`session.cookie_samesite`, `SESSION_COOKIE_SAMESITE`) absente ou définie à `None`.
- Présence de `SameSite=None` sans `Secure` associé dans la même déclaration de cookie.

## Remédiation
- Définir `SameSite=Lax` par défaut pour les cookies de session, ou `Strict` quand aucune navigation cross-site légitime n'est nécessaire.
- Réserver `SameSite=None` aux cas justifiés et toujours l'associer à `Secure`.
- Activer le réglage global du framework pour appliquer `SameSite` par défaut à tous les cookies émis.
- Compléter par une protection CSRF applicative (token anti-CSRF) plutôt que de reposer uniquement sur `SameSite`.
- Voir `rules/remediation/missing-samesite.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js/missing-samesite/`.

## Références
- OWASP Cheat Sheet: Session Management Cheat Sheet
- CWE-1275: Sensitive Cookie with Improper SameSite Attribute
