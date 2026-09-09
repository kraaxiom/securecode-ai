---
id: request-desynchronization
category: modern
cwe: CWE-444
owasp: A04:2021-Insecure Design
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# Request Smuggling / Desynchronization HTTP

## Description
La désynchronisation de requêtes HTTP (HTTP Request Smuggling) exploite une divergence d'interprétation des limites de requêtes entre deux composants d'une chaîne HTTP (reverse proxy et serveur d'origine), typiquement liée aux en-têtes `Content-Length` et `Transfer-Encoding` traités différemment par chacun. Un attaquant peut construire une requête ambiguë qui est comprise comme une seule requête par le frontal mais comme deux par le serveur d'origine (ou l'inverse), permettant d'injecter une requête cachée qui sera traitée dans le contexte d'un autre utilisateur, contournant des contrôles de sécurité au niveau du frontal.

## Où ça apparaît typiquement
- Architectures avec reverse proxy/load balancer devant un serveur applicatif, chacun avec sa propre implémentation HTTP.
- Requêtes contenant simultanément les en-têtes `Content-Length` et `Transfer-Encoding: chunked`.
- Composants intermédiaires (CDN, WAF) et serveurs d'origine de fournisseurs/versions différents avec des tolérances de parsing HTTP différentes.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Chaîne d'infrastructure combinant plusieurs implémentations HTTP distinctes (proxy A + serveur B) sans normalisation stricte des en-têtes.
- Configuration serveur qui ne rejette pas explicitement les requêtes contenant à la fois `Content-Length` et `Transfer-Encoding`.
- Absence de bascule complète vers HTTP/2 de bout en bout (le request smuggling classique cible surtout HTTP/1.1).

## Remédiation
- Configurer chaque composant de la chaîne pour rejeter strictement les requêtes ambiguës (présence simultanée de `Content-Length` et `Transfer-Encoding`).
- Normaliser les requêtes au niveau du frontal avant transmission au backend, en utilisant la même implémentation/version HTTP autant que possible.
- Privilégier HTTP/2 de bout en bout entre les composants internes lorsque c'est possible, pour éliminer la classe de vulnérabilité liée au chunked encoding.
- Voir `rules/remediation/request-desynchronization.md` pour les diffs de configuration.

## Exemple avant/après
Voir `examples/config/request-desynchronization/` (configuration proxy avant/après durcissement du parsing HTTP).

## Références
- OWASP: HTTP Request Smuggling
- CWE-444: Inconsistent Interpretation of HTTP Requests ('HTTP Request/Response Smuggling')
