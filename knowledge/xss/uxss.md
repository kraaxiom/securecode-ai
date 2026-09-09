---
id: uxss
category: xss
cwe: CWE-79
owasp: A03:2021-Injection
severity_default: critical
languages: [js]
---

# Universal Cross-Site Scripting (UXSS)

## Description
Le UXSS désigne une classe de vulnérabilités où une faille dans le navigateur lui-même, une extension, ou un composant tiers largement intégré (SDK publicitaire, widget embarqué) permet de contourner la Same-Origin Policy et d'exécuter du script dans le contexte de n'importe quel site visité par la victime, indépendamment d'une faille applicative propre au site ciblé. Bien que la cause première ne soit généralement pas le code de l'application, celle-ci reste responsable de limiter la surface d'exposition via les composants tiers qu'elle intègre.

## Où ça apparaît typiquement
- Intégration d'extensions de navigateur, de SDK tiers ou de widgets embarqués (chat, publicité, analytics) exécutés dans le contexte de la page.
- Iframes tierces avec des permissions excessives (`sandbox` absent ou trop permissif) intégrées dans l'application.
- Dépendance à des bibliothèques JavaScript tierces chargées depuis un CDN sans intégrité vérifiée.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Inclusion de scripts tiers sans attribut `integrity` (Subresource Integrity) ni restriction via Content Security Policy.
- Iframes tierces intégrées sans attribut `sandbox` restrictif adapté au besoin réel.
- Absence de politique de permissions (`Permissions-Policy`) limitant les capacités des composants tiers embarqués.

## Remédiation
- Restreindre strictement les sources de script autorisées via une Content Security Policy, et utiliser Subresource Integrity pour toute dépendance tierce chargée depuis un CDN.
- Isoler les composants tiers non indispensables dans des iframes avec attribut `sandbox` minimal.
- Maintenir à jour les navigateurs ciblés et suivre les avis de sécurité des composants tiers intégrés.
- Voir `rules/remediation/uxss.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/nodejs-express/uxss/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Content Security Policy Cheat Sheet
- CWE-79: Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')
