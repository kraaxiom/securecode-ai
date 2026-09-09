---
id: code-injection
category: injections
cwe: CWE-94
owasp: A03:2021-Injection
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# Code Injection

## Description
L'injection de code survient lorsqu'une application évalue ou exécute dynamiquement une chaîne de caractères contrôlée, en tout ou partie, par un utilisateur (via `eval`, `exec`, désérialisation dynamique, ou chargement de code à la volée). Contrairement à l'injection de commande qui cible le shell système, l'injection de code cible directement l'interpréteur du langage de l'application, offrant à l'attaquant un contrôle total sur son exécution.

## Où ça apparaît typiquement
- Utilisation de `eval()`, `exec()`, `Function()`, `create_function()` avec une chaîne construite à partir d'une entrée utilisateur.
- Moteurs de règles métier ou de calcul de formules interprétant du code fourni par l'utilisateur.
- Désérialisation d'objets ou de scripts (YAML, pickle, PHP `unserialize`) provenant de sources non fiables.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel direct à `eval`/`exec`/`Function`/`vm.runInNewContext` avec une variable non constante.
- Concaténation d'entrée utilisateur dans une chaîne ensuite interprétée comme du code.
- Désérialisation de données non fiables sans liste blanche de types autorisés.

## Remédiation
- Éviter totalement l'évaluation dynamique de code à partir d'entrées utilisateur.
- Remplacer les mécanismes d'`eval` par des structures de données déclaratives (JSON, mapping de fonctions autorisées).
- Utiliser des désérialiseurs sûrs (ex: `JSON.parse` plutôt que `eval`, listes blanches de classes en Java/PHP).
- Voir `rules/remediation/code-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/code-injection/`.

## Références
- OWASP Cheat Sheet: Deserialization Cheat Sheet
- CWE-94: Improper Control of Generation of Code ('Code Injection')
