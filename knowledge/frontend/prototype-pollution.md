---
id: prototype-pollution
category: frontend
cwe: CWE-1321
owasp: A03:2021-Injection
severity_default: high
languages: [js, ts]
---

# Prototype Pollution

## Description
La pollution de prototype est une vulnérabilité propre à JavaScript où un attaquant parvient à injecter ou modifier des propriétés sur `Object.prototype` (ou un autre prototype partagé), généralement via une fusion (`merge`), un clonage profond ou un parsing JSON non contrôlé de données utilisateur contenant des clés spéciales comme `__proto__`, `constructor` ou `prototype`. Une fois le prototype global pollué, tous les objets de l'application héritent de la propriété injectée, ce qui peut mener à un déni de service, un contournement de logique de sécurité, voire dans certains cas une exécution de code selon la façon dont la propriété polluée est ensuite utilisée.

## Où ça apparaît typiquement
- Fonctions de fusion d'objets maison ou de bibliothèques (`merge`, `extend`, `deepMerge`) appliquées à un objet JSON venant de l'utilisateur (body de requête, query string).
- Parsing de paramètres d'URL en objets imbriqués (ex: `qs.parse` avec options laxistes) sans filtrage des clés.
- Frameworks de templating ou de configuration acceptant des chemins de clés dynamiques dérivés d'entrée utilisateur (`lodash.set(obj, userPath, value)`).
- Désérialisation JSON suivie d'une assignation récursive de propriétés sans liste blanche de clés autorisées.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fonction de fusion/assignation récursive qui itère sur les clés d'un objet utilisateur sans exclure explicitement `__proto__`, `constructor` et `prototype`.
- Utilisation de `lodash.merge`, `lodash.set`, ou équivalents avec un chemin/objet dérivé directement d'une entrée utilisateur.
- Absence de `Object.create(null)` ou de `Map` pour stocker des données utilisateur sous forme de dictionnaire dynamique.
- Versions de bibliothèques connues pour des CVE de pollution de prototype non mises à jour (à croiser avec un scan de dépendances).

## Remédiation
- Filtrer explicitement les clés dangereuses (`__proto__`, `constructor`, `prototype`) lors de toute fusion/assignation dynamique d'objets issus d'entrée utilisateur.
- Utiliser `Object.create(null)` ou `Map` pour les structures de données indexées par des clés utilisateur.
- Geler les prototypes critiques (`Object.freeze(Object.prototype)`) en défense en profondeur si l'architecture le permet.
- Maintenir à jour les bibliothèques de manipulation d'objets et surveiller leurs CVE.
- Voir `rules/remediation/prototype-pollution.md` pour les diffs par langage/framework.

## Exemple avant/après
Voir `examples/js/prototype-pollution/`.

## Références
- OWASP Cheat Sheet: Prototype Pollution Prevention Cheat Sheet
- CWE-1321: Improperly Controlled Modification of Object Prototype Attributes ('Prototype Pollution')
