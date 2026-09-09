---
id: os-command-injection
category: injections
cwe: CWE-78
owasp: A03:2021-Injection
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# OS Command Injection

## Description
L'injection de commande système survient lorsqu'une entrée utilisateur non neutralisée est intégrée dans un appel exécutant une commande shell, permettant à un attaquant d'injecter des méta-caractères shell (`;`, `|`, `&&`, backticks) pour exécuter des commandes arbitraires sur le serveur. C'est l'une des vulnérabilités les plus critiques car elle mène généralement à une exécution de code à distance complète.

## Où ça apparaît typiquement
- Appels à des fonctions d'exécution shell (`exec`, `system`, `popen`, `subprocess.run(shell=True)`, `Runtime.exec`) avec une chaîne construite à partir d'entrées utilisateur.
- Fonctionnalités de conversion de fichiers, ping, traitement d'image ou d'appel à des outils externes en ligne de commande.
- Scripts d'automatisation ou de build recevant des paramètres externes non validés.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Concaténation de variables non filtrées dans une chaîne passée à une fonction d'exécution shell.
- Utilisation de l'option shell activée (`shell=True`, `exec()` de chaîne complète) alors qu'un appel direct au binaire suffirait.
- Absence de liste blanche stricte sur les valeurs transmises en argument de commande.

## Remédiation
- Éviter tout appel shell avec entrée utilisateur; privilégier les API d'exécution qui prennent un tableau d'arguments (pas d'interprétation shell).
- Valider les entrées selon une liste blanche stricte (ex: extension de fichier, format numérique).
- Exécuter les processus externes avec les privilèges minimaux nécessaires (sandboxing, utilisateur dédié).
- Voir `rules/remediation/os-command-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/python-flask/os-command-injection/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: OS Command Injection Defense
- CWE-78: Improper Neutralization of Special Elements used in an OS Command ('OS Command Injection')
