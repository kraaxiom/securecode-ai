---
id: jwt-none
category: auth
cwe: CWE-347
owasp: A02:2021-Cryptographic Failures
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# JWT `alg: none`

## Description
La spécification JWT autorise un algorithme de signature `none`, signifiant qu'aucune signature n'est appliquée au token. Si le code de vérification côté serveur accepte cet algorithme ou ne le rejette pas explicitement, un attaquant peut fabriquer un token dont il contrôle intégralement le contenu (y compris les claims de rôle ou d'identité) sans posséder aucun secret, en se contentant de définir `alg: none` dans l'en-tête et de laisser la signature vide.

## Où ça apparaît typiquement
- Vérificateurs JWT maison ou bibliothèques anciennes n'excluant pas explicitement l'algorithme `none`.
- Configurations de bibliothèques JWT laissant la liste d'algorithmes acceptés vide ou par défaut.
- Intégrations tierces réutilisant un decodeur JWT générique sans configuration stricte.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Vérification JWT ne fournissant pas de liste explicite et restreinte d'algorithmes acceptés.
- Utilisation d'une fonction de décodage (`decode`) sans vérification de signature activée par défaut, réutilisée par erreur en contexte d'authentification.
- Absence de test unitaire vérifiant que les tokens `alg: none` sont rejetés.

## Remédiation
- Configurer explicitement la bibliothèque JWT pour n'accepter qu'une liste fermée d'algorithmes forts (ex. `RS256`), en excluant systématiquement `none`.
- Ne jamais utiliser une fonction de simple décodage (sans vérification de signature) pour une décision d'authentification ou d'autorisation.
- Ajouter des tests de non-régression couvrant explicitement le rejet des tokens non signés.
- Voir `rules/remediation/jwt-none.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/python-django/jwt-none/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: JSON Web Token for Java
- CWE-347: Improper Verification of Cryptographic Signature
