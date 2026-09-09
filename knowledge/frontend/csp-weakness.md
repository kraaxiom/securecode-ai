---
id: csp-weakness
category: frontend
cwe: CWE-693
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: [js, ts, php, python, java]
---

# Faiblesse de Content Security Policy

## Description
La Content Security Policy (CSP) est un en-tête HTTP qui restreint les sources depuis lesquelles un navigateur peut charger et exécuter des scripts, styles ou autres ressources, réduisant l'impact des failles XSS. Une CSP faible — absente, trop permissive, ou mal configurée — ne remplit plus son rôle de mécanisme de défense en profondeur. Les erreurs courantes incluent l'usage de `unsafe-inline`, `unsafe-eval`, de jokers larges (`*`), ou l'absence pure et simple de l'en-tête.

## Où ça apparaît typiquement
- Configuration serveur/CDN ne définissant aucun en-tête `Content-Security-Policy`.
- Politiques générées automatiquement par des frameworks avec `script-src 'unsafe-inline' 'unsafe-eval'` laissé en production pour la simplicité de développement.
- Directives `default-src *` ou `script-src *` trop larges.
- CSP définie uniquement en mode `report-only` sans jamais passer en mode bloquant.
- Absence de directive `object-src 'none'` ou `base-uri 'self'`, laissant des vecteurs de contournement.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- En-tête `Content-Security-Policy` absent des réponses HTML.
- Présence de `unsafe-inline` et/ou `unsafe-eval` dans `script-src` en environnement de production.
- Directives `default-src`, `script-src` ou `style-src` définies avec un joker (`*`) trop large.
- Absence de `object-src 'none'` et `base-uri 'self'`.
- CSP restant en `Content-Security-Policy-Report-Only` sans jamais être appliquée en mode bloquant.

## Remédiation
- Définir une CSP restrictive basée sur une liste blanche explicite de sources de confiance.
- Éviter `unsafe-inline`/`unsafe-eval`; privilégier les nonces ou hashes pour les scripts inline légitimes.
- Ajouter `object-src 'none'` et `base-uri 'self'` par défaut.
- Tester la politique en mode `report-only` avant de l'appliquer strictement, puis migrer en mode bloquant.
- Voir `rules/remediation/csp-weakness.md` pour les diffs par langage/framework.

## Exemple avant/après
Voir `examples/js/csp-weakness/`.

## Références
- OWASP Cheat Sheet: Content Security Policy Cheat Sheet
- CWE-693: Protection Mechanism Failure
