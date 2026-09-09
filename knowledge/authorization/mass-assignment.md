---
id: mass-assignment
category: authorization
cwe: CWE-915
owasp: A08:2021-Software and Data Integrity Failures
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Mass Assignment

## Description
Le mass assignment survient quand une application lie automatiquement les champs d'une requête entrante (JSON, formulaire) directement aux attributs d'un modèle ou d'une entité, sans restreindre explicitement quels champs sont autorisés à être modifiés par le client. Un attaquant peut alors injecter des champs supplémentaires non prévus dans le formulaire d'origine (ex. `role`, `isAdmin`, `balance`) pour modifier des attributs sensibles auxquels il ne devrait pas avoir accès.

## Où ça apparaît typiquement
- Frameworks avec liaison automatique objet-requête (binding ORM) utilisée sans liste blanche de champs autorisés.
- Endpoints de mise à jour de profil utilisateur acceptant l'intégralité du corps JSON pour peupler le modèle.
- API de création de ressources réutilisant le même modèle/DTO pour l'entrée utilisateur et pour la représentation interne complète.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Liaison directe du corps de requête à un modèle/entité complet (`Model.create(req.body)` ou équivalent) sans filtrage de champs.
- Absence de DTO ou de schéma de validation dédié à l'entrée, distinct du modèle de données interne complet.
- Attributs sensibles (rôle, statut, solde, drapeaux internes) portés par le même modèle que les champs modifiables par l'utilisateur, sans séparation explicite.

## Remédiation
- Utiliser une liste blanche explicite des champs autorisés en entrée pour chaque opération de création/mise à jour, plutôt qu'une liaison automatique complète.
- Séparer le modèle de données interne des objets d'entrée (DTO) exposés à l'utilisateur.
- Marquer explicitement les attributs sensibles comme non assignables en masse au niveau du framework/ORM.
- Voir `rules/remediation/mass-assignment.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/mass-assignment/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Mass Assignment
- CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes
