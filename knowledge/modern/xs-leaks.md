---
id: xs-leaks
category: modern
cwe: CWE-203
owasp: A01:2021-Broken Access Control
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# XS-Leaks (Cross-Site Leaks)

## Description
Les XS-Leaks regroupent une famille de techniques permettant à un site malveillant d'inférer des informations sur l'état d'authentification ou le contenu d'un autre site, sans violer directement la politique de même origine, en observant des effets de bord détectables cross-origin : temps de réponse, présence/absence d'une redirection, taille de frame, nombre d'entrées dans l'historique, ou déclenchement d'événements `onload`/`onerror`. Ces fuites d'information peuvent révéler si une victime est connectée, appartient à un groupe spécifique, ou possède certaines données, sans jamais lire directement le contenu de la page ciblée.

## Où ça apparaît typiquement
- Pages dont le comportement (redirection, code de statut, présence de contenu) diffère selon l'état d'authentification, sans protection `Cross-Origin` adéquate.
- Endpoints qui renvoient une réponse de taille ou de temps de traitement variable selon des données sensibles (recherche, existence d'un compte).
- Absence des en-têtes `Cross-Origin-Opener-Policy`, `Cross-Origin-Resource-Policy` et `Cross-Origin-Embedder-Policy`.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence des en-têtes `Cross-Origin-Opener-Policy: same-origin` et `Cross-Origin-Resource-Policy: same-origin` sur les réponses sensibles.
- Endpoints dont le comportement observable (redirection vs 200, taille de réponse) varie directement selon l'état de session sans nécessité fonctionnelle.
- Absence de `SameSite` strict sur les cookies de session, augmentant la surface d'attaques combinées.

## Remédiation
- Définir `Cross-Origin-Opener-Policy: same-origin`, `Cross-Origin-Embedder-Policy: require-corp` et `Cross-Origin-Resource-Policy: same-origin` sur les réponses sensibles.
- Uniformiser les comportements observables (temps de réponse, codes de statut) entre états authentifié/non authentifié quand c'est possible.
- Utiliser `SameSite=Strict` ou `Lax` sur les cookies de session pour réduire la surface d'attaque cross-site.
- Voir `rules/remediation/xs-leaks.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/xs-leaks/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP: XS-Leaks Wiki (XS-Leaks Attacks)
- CWE-203: Observable Discrepancy
