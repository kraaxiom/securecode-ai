---
id: verbose-errors
category: server
cwe: CWE-209
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Messages d'erreur verbeux exposés

## Description
Des messages d'erreur verbeux exposés au client révèlent des détails internes de l'application (nom et version du framework/serveur, requêtes SQL, chemins de fichiers, structure de la base de données) lorsqu'une exception non gérée survient. Même sans stack trace complète (voir `stack-trace`), un simple message d'erreur de bibliothèque ou de base de données peut suffire à un attaquant pour affiner ses tentatives d'exploitation.

## Où ça apparaît typiquement
- Erreurs de base de données renvoyées telles quelles au client (`SQLSTATE[...]`, `ORA-00933`, etc.).
- Messages d'exception de framework affichés dans la réponse HTTP au lieu d'une page d'erreur générique.
- API renvoyant le message d'exception brut dans le corps JSON de la réponse.
- Absence de gestion centralisée des erreurs (chaque route gère ses erreurs différemment, certaines fuient des détails).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Réponse HTTP en erreur (4xx/5xx) contenant un message technique détaillé (nom de driver DB, requête, chemin de fichier) plutôt qu'un message générique.
- Absence de gestionnaire d'erreurs global/middleware centralisant la mise en forme des réponses d'erreur.
- Messages d'exception de bibliothèques tierces propagés sans transformation jusqu'au client.

## Remédiation
- Mettre en place un gestionnaire d'erreurs centralisé qui renvoie des messages génériques au client et journalise les détails côté serveur uniquement.
- Ne jamais propager directement une exception de bibliothèque/DB dans la réponse HTTP.
- Définir des codes d'erreur applicatifs stables et documentés plutôt que des messages techniques bruts.
- Voir `rules/remediation/verbose-errors.md`.

## Exemple avant/après
Voir `examples/php/verbose-errors/`.

## Références
- OWASP Top 10 2021 — A05: Security Misconfiguration
- OWASP Cheat Sheet: Error Handling
- CWE-209: Generation of Error Message Containing Sensitive Information
