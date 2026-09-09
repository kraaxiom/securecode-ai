---
id: command-injection
category: injections
cwe: CWE-78
owasp: A03:2021-Injection
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# Command Injection

## Description
L'injection de commande se produit lorsqu'une application transmet à un interpréteur de commandes du système d'exploitation une chaîne construite à partir d'une entrée utilisateur non neutralisée. L'attaquant peut alors ajouter des métacaractères shell (`;`, `&&`, `|`, backticks) pour exécuter des commandes arbitraires avec les privilèges du processus applicatif.

## Où ça apparaît typiquement
- Appels à `system()`, `exec()`, `popen()`, `subprocess.call(..., shell=True)`, `Runtime.exec()` avec une chaîne construite dynamiquement.
- Fonctionnalités de conversion de fichiers, ping/traceroute réseau, ou intégration d'outils externes (ImageMagick, ffmpeg) pilotées par des paramètres utilisateur.
- Scripts d'administration exposés via une interface web qui relaient des arguments à un binaire système.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Concaténation d'une entrée utilisateur dans une chaîne de commande shell.
- Utilisation de `shell=True` (Python), `exec()`/`system()` (PHP, Node) ou équivalent avec paramètre dynamique.
- Absence de liste blanche stricte sur les valeurs autorisées (ex: nom de fichier, adresse IP).

## Remédiation
- Éviter l'appel à un shell ; utiliser des API d'exécution de processus avec arguments passés en tableau (pas de concaténation de chaîne).
- Valider les entrées avec une liste blanche stricte de caractères/valeurs autorisés.
- Exécuter le processus avec le moindre privilège possible (utilisateur dédié, sandbox, conteneur).
- Voir `rules/remediation/command-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/command-injection/`.

## Références
- OWASP Cheat Sheet: OS Command Injection Defense
- CWE-78: Improper Neutralization of Special Elements used in an OS Command ('OS Command Injection')
