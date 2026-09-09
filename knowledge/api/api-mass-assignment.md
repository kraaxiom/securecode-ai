---
id: api-mass-assignment
category: api
cwe: CWE-915
owasp: API6:2019-Mass Assignment
severity_default: high
languages: [php, js, python, java, csharp, go]
---

# Mass Assignment sur API

## Description
Le mass assignment survient lorsqu'une API lie automatiquement les champs d'une requête entrante (JSON, formulaire) directement aux attributs d'un objet métier ou d'un modèle de base de données, sans liste blanche explicite des champs autorisés. Un attaquant peut alors injecter des champs supplémentaires non prévus par le formulaire d'origine, comme un rôle, un statut de vérification ou un solde, et les faire persister. Ce problème est fréquent avec les ORM et frameworks qui offrent un binding automatique "pratique" par défaut.

## Où ça apparaît typiquement
- Endpoints de création/mise à jour (`POST`, `PUT`, `PATCH`) qui passent le body JSON directement à un constructeur de modèle ou une méthode `create()`/`update()`/`save()`.
- Utilisation de fonctions de "binding" automatique d'un framework (`Model::create($request->all())`, désérialisation directe JSON vers entité sans DTO).
- Formulaires d'inscription ou de profil utilisateur permettant de modifier des champs sensibles (`role`, `isAdmin`, `balance`, `verified`) non censés être exposés au client.
- APIs REST/GraphQL qui exposent le même modèle interne en entrée et en sortie sans distinction.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel direct du corps de requête complet à une opération de persistance (`Model::create($request->all())`, `Object.assign(entity, req.body)`, désérialisation JSON directe sur une entité JPA/Hibernate).
- Absence de DTO/schéma d'entrée distinct du modèle de persistance.
- Absence de liste blanche (`$fillable`, `@JsonIgnoreProperties`, allowlist explicite de champs) ou usage d'une liste noire incomplète.
- Champs sensibles (`role`, `is_admin`, `status`, `balance`) présents dans le même modèle que les champs utilisateur ordinaires sans annotation de protection.

## Remédiation
- Utiliser des DTOs ou schémas de validation d'entrée distincts des modèles de persistance, avec une liste blanche explicite des champs autorisés par endpoint.
- Configurer les mécanismes de protection natifs des frameworks (`$fillable`/`$guarded` en Laravel, `@JsonIgnore` en Java, exclusion explicite de champs en Django REST serializers).
- Ne jamais exposer les champs privilégiés (rôle, statut, identifiants internes) dans les schémas d'entrée acceptés par le client.
- Valider côté serveur toute tentative de modification de champs sensibles indépendamment de la présence dans le payload.
- Voir `rules/remediation/api-mass-assignment.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/api-mass-assignment/` (et les répertoires équivalents pour js, python, java, csharp, go).

## Références
- OWASP API Security Top 10 2019: API6-Mass Assignment
- CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes
