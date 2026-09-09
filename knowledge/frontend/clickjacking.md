---
id: clickjacking
category: frontend
cwe: CWE-1021
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: [js, ts, php, python, java]
---

# Clickjacking

## Description
Le clickjacking (détournement de clic) consiste à superposer une page légitime dans un `<iframe>` invisible ou transparent au-dessus d'un contenu leurre, pour amener l'utilisateur à cliquer sur des éléments de la page ciblée sans en avoir conscience. L'attaque exploite l'absence de restriction sur l'intégration de la page dans un cadre tiers. Elle permet de faire exécuter des actions sensibles (validation d'un paiement, changement de paramètres, like/follow) à l'insu de la victime.

## Où ça apparaît typiquement
- Pages d'administration, de paiement ou de changement de paramètres de compte servies sans en-tête anti-framing.
- Applications qui autorisent l'intégration en iframe pour des besoins légitimes (widgets, SSO) mais sans restreindre les origines autorisées.
- Configuration serveur ou CDN qui ne définit pas d'en-tête de sécurité par défaut sur l'ensemble des routes.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de l'en-tête HTTP `X-Frame-Options` (`DENY` ou `SAMEORIGIN`) dans les réponses des pages sensibles.
- Absence de directive `frame-ancestors` dans la Content-Security-Policy.
- Présence de code applicatif désactivant explicitement ces protections (middleware ou configuration qui supprime ces en-têtes).
- Absence de "frame-busting" JavaScript côté legacy sur les applications ne pouvant pas définir d'en-têtes.

## Remédiation
- Définir l'en-tête `X-Frame-Options: DENY` ou `SAMEORIGIN` sur toutes les réponses HTML sensibles.
- Ajouter une directive CSP `frame-ancestors 'none'` ou `frame-ancestors 'self'` (remplace `X-Frame-Options` sur les navigateurs modernes).
- Si l'intégration en iframe est nécessaire, restreindre explicitement `frame-ancestors` aux origines de confiance.
- Voir `rules/remediation/clickjacking.md` pour les diffs par langage/framework.

## Exemple avant/après
Voir `examples/js/clickjacking/`.

## Références
- OWASP Cheat Sheet: Clickjacking Defense Cheat Sheet
- CWE-1021: Improper Restriction of Rendered UI Layers or Frames
