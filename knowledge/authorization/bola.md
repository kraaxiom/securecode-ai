---
id: bola
category: authorization
cwe: CWE-639
owasp: A01:2021-Broken Access Control
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Broken Object Level Authorization (BOLA)

## Description
La BOLA est l'équivalent de l'IDOR dans le contexte des API (RESTful ou GraphQL) : un endpoint accepte un identifiant d'objet fourni par le client et effectue une opération (lecture, mise à jour, suppression) sans vérifier que l'appelant est autorisé à agir sur cet objet précis. C'est la vulnérabilité la plus fréquente listée dans l'OWASP API Security Top 10.

## Où ça apparaît typiquement
- Endpoints d'API REST manipulant une ressource identifiée par ID dans le chemin ou le corps de la requête (`PUT /api/orders/{id}`).
- Résolveurs GraphQL récupérant un objet par ID sans vérification d'appartenance dans le resolver.
- API multi-tenant où l'isolation entre tenants repose uniquement sur un filtre côté client.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Handler d'API récupérant/modifiant un objet par ID transmis par le client sans clause de filtrage sur le propriétaire ou le tenant.
- Resolver GraphQL exposant une requête paramétrée par ID sans vérification d'autorisation dans le resolver lui-même.
- Tests d'API limités aux cas nominaux, sans scénario d'accès croisé entre deux comptes différents.

## Remédiation
- Vérifier systématiquement, au niveau de chaque endpoint/resolver manipulant un objet par ID, que l'appelant est autorisé sur cet objet précis (pas seulement authentifié).
- Appliquer un filtrage d'appartenance au niveau de la requête de données elle-même, pas en post-traitement.
- Ajouter des tests d'API dédiés à l'accès croisé entre comptes/tenants distincts.
- Voir `rules/remediation/bola.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/bola/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP API Security Top 10: API1:2023 Broken Object Level Authorization
- CWE-639: Authorization Bypass Through User-Controlled Key
