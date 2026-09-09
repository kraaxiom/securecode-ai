---
id: missing-authentication
category: websocket
cwe: CWE-306
owasp: A01:2021-Broken Access Control
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# Absence d'authentification sur un canal WebSocket

## Description
Ce défaut survient quand un serveur accepte des connexions WebSocket et traite des messages sensibles (lecture/écriture de données, actions administratives) sans jamais vérifier l'identité de l'appelant. Contrairement aux endpoints REST où un middleware d'authentification est souvent appliqué globalement, les routes WebSocket sont fréquemment ajoutées "à côté" du framework principal et oublient d'hériter des contrôles d'accès existants.

## Où ça apparaît typiquement
- Serveurs WebSocket exposés directement sans passer par le middleware d'authentification de l'application.
- Handshake qui accepte toute connexion (`ws.on('connection')`) puis vérifie l'identité seulement sur certains types de messages, laissant d'autres actions accessibles anonymement.
- Canaux de diffusion (pub/sub) internes exposés publiquement sans contrôle d'accès pensé pour un usage externe.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Callback `on('connection')` qui ne vérifie ni cookie de session, ni token, ni header d'autorisation avant d'accepter la connexion.
- Logique métier sensible exécutée directement dans le handler de message sans passage par une couche d'autorisation.
- Absence de test unitaire/E2E vérifiant qu'une connexion non authentifiée est rejetée.

## Remédiation
- Exiger une authentification explicite (token signé, session validée) avant ou pendant le handshake WebSocket, avant d'accepter tout message.
- Appliquer une vérification d'autorisation par type de message/action, pas seulement à la connexion.
- Fermer immédiatement la connexion (code 1008 - Policy Violation) si l'authentification échoue ou expire.
- Voir `rules/remediation/missing-authentication.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/missing-authentication/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Top 10: A01:2021-Broken Access Control
- CWE-306: Missing Authentication for Critical Function
