---
id: blind-ssrf
category: ssrf
cwe: CWE-918
owasp: A10:2021-Server-Side Request Forgery (SSRF)
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Blind SSRF

## Description
La SSRF aveugle (blind SSRF) est une variante où l'application effectue bien une requête sortante contrôlée par l'attaquant, mais ne retourne aucune information sur le résultat de cette requête (statut, contenu, erreur) dans la réponse observable. L'exploitation repose alors sur des canaux indirects (délai de réponse, callbacks DNS/HTTP hors bande) plutôt que sur la lecture directe du contenu récupéré, ce qui la rend plus difficile à détecter mais tout aussi dangereuse pour le scan de réseau interne ou l'exfiltration de métadonnées.

## Où ça apparaît typiquement
- Traitement asynchrone de webhooks ou de notifications sortantes dont le résultat n'est jamais renvoyé au client.
- Systèmes de validation d'URL "fire and forget" (vérification d'un lien avant enregistrement, sans afficher le contenu récupéré).
- Parsers de documents (XML, PDF) déclenchant des requêtes externes sans exposer d'erreur détaillée au client.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Requête sortante déclenchée par une entrée utilisateur dans un contexte où la réponse n'est jamais renvoyée à l'appelant (job asynchrone, worker, callback).
- Absence de whitelist de destinations et de filtrage des plages d'adresses internes, identique aux cas de SSRF classique mais sans retour visible.
- Traitement de formats de fichiers connus pour déclencher des résolutions d'entités externes (XML External Entity) sans désactivation de cette fonctionnalité.

## Remédiation
- Appliquer les mêmes contrôles que pour une SSRF classique (whitelist, résolution/validation d'IP, isolation réseau) même en l'absence de retour visible.
- Désactiver la résolution d'entités externes dans les parseurs XML/documents.
- Surveiller et journaliser les requêtes sortantes du serveur pour détecter des tentatives de scan interne.
- Voir `rules/remediation/blind-ssrf.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/python-django/blind-ssrf/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Server Side Request Forgery Prevention
- CWE-918: Server-Side Request Forgery (SSRF)
