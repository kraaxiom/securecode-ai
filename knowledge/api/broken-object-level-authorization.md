---
id: broken-object-level-authorization
category: api
cwe: CWE-639
owasp: API1:2023-Broken Object Level Authorization
severity_default: high
languages: [php, js, python, java, csharp, go]
---

# Broken Object Level Authorization (BOLA)

## Description
Le Broken Object Level Authorization survient lorsqu'une API accepte un identifiant d'objet fourni par le client (dans l'URL, le body ou un paramètre) et effectue une opération sur celui-ci sans vérifier que l'utilisateur authentifié est réellement autorisé à accéder à cet objet précis. C'est la vulnérabilité la plus fréquente et la plus critique du top 10 API OWASP, car elle permet à un utilisateur légitime d'accéder aux données d'un autre utilisateur simplement en changeant un identifiant. L'authentification seule ne suffit pas : c'est l'absence de contrôle d'autorisation par ressource qui est en cause.

## Où ça apparaît typiquement
- Endpoints du type `GET /api/orders/{id}`, `GET /api/users/{id}/invoices` où `{id}` est utilisé tel quel pour la requête base de données.
- Opérations de modification/suppression (`PUT`, `DELETE`) ciblant une ressource par identifiant sans vérification de propriété.
- APIs exposant des identifiants séquentiels ou prévisibles (IDs numériques auto-incrémentés) facilitant l'énumération.
- Endpoints internes réutilisés tels quels côté mobile/API publique sans réintroduire les contrôles d'autorisation applicatifs.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Requête de récupération/modification d'objet par ID sans clause de filtrage sur l'utilisateur courant (`WHERE id = :id` sans `AND user_id = :current_user`).
- Contrôleur qui récupère l'objet directement depuis le paramètre de route sans passer par une vérification d'appartenance ou une couche de policy/autorisation.
- Absence de middleware ou décorateur d'autorisation au niveau ressource (vérification uniquement au niveau authentification, pas au niveau objet).
- Fonctions génériques de type "get by id" partagées entre plusieurs contextes utilisateurs sans paramètre de scope.

## Remédiation
- Vérifier systématiquement, à chaque accès à un objet, que l'utilisateur authentifié est propriétaire ou explicitement autorisé sur cette ressource précise (pas seulement authentifié).
- Centraliser la logique d'autorisation dans une couche dédiée (policies, guards) plutôt que de la dupliquer dans chaque contrôleur.
- Préférer des identifiants non devinables (UUID) en complément des contrôles d'autorisation, sans s'y substituer.
- Ajouter des tests d'autorisation automatisés couvrant l'accès croisé entre comptes utilisateurs.
- Voir `rules/remediation/broken-object-level-authorization.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/broken-object-level-authorization/` (et les répertoires équivalents pour js, python, java, csharp, go).

## Références
- OWASP API Security Top 10 2023: API1-Broken Object Level Authorization
- CWE-639: Authorization Bypass Through User-Controlled Key
