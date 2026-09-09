---
id: missing-httponly
category: sessions
cwe: CWE-1004
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: [php, js, python, java, csharp, go]
---

# Cookie sans attribut HttpOnly

## Description
L'attribut `HttpOnly` d'un cookie empêche son accès via JavaScript côté client (`document.cookie`). Quand un cookie de session (ou tout cookie sensible) est émis sans cet attribut, une faille XSS même mineure sur le site permet à un attaquant de voler directement le cookie et de détourner la session de la victime. C'est une défense en profondeur essentielle : elle ne corrige pas le XSS, mais réduit fortement son impact sur les sessions.

## Où ça apparaît typiquement
- Configuration manuelle de cookies de session (`setcookie`, `res.cookie`, `response.set_cookie`, `Set-Cookie` HTTP header) sans le flag `HttpOnly`.
- Frameworks configurés avec des paramètres de session par défaut non durcis.
- Cookies applicatifs personnalisés (tokens CSRF, préférences) qui n'ont pas besoin d'être lus en JS mais le sont par défaut.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à une fonction de création de cookie sans l'option `httponly: true` / `HttpOnly` explicite.
- En-tête `Set-Cookie` construit manuellement sans le mot-clé `HttpOnly`.
- Configuration de session du framework (`session.cookie_httponly`, `SESSION_COOKIE_HTTPONLY`) définie à `false` ou absente.

## Remédiation
- Définir systématiquement `HttpOnly` sur tout cookie de session ou cookie sensible qui n'a pas besoin d'être lu par du JavaScript.
- Activer les réglages globaux du framework pour appliquer `HttpOnly` par défaut à tous les cookies de session.
- Combiner avec `Secure` et `SameSite` pour une protection complète des cookies.
- Voir `rules/remediation/missing-httponly.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js/missing-httponly/`.

## Références
- OWASP Cheat Sheet: Session Management Cheat Sheet
- CWE-1004: Sensitive Cookie Without 'HttpOnly' Flag
