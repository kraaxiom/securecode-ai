---
id: debug-mode
category: server
cwe: CWE-489
owasp: A05:2021-Security Misconfiguration
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Mode debug activé en production

## Description
Un mode debug activé en production laisse le framework afficher des informations de diagnostic détaillées (traces d'exécution, variables d'environnement, requêtes SQL, chemins de fichiers serveur, parfois une console interactive) directement dans les réponses HTTP en cas d'erreur. Cela offre à un attaquant une cartographie complète de l'application et, dans le pire cas (ex: Werkzeug debugger, Symfony profiler ouvert), un accès à l'exécution de code arbitraire.

## Où ça apparaît typiquement
- Variable d'environnement de framework (`APP_DEBUG=true`, `DEBUG=True`, `ASPNETCORE_ENVIRONMENT=Development`) laissée active en production.
- Pages d'erreur détaillées (stack trace complète, extraits de code source) visibles par les utilisateurs finaux.
- Consoles de debug interactives (Werkzeug/Flask debugger, Rails console web) accessibles sans authentification.
- Fichiers de configuration de déploiement copiés depuis l'environnement de développement sans changement de valeurs.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Configuration explicite (`.env`, `settings.py`, `web.config`, `appsettings.json`) avec un flag debug à `true`/`Development` dans un contexte de déploiement production.
- Réponse d'erreur HTTP contenant une stack trace, des chemins de fichiers serveur ou du code source.
- Endpoint de type console/profiler accessible sans authentification.

## Remédiation
- Désactiver strictement tout mode debug en production (`APP_DEBUG=false`, `DEBUG=False`, `ASPNETCORE_ENVIRONMENT=Production`).
- Afficher des pages d'erreur génériques aux utilisateurs, journaliser les détails uniquement côté serveur (logs internes).
- Automatiser la vérification de la configuration d'environnement dans le pipeline de déploiement (fail-fast si debug actif en prod).
- Voir `rules/remediation/debug-mode.md`.

## Exemple avant/après
Voir `examples/php/debug-mode/`.

## Références
- OWASP Top 10 2021 — A05: Security Misconfiguration
- OWASP Testing Guide: Testing for Error Handling
- CWE-489: Active Debug Code
