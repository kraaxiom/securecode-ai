---
id: dom-clobbering
category: frontend
cwe: CWE-79
owasp: A03:2021-Injection
severity_default: medium
languages: [js, ts]
---

# DOM Clobbering

## Description
Le DOM clobbering est une technique où un attaquant injecte des éléments HTML (souvent via un point d'injection HTML limité qui ne permet pas d'exécuter du JavaScript directement) dont les attributs `id` ou `name` viennent "écraser" des variables globales JavaScript ou des propriétés d'objets DOM attendues par l'application. Un script qui référence `window.config` ou `document.getElementById('foo').value` sans vérification de type peut alors se voir substituer un élément HTML contrôlé par l'attaquant, détournant la logique applicative.

## Où ça apparaît typiquement
- Applications autorisant un sous-ensemble de HTML utilisateur (commentaires, contenu riche) sans sanitisation stricte des attributs `id`/`name`.
- Scripts JavaScript référençant des variables globales implicites (`window.X`, `document.X`) censées être définies uniquement en JS, sans vérifier leur type avant usage.
- Code legacy s'appuyant sur des identifiants globaux de formulaire (`document.forms.login.action`) accessibles par simple nommage HTML.
- Bibliothèques tierces vérifiant l'existence d'une configuration globale (`if (window.appConfig) {...}`) sans valider qu'il s'agit bien d'un objet JS.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Utilisation de variables globales implicites (`window.NOM`) sans validation de type (`typeof x === 'object'`) avant usage sensible.
- Sanitisation HTML qui autorise les attributs `id` et `name` sans restriction sur les valeurs réservées utilisées par le code applicatif.
- Code JS qui fait confiance à `document.getElementById(...)` pour récupérer une configuration ou une URL sans vérifier qu'il s'agit d'une valeur attendue.
- Absence de préfixage ou d'espace de nommage dédié pour les identifiants DOM afin d'éviter les collisions avec des variables globales sensibles.

## Remédiation
- Toujours valider le type des objets globaux avant utilisation (`instanceof`, vérification de forme explicite) plutôt que de faire confiance à leur simple présence.
- Éviter de dépendre de variables globales implicites créées automatiquement par le DOM (named property access).
- Sanitiser strictement le HTML utilisateur, notamment en filtrant/normalisant les attributs `id` et `name`.
- Utiliser des modules JS encapsulés (scope local) plutôt que des variables globales pour la configuration sensible.
- Voir `rules/remediation/dom-clobbering.md` pour les diffs par langage/framework.

## Exemple avant/après
Voir `examples/js/dom-clobbering/`.

## Références
- OWASP Cheat Sheet: DOM based XSS Prevention Cheat Sheet
- CWE-79: Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')
