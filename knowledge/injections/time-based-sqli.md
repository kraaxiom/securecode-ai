---
id: time-based-sqli
category: injections
cwe: CWE-89
owasp: A03:2021-Injection
severity_default: high
languages: [php, js, python, java, csharp, go]
---

# Time-Based Blind SQL Injection

## Description
L'injection SQL aveugle basée sur le temps est une variante d'injection SQL où l'attaquant n'obtient aucune donnée ni différence visible dans la réponse, mais infère des informations en observant le délai de réponse du serveur après avoir injecté une condition provoquant une pause conditionnelle (ex: `SLEEP`, `WAITFOR DELAY`, `pg_sleep`). Elle est utilisée lorsque ni les données ni les messages d'erreur ne sont exposés, rendant la détection plus difficile.

## Où ça apparaît typiquement
- Mêmes points d'entrée que toute injection SQL classique (paramètres, filtres, tris) mais dans un contexte où l'application ne retourne ni données ni erreurs détaillées.
- API retournant des réponses génériques identiques quel que soit le résultat de la requête interne.
- Requêtes construites par concaténation à partir d'entrées utilisateur dans des applications avec gestion d'erreur "silencieuse".

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Concaténation de variable non échappée dans une chaîne SQL, combinée à une gestion d'erreur qui masque les détails de la base de données.
- Absence de limite de temps d'exécution (timeout) configurée côté application ou base de données pour les requêtes utilisateur.
- Journalisation insuffisante empêchant de détecter des pics de latence anormaux corrélés à des requêtes suspectes.

## Remédiation
- Utiliser systématiquement des requêtes préparées avec paramètres liés (élimine la classe de vulnérabilité, pas seulement la variante).
- Configurer des timeouts stricts d'exécution de requête côté base de données.
- Mettre en place une surveillance des temps de réponse anormaux comme indicateur de détection complémentaire.
- Voir `rules/remediation/time-based-sqli.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/python-flask/time-based-sqli/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: SQL Injection Prevention
- CWE-89: Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')
