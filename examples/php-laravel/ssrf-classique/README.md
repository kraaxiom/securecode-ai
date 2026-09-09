## Vulnérabilité
Server-Side Request Forgery (SSRF) classique — CWE-918.

## Impact
Le contrôleur `UrlPreviewController::preview` effectue une requête HTTP sortante vers une URL entièrement contrôlée par l'utilisateur, sans aucune restriction. Un attaquant peut détourner le serveur pour interroger des ressources internes normalement inaccessibles depuis l'extérieur : interfaces d'administration, services internes, bases de données exposées sur le réseau privé, ou API cloud de métadonnées. Il peut aussi exploiter les redirections HTTP suivies automatiquement pour contourner un filtrage basé uniquement sur l'URL initiale.

## Cause racine
Absence totale de validation de la destination réseau avant l'émission de la requête sortante : ni whitelist de domaines, ni vérification de schéma, ni résolution/validation de l'adresse IP réelle, ni contrôle des redirections.

## Correction
- Restriction des destinations à une whitelist explicite de domaines métier (`ALLOWED_HOSTS`).
- Schéma limité à `https` uniquement.
- Résolution DNS explicite du domaine puis validation de l'IP obtenue contre les plages privées/loopback/link-local (`FILTER_FLAG_NO_PRIV_RANGE | FILTER_FLAG_NO_RES_RANGE`).
- Désactivation des redirections HTTP automatiques (`allow_redirects: false`).
- Timeout court sur la requête sortante pour limiter l'impact d'un scan interne.

## Références
- CWE-918: Server-Side Request Forgery (SSRF)
- OWASP: A10:2021-Server-Side Request Forgery (SSRF)
- OWASP Cheat Sheet: Server Side Request Forgery Prevention
