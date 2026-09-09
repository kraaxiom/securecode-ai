---
id: stack-trace
category: server
cwe: CWE-209
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Stack trace exposée au client

## Description
L'exposition d'une stack trace complète dans une réponse HTTP d'erreur révèle la structure interne de l'application : noms de classes, chemins absolus sur le serveur, versions de bibliothèques, et parfois des fragments de code source ou de requêtes. Ces informations facilitent grandement la reconnaissance pour un attaquant qui cherche des vulnérabilités connues correspondant aux versions de composants identifiées.

## Où ça apparaît typiquement
- Exceptions non interceptées remontant jusqu'au serveur web, qui affiche la trace par défaut (page d'erreur de développement laissée active).
- API renvoyant le champ `stackTrace`/`trace` dans le corps JSON d'une réponse d'erreur.
- Frameworks avec page d'erreur "whitelabel" ou de debug non désactivée en production.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Réponse HTTP contenant des lignes de trace typiques (chemins de fichiers `.php`/`.py`/`.java`, numéros de ligne, noms de méthodes internes).
- Champ `trace`/`stackTrace`/`exception` présent dans un payload JSON d'erreur retourné au client.
- Configuration de gestion d'erreurs de framework encore sur le mode par défaut de développement.

## Remédiation
- Intercepter toutes les exceptions au niveau applicatif (middleware/handler global) et ne jamais renvoyer la trace au client.
- Journaliser la stack trace complète uniquement dans les logs serveur, avec un identifiant de corrélation renvoyé au client pour le support.
- Configurer explicitement le mode production de chaque framework pour désactiver l'affichage des traces.
- Voir `rules/remediation/stack-trace.md`.

## Exemple avant/après
Voir `examples/java/stack-trace/`.

## Références
- OWASP Top 10 2021 — A05: Security Misconfiguration
- OWASP Cheat Sheet: Error Handling
- CWE-209: Generation of Error Message Containing Sensitive Information
