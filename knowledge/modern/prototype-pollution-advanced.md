---
id: prototype-pollution-advanced
category: modern
cwe: CWE-1321
owasp: A03:2021-Injection
severity_default: high
languages: [js]
---

# Prototype Pollution avancée

## Description
La pollution de prototype JavaScript survient lorsqu'une fusion (`merge`), un clonage profond ou une désérialisation d'objet permet à un attaquant de contrôler des clés comme `__proto__`, `constructor.prototype` ou `prototype`, modifiant ainsi le prototype global `Object.prototype`. Dans sa forme avancée, cette altération se propage à l'ensemble de l'application : elle peut désactiver des vérifications de sécurité implicites, injecter des propriétés dans des objets de configuration internes, ou, combinée à un gadget d'exécution (ex: un moteur de template ou un `child_process`), mener à une exécution de code arbitraire côté serveur.

## Où ça apparaît typiquement
- Fonctions utilitaires de fusion/clonage d'objets maison ou de bibliothèques non patchées (anciennes versions de `lodash.merge`, `deepmerge`, `jQuery.extend`).
- Désérialisation de JSON utilisateur directement fusionné dans un objet de configuration ou d'état applicatif.
- Frameworks côté serveur (Node.js) où un objet pollué au niveau global affecte toutes les requêtes suivantes du même processus.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fonction de fusion récursive d'objets sans exclusion explicite des clés `__proto__`, `constructor`, `prototype`.
- Dépendances de fusion/clonage (`lodash`, `merge`, `deepmerge`) utilisées avec une version antérieure aux correctifs connus de pollution de prototype.
- `JSON.parse` d'une entrée utilisateur directement fusionné dans un objet partagé entre requêtes sans validation de schéma.

## Remédiation
- Utiliser `Object.create(null)` ou des `Map` pour les objets construits à partir d'entrées utilisateur, plutôt que des objets littéraux.
- Filtrer explicitement les clés dangereuses (`__proto__`, `constructor`, `prototype`) dans toute fonction de fusion/assignation récursive.
- Maintenir à jour les bibliothèques de fusion/clonage et valider le schéma des objets désérialisés avant utilisation.
- Voir `rules/remediation/prototype-pollution-advanced.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/prototype-pollution-advanced/` (à créer selon le même schéma).

## Références
- OWASP: Prototype Pollution Prevention Cheat Sheet
- CWE-1321: Improperly Controlled Modification of Object Prototype Attributes ('Prototype Pollution')
