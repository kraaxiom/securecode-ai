---
id: message-injection
category: websocket
cwe: CWE-20
owasp: A03:2021-Injection
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Injection dans les messages WebSocket

## Description
L'injection dans les messages WebSocket survient lorsque le contenu reçu sur un canal temps réel est traité comme du code, une commande ou une requête sans validation ni encodage, exactement comme une entrée HTTP classique. Comme les messages WebSocket échappent souvent aux couches de validation habituelles (WAF, middleware de sanitation HTTP), ils constituent un vecteur privilégié pour propager des injections SQL, NoSQL, de commandes ou du XSS stocké/diffusé en direct à d'autres clients.

## Où ça apparaît typiquement
- Serveurs de chat ou de notifications qui rediffusent (`broadcast`) le contenu d'un message à d'autres clients sans encodage.
- Handlers de messages WebSocket qui utilisent le payload directement dans une requête base de données.
- Applications qui désérialisent le message (JSON, MessagePack) et l'utilisent pour construire des commandes système ou des requêtes dynamiques.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fonction `on('message', ...)` qui retransmet directement le payload reçu à d'autres sockets sans échappement HTML/JS.
- Contenu du message WebSocket concaténé dans une requête SQL/NoSQL ou une commande shell.
- Absence de schéma de validation (type, longueur, format) appliqué aux messages entrants.
- Désérialisation du message sans validation stricte du type attendu.

## Remédiation
- Traiter chaque message WebSocket comme une entrée non fiable : valider son schéma, son type et sa taille avant traitement.
- Encoder/échapper tout contenu utilisateur avant de le rediffuser à d'autres clients (contexte HTML/JS).
- Utiliser des requêtes paramétrées pour toute interaction avec une base de données déclenchée par un message.
- Voir `rules/remediation/message-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/message-injection/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Testing Guide: Testing WebSockets
- CWE-20: Improper Input Validation
