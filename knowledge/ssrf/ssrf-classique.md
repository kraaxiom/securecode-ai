---
id: ssrf-classique
category: ssrf
cwe: CWE-918
owasp: A10:2021-Server-Side Request Forgery (SSRF)
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Server-Side Request Forgery (SSRF) classique

## Description
Une SSRF survient quand une application effectue une requête réseau sortante vers une URL ou un hôte dont tout ou partie est contrôlé par l'utilisateur, sans validation stricte de la destination. L'attaquant détourne le serveur pour qu'il émette des requêtes vers des ressources internes normalement inaccessibles depuis l'extérieur : services internes, interfaces d'administration, ou API cloud de métadonnées.

## Où ça apparaît typiquement
- Fonctionnalités de récupération d'aperçu (webhook, preview d'URL, fetch d'avatar/image distante).
- Intégrations qui appellent une API tierce dont l'URL de base est configurable par l'utilisateur.
- Traitement de fichiers XML/PDF/SVG pouvant référencer des ressources externes.
- Proxys applicatifs génériques (`?url=`).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à une bibliothèque HTTP (client HTTP, `fetch`, `curl`, `requests`) avec une URL/hôte provenant directement d'une entrée utilisateur.
- Absence de whitelist de domaines/hôtes autorisés pour les requêtes sortantes.
- Absence de filtrage des plages d'adresses privées/loopback/link-local avant l'émission de la requête.

## Remédiation
- Restreindre les destinations à une whitelist explicite de domaines/hôtes métier.
- Résoudre le DNS et valider l'adresse IP résultante contre les plages privées/réservées avant chaque requête sortante (voir `dns-rebinding.md` pour le contournement par résolution différée).
- Désactiver les redirections automatiques ou revalider chaque redirection.
- Isoler le service effectuant des requêtes sortantes dans un segment réseau sans accès aux ressources internes sensibles.
- Voir `rules/remediation/ssrf-classique.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/ssrf-classique/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Server Side Request Forgery Prevention
- CWE-918: Server-Side Request Forgery (SSRF)
