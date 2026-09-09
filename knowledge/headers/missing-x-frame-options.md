---
id: missing-x-frame-options
category: headers
cwe: CWE-1021
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# X-Frame-Options manquant

## Description
L'absence de l'en-tête `X-Frame-Options` (ou d'une directive `frame-ancestors` en CSP) permet à un site tiers malveillant d'intégrer la page cible dans une `<iframe>` et de mener une attaque de clickjacking : l'utilisateur croit interagir avec l'interface visible mais ses clics sont en réalité capturés par la page piégée pour déclencher des actions non voulues sur le site légitime (validation de formulaire, changement de paramètres, etc.).

## Où ça apparaît typiquement
- Pages contenant des actions sensibles (changement de mot de passe, transfert, validation) servies sans protection anti-framing.
- Configuration serveur web ou application ne définissant jamais `X-Frame-Options` ni `Content-Security-Policy: frame-ancestors`.
- Applications legacy n'ayant jamais intégré les en-têtes de sécurité modernes.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- En-tête `X-Frame-Options` absent des réponses HTTP.
- Absence de directive `frame-ancestors` dans une éventuelle CSP existante.
- Aucun test automatisé ou middleware de sécurité vérifiant les en-têtes anti-clickjacking.

## Remédiation
- Ajouter `X-Frame-Options: DENY` ou `SAMEORIGIN` selon le besoin métier (rarement besoin d'autoriser le framing cross-origin).
- Préférer à terme la directive CSP `frame-ancestors 'self'` (ou liste explicite de domaines), plus flexible et standard.
- Appliquer ces en-têtes de façon globale via middleware/reverse proxy plutôt que page par page.
- Voir `rules/remediation/missing-x-frame-options.md`.

## Exemple avant/après
Voir `examples/js/missing-x-frame-options/`.

## Références
- OWASP Cheat Sheet: Clickjacking Defense
- OWASP Secure Headers Project
- CWE-1021: Improper Restriction of Rendered UI Layers or Frames
