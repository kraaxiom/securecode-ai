---
id: error-sqli
category: injections
cwe: CWE-89
owasp: A03:2021-Injection
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Error-based SQL Injection

## Description
L'injection SQL basée sur les erreurs exploite les messages d'erreur détaillés renvoyés par le serveur de base de données (ou relayés par l'application) pour extraire des informations. En provoquant volontairement une erreur de syntaxe ou de type contenant le résultat d'une sous-requête, l'attaquant récupère des données directement dans le message d'erreur affiché au client.

## Où ça apparaît typiquement
- Environnements de développement ou mal configurés où les erreurs SQL brutes sont renvoyées au client (stack trace, message du driver).
- Requêtes construites par concaténation, particulièrement vulnérables à des erreurs de syntaxe déclenchables.
- APIs qui journalisent ou renvoient le message d'exception complet dans la réponse HTTP.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Configuration `display_errors`/`debug=True` activée en production, ou gestion d'exception qui renvoie `e.getMessage()` au client.
- Requêtes SQL construites par concaténation sans requêtes préparées.
- Absence de page d'erreur générique pour les échecs de requête base de données.

## Remédiation
- Désactiver l'affichage des erreurs détaillées en production ; renvoyer des messages génériques au client et journaliser le détail côté serveur uniquement.
- Utiliser systématiquement des requêtes préparées avec paramètres liés.
- Mettre en place une gestion d'exception centralisée qui ne propage jamais le message natif du driver SQL au client.
- Voir `rules/remediation/error-sqli.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/error-sqli/`.

## Références
- OWASP Cheat Sheet: SQL Injection Prevention
- CWE-89: Improper Neutralization of Special Elements used in an SQL Command
