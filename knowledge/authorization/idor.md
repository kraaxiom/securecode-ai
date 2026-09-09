---
id: idor
category: authorization
cwe: CWE-639
owasp: A01:2021-Broken Access Control
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Insecure Direct Object Reference (IDOR)

## Description
Une IDOR survient quand une application expose une référence directe à un objet interne (ID de base de données, nom de fichier, clé) dans une requête, et se contente de vérifier que l'utilisateur est authentifié sans vérifier qu'il est bien autorisé à accéder à cet objet précis. En modifiant simplement l'identifiant dans la requête, un attaquant peut accéder aux données ou ressources d'un autre utilisateur.

## Où ça apparaît typiquement
- Endpoints de type `GET /invoices/{id}` ou `GET /users/{id}/profile` où `{id}` est un identifiant séquentiel ou devinable.
- Téléchargement de fichiers ou documents référencés par un identifiant dans l'URL.
- API REST/GraphQL retournant un objet par identifiant sans vérifier son appartenance à l'utilisateur courant.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Requête de récupération d'objet par identifiant fourni par le client, sans clause de filtrage sur le propriétaire/tenant dans la requête de base de données.
- Vérification d'autorisation limitée à "l'utilisateur est connecté" sans vérification "l'utilisateur possède/peut accéder à cet objet précis".
- Identifiants d'objets séquentiels ou prévisibles combinés à l'absence de contrôle d'appartenance.

## Remédiation
- Toujours inclure une clause de vérification d'appartenance (ex. `WHERE user_id = :current_user`) dans toute requête de récupération d'objet par identifiant.
- Centraliser la vérification d'autorisation par objet dans une couche dédiée (policy/guard) appliquée systématiquement.
- Envisager des identifiants non séquentiels (UUID) pour réduire la surface de découverte, en complément (pas en remplacement) du contrôle d'autorisation.
- Voir `rules/remediation/idor.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/idor/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Authorization
- CWE-639: Authorization Bypass Through User-Controlled Key
