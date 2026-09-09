# Utilisation de SHA-1 (fonction de hachage affaiblie) — Python/Django

`vulnerable.py` implémente un stockage de mot de passe fait maison basé sur `hashlib.sha1()` (non salé, sans facteur de coût) ainsi qu'une fonction de signature d'export de données sensibles utilisant une concaténation `secret + contenu` hachée en SHA-1 (CWE-328, Use of Weak Hash). SHA-1 est cryptographiquement cassé depuis l'attaque "SHAttered" (2017) et reste, de plus, beaucoup trop rapide pour un usage de hachage de mot de passe.

`fixed.py` remplace le hachage fait maison par le système d'authentification natif de Django (`set_password()` / `authenticate()`) configuré avec `Argon2PasswordHasher`, et remplace la signature d'export par un HMAC-SHA256 correctement construit via `hmac.new()`.

## Pourquoi c'est dangereux
- SHA-1 est vulnérable aux attaques par collision, ce qui compromet toute preuve d'intégrité ou de signature qui en dépend.
- Utilisé pour des mots de passe, SHA-1 non salé et non ralenti est trivialement cassable par rainbow table ou par force brute GPU en cas de fuite de la base.
- Une concaténation `secret + message` hachée n'est pas un HMAC : elle est vulnérable aux attaques d'extension de longueur (length extension), propres aux fonctions de hachage de type Merkle-Damgård comme SHA-1.

## Explication du correctif
- Délégation du hachage de mot de passe au framework Django (`PASSWORD_HASHERS`), avec Argon2id en tête de liste : salage automatique, facteur de coût réglable, résistance au brute force.
- Utilisation d'`authenticate()` / `check_password()` pour une comparaison à temps constant, évitant les attaques par timing.
- Remplacement du hachage SHA-1 par un HMAC-SHA256 construit avec l'API `hmac.new()`, immunisée contre les attaques d'extension de longueur et les collisions.

## Notes résiduelles
- Les mots de passe existants hachés en SHA-1 doivent être migrés progressivement (re-hachage à la prochaine connexion réussie de l'utilisateur).
- Vérifier qu'aucun certificat ou artefact signé en SHA-1 ne subsiste dans la chaîne de confiance de l'infrastructure.
