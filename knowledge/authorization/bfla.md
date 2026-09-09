---
id: bfla
category: authorization
cwe: CWE-862
owasp: A01:2021-Broken Access Control
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Broken Function Level Authorization (BFLA)

## Description
La BFLA survient quand une application ne vérifie pas correctement qu'un utilisateur dispose du rôle ou du niveau de privilège requis pour accéder à une fonctionnalité donnée, indépendamment de la ressource ciblée. Contrairement à la BOLA (contrôle au niveau de l'objet), la BFLA concerne le contrôle au niveau de la fonction/action elle-même : un utilisateur standard peut par exemple appeler directement un endpoint d'administration s'il en connaît l'existence, faute de vérification de rôle.

## Où ça apparaît typiquement
- Endpoints d'administration (`/admin/*`, `DELETE /users/{id}`) protégés uniquement par l'authentification, sans vérification de rôle.
- Fonctions sensibles exposées dans l'API mais non présentes dans l'interface utilisateur standard (sécurité par obscurité).
- Différences de contrôle d'autorisation entre l'interface web (qui masque le bouton) et l'API sous-jacente (qui exécute l'action sans revérifier le rôle).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Endpoint effectuant une action privilégiée sans middleware/décorateur de vérification de rôle explicite.
- Contrôle de rôle effectué uniquement côté interface utilisateur (masquage de bouton/menu) sans contrôle équivalent côté serveur.
- Incohérence entre les rôles vérifiés selon les différentes routes exposant une même fonctionnalité.

## Remédiation
- Appliquer une vérification de rôle/permission explicite et centralisée sur chaque endpoint exposant une fonctionnalité sensible, y compris ceux non visibles dans l'interface.
- Ne jamais s'appuyer sur l'absence d'affichage d'une fonctionnalité côté client comme mesure de sécurité.
- Cartographier les fonctionnalités par niveau de privilège requis et auditer systématiquement leur couverture par des tests d'autorisation.
- Voir `rules/remediation/bfla.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/python-django/bfla/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP API Security Top 10: API5:2023 Broken Function Level Authorization
- CWE-862: Missing Authorization
