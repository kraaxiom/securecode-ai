---
id: cswsh
category: websocket
cwe: CWE-352
owasp: A01:2021-Broken Access Control
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Cross-Site WebSocket Hijacking (CSWSH)

## Description
Le Cross-Site WebSocket Hijacking est l'équivalent du CSRF pour les connexions WebSocket : un serveur accepte l'établissement d'une connexion `ws://`/`wss://` en se basant uniquement sur les cookies de session du navigateur, sans vérifier l'origine de la requête ni exiger de jeton anti-CSRF. Une page malveillante hébergée sur un autre domaine peut alors ouvrir une connexion WebSocket authentifiée au nom de la victime et échanger des données avec le serveur. Contrairement aux requêtes HTTP classiques, la politique de même origine (SOP) ne s'applique pas nativement aux handshakes WebSocket.

## Où ça apparaît typiquement
- Endpoints WebSocket (`/ws`, `/socket`, `/realtime`) authentifiés uniquement via cookie de session.
- Handshake WebSocket qui ne vérifie pas l'en-tête `Origin` côté serveur.
- Frameworks temps réel (Socket.IO, SignalR, ActionCable, ws) utilisés avec la configuration par défaut sans contrôle d'origine explicite.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Code de handshake WebSocket qui accepte la connexion sans lire ni valider `request.headers['Origin']`.
- Authentification basée uniquement sur un cookie de session, sans jeton distinct transmis dans le message initial.
- Absence de liste blanche d'origines dans la configuration du serveur WebSocket.
- Middleware CORS configuré pour l'HTTP classique mais absent sur la route d'upgrade WebSocket.

## Remédiation
- Valider strictement l'en-tête `Origin` lors du handshake contre une liste blanche de domaines autorisés.
- Ne pas s'appuyer uniquement sur les cookies : exiger un jeton d'authentification propre (ex: token signé transmis dans le premier message ou en query param signé à courte durée de vie) en complément.
- Utiliser des cookies `SameSite=Strict` ou `Lax` pour limiter l'envoi cross-site du cookie de session lors du handshake.
- Voir `rules/remediation/cswsh.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/cswsh/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP: Cross-Site WebSocket Hijacking
- CWE-352: Cross-Site Request Forgery (CSRF)
