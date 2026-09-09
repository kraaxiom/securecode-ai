---
id: redis-ssrf
category: ssrf
cwe: CWE-918
owasp: A10:2021-Server-Side Request Forgery (SSRF)
severity_default: critical
languages: [php, python, js, java, go, csharp, rust]
---

# SSRF ciblant Redis

## Description
Redis expose par défaut un protocole texte simple (RESP) sans authentification ni chiffrement, conçu pour un usage strictement interne. Une SSRF permettant à un attaquant de faire émettre des requêtes HTTP (ou d'ouvrir des connexions bas niveau) vers une instance Redis interne peut, selon le contexte, permettre de lire/écrire des clés, voire d'obtenir une exécution de code lorsque Redis est combiné à des mécanismes de persistance ou de modules mal sécurisés. Cette faille est aggravée quand le protocole HTTP tolère des injections de commandes brutes vers le port Redis.

## Où ça apparaît typiquement
- Applications permettant de configurer une URL/hôte de cache ou de file de tâches côté utilisateur.
- Fonctions de "webhook"/proxy générique qui peuvent être détournées pour atteindre le port Redis interne.
- Environnements cloud où Redis est accessible sans authentification à toute machine du même réseau virtuel.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Construction d'une destination réseau (hôte/port) à partir d'une entrée utilisateur, sans whitelist, dans un contexte où le réseau interne héberge une instance Redis.
- Absence d'authentification (`requirepass`/ACL) sur les instances Redis accessibles depuis le service applicatif.
- Absence de segmentation réseau isolant Redis des composants exposés à des entrées utilisateur non fiables.

## Remédiation
- Activer l'authentification et les ACL Redis, même en réseau interne, et lier Redis à une interface non exposée publiquement.
- Restreindre strictement les destinations autorisées pour toute requête sortante contrôlable par l'utilisateur (whitelist d'hôtes).
- Segmenter le réseau pour qu'aucun service traitant des entrées utilisateur non fiables ne puisse atteindre directement le port Redis.
- Voir `rules/remediation/redis-ssrf.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/python-django/redis-ssrf/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Server Side Request Forgery Prevention
- CWE-918: Server-Side Request Forgery (SSRF)
