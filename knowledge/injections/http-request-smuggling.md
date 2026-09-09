---
id: http-request-smuggling
category: injections
cwe: CWE-444
owasp: A03:2021-Injection
severity_default: high
languages: [php, js, python, java, csharp, go]
---

# HTTP Request Smuggling

## Description
La contrebande de requêtes HTTP (request smuggling) exploite des divergences d'interprétation entre deux serveurs HTTP en chaîne (typiquement un frontal/proxy et un backend) sur la limite entre deux requêtes, en général via une ambiguïté entre les en-têtes `Content-Length` et `Transfer-Encoding`. Un attaquant peut ainsi faire traiter une partie de sa requête comme le début d'une requête distincte, ce qui permet de contourner des contrôles d'accès, empoisonner le cache ou détourner les requêtes d'autres utilisateurs.

## Où ça apparaît typiquement
- Architectures avec un reverse proxy, load balancer ou CDN placé devant un serveur d'application, chacun ayant sa propre implémentation HTTP/1.1.
- Environnements où plusieurs technologies (Nginx, Apache, serveurs applicatifs) coexistent avec des versions ou configurations divergentes du parsing HTTP.
- Endpoints acceptant à la fois `Content-Length` et `Transfer-Encoding: chunked` sans normalisation stricte.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Présence simultanée d'en-têtes `Content-Length` et `Transfer-Encoding` dans le traitement des requêtes entrantes sans rejet explicite du cas ambigu.
- Configuration de proxy ne forçant pas HTTP/2 de bout en bout ou ne normalisant pas les requêtes avant transmission au backend.
- Absence de tests de cohérence entre les composants de la chaîne HTTP sur le parsing des délimiteurs de requêtes.

## Remédiation
- Utiliser HTTP/2 de bout en bout entre le proxy et le backend lorsque c'est possible.
- Configurer le proxy pour rejeter les requêtes ambiguës (présence conjointe de `Content-Length` et `Transfer-Encoding`).
- Maintenir à jour les serveurs HTTP et proxys pour bénéficier des correctifs de parsing.
- Voir `rules/remediation/http-request-smuggling.md` pour les diffs par langage/infrastructure.

## Exemple avant/après
Voir `examples/infra/http-request-smuggling/` (configuration de proxy avant/après, à créer selon le même schéma).

## Références
- OWASP Testing Guide: Testing for HTTP Request Smuggling
- CWE-444: Inconsistent Interpretation of HTTP Requests ('HTTP Request/Response Smuggling')
