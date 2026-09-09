---
id: elasticsearch-ssrf
category: ssrf
cwe: CWE-918
owasp: A10:2021-Server-Side Request Forgery (SSRF)
severity_default: high
languages: [java, python, js, php, go]
---

# SSRF ciblant Elasticsearch

## Description
Cette variante de SSRF cible spécifiquement les instances Elasticsearch (ou clusters de recherche similaires) exposées uniquement sur le réseau interne. Un attaquant qui parvient à faire émettre une requête HTTP arbitraire par le serveur applicatif peut interroger directement l'API REST d'Elasticsearch (souvent non authentifiée par défaut) pour lister les index, extraire des données sensibles, ou dans certaines configurations legacy déclencher l'exécution de scripts côté serveur.

## Où ça apparaît typiquement
- Applications exposant un paramètre permettant de configurer l'hôte/URL du backend de recherche.
- Fonctionnalités de proxy générique vers un cluster interne (dashboards, outils d'administration).
- Microservices qui relaient des requêtes utilisateur vers Elasticsearch sans validation de destination.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Construction d'une URL vers un cluster Elasticsearch à partir d'une entrée utilisateur (hôte, port ou chemin d'API).
- Absence d'authentification/autorisation sur l'accès réseau au cluster (accessible depuis le service applicatif sans contrôle supplémentaire).
- Absence de whitelist stricte limitant les requêtes sortantes du service applicatif au seul hôte Elasticsearch attendu.

## Remédiation
- Ne jamais permettre à un utilisateur de contrôler l'hôte/port de destination des requêtes vers le cluster de recherche.
- Activer l'authentification et le contrôle d'accès sur l'API Elasticsearch, même en réseau interne.
- Restreindre l'accès réseau au cluster au seul service applicatif légitime (segmentation réseau, règles de pare-feu).
- Voir `rules/remediation/elasticsearch-ssrf.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/java-spring/elasticsearch-ssrf/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Server Side Request Forgery Prevention
- CWE-918: Server-Side Request Forgery (SSRF)
