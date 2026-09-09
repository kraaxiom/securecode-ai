---
id: missing-permissions-policy
category: headers
cwe: CWE-1021
owasp: A05:2021-Security Misconfiguration
severity_default: low
languages: [php, js, python, java, csharp, go, rust]
---

# Permissions-Policy manquant

## Description
L'absence d'en-tête `Permissions-Policy` (anciennement `Feature-Policy`) laisse toutes les API sensibles du navigateur (caméra, micro, géolocalisation, capteurs, autoplay) accessibles par défaut à la page et à tout contenu tiers embarqué (iframes, scripts publicitaires). En cas de compromission par XSS ou d'intégration d'un widget tiers malveillant, ces API peuvent être détournées pour espionner l'utilisateur.

## Où ça apparaît typiquement
- Applications intégrant des iframes ou scripts tiers (publicité, widgets) sans restreindre leurs permissions.
- Sites n'utilisant aucune des API sensibles mais ne les désactivant pas explicitement, laissant une surface d'attaque inutile en cas de XSS.
- Absence totale de politique de permissions dans la configuration serveur/application.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- En-tête `Permissions-Policy` absent des réponses HTTP.
- Iframes tierces intégrées sans attribut `allow` restrictif.

## Remédiation
- Définir un `Permissions-Policy` restrictif désactivant les fonctionnalités non utilisées (`camera=(), microphone=(), geolocation=()`, etc.).
- Restreindre explicitement les permissions accordées aux iframes tierces via l'attribut `allow` en complément de l'en-tête serveur.
- Revoir régulièrement la liste des fonctionnalités autorisées à mesure que l'application évolue.
- Voir `rules/remediation/missing-permissions-policy.md`.

## Exemple avant/après
Voir `examples/js/missing-permissions-policy/`.

## Références
- OWASP Secure Headers Project
- W3C Permissions Policy specification
- CWE-1021: Improper Restriction of Rendered UI Layers or Frames
