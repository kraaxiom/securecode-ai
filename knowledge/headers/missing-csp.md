---
id: missing-csp
category: headers
cwe: CWE-1021
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Content-Security-Policy manquant

## Description
L'absence d'en-tête `Content-Security-Policy` (CSP) prive le navigateur d'une directive lui indiquant quelles sources de scripts, styles, images ou connexions sont autorisées sur la page. Sans CSP, une faille XSS exploitable devient bien plus dangereuse : le navigateur exécutera n'importe quel script injecté sans restriction, alors qu'une CSP correctement configurée peut empêcher l'exécution de scripts inline ou provenant de domaines non autorisés.

## Où ça apparaît typiquement
- Réponses HTTP de l'application ne définissant jamais l'en-tête `Content-Security-Policy`.
- Configuration serveur web (Nginx, Apache, IIS) sans directive d'ajout d'en-tête de sécurité.
- Frameworks backend n'ayant pas de middleware de sécurité activé par défaut.
- Applications avec beaucoup de contenu tiers (widgets, analytics) où une CSP n'a jamais été mise en place par crainte de tout casser.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- En-tête `Content-Security-Policy` absent des réponses HTTP (vérifiable via les headers de réponse).
- Middleware de sécurité (helmet.js, secure_headers, etc.) absent ou non configuré dans le projet.
- Configuration serveur web sans directive `add_header Content-Security-Policy`.

## Remédiation
- Définir une CSP restrictive dès le départ (`default-src 'self'`) puis l'affiner par directive (`script-src`, `style-src`, `img-src`, `connect-src`).
- Éviter `'unsafe-inline'` et `'unsafe-eval'` ; utiliser des nonces ou hashes pour les scripts inline nécessaires.
- Déployer d'abord en mode `Content-Security-Policy-Report-Only` pour mesurer l'impact avant application stricte.
- Automatiser l'ajout de l'en-tête via un middleware ou la configuration du serveur web/reverse proxy.
- Voir `rules/remediation/missing-csp.md`.

## Exemple avant/après
Voir `examples/js/missing-csp/`.

## Références
- OWASP Secure Headers Project
- MDN Web Docs: Content-Security-Policy
- CWE-1021: Improper Restriction of Rendered UI Layers or Frames
