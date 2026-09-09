---
id: mutation-xss
category: xss
cwe: CWE-79
owasp: A03:2021-Injection
severity_default: high
languages: [js]
---

# Mutation-based Cross-Site Scripting (mXSS)

## Description
Le XSS par mutation exploite les différences de comportement entre l'étape de sanitisation d'un contenu HTML et son interprétation finale par le moteur de rendu du navigateur (parser HTML). Un contenu jugé sûr par le sanitiseur peut être réécrit ("muté") par le navigateur lors de son insertion dans le DOM (notamment via des allers-retours `innerHTML`), faisant réapparaître une structure active après la vérification, ce qui contourne la protection même lorsque le code de sanitisation est correct en apparence.

## Où ça apparaît typiquement
- Sanitisation HTML effectuée sur une représentation sérialisée puis réinsérée dans le DOM, où le navigateur réinterprète et normalise le markup différemment.
- Éditeurs riches (WYSIWYG) qui font des allers-retours multiples entre modèle DOM et chaîne HTML.
- Bibliothèques de sanitisation obsolètes ou mal configurées ne tenant pas compte des quirks de parsing du navigateur cible.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Utilisation d'une bibliothèque de sanitisation HTML ancienne ou non maintenue, connue pour des contournements par mutation.
- Sanitisation effectuée côté serveur uniquement, sans revalidation côté client juste avant insertion dans le DOM via `innerHTML`.
- Allers-retours multiples de contenu entre chaîne HTML et DOM (parse → sérialise → reparse) sans re-sanitisation à chaque étape.

## Remédiation
- Utiliser une bibliothèque de sanitisation HTML activement maintenue et connue pour couvrir les vecteurs de mutation (mise à jour régulière).
- Limiter au strict minimum les allers-retours entre représentation chaîne et DOM pour du contenu non fiable.
- Appliquer une Content Security Policy stricte en défense en profondeur, indépendamment de la sanitisation.
- Voir `rules/remediation/mutation-xss.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/nodejs-express/mutation-xss/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Cross Site Scripting Prevention / DOM based XSS Prevention
- CWE-79: Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')
