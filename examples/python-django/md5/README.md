# Hachage MD5 des mots de passe — Python/Django

`vulnerable.py` hache les mots de passe utilisateurs avec `hashlib.md5()` avant stockage, à l'inscription comme à la connexion (CWE-328, Use of Weak Hash). MD5 est cassé depuis 2004, ne comporte ni sel ni facteur de coût, et est cassable par force brute GPU/table arc-en-ciel en un temps très court.

`fixed.py` remplace MD5 par Argon2id via la bibliothèque `argon2-cffi`, avec sel automatique, facteur de coût configurable, vérification sécurisée et re-hachage progressif lorsque les paramètres doivent être renforcés.

## Pourquoi c'est dangereux
- MD5 permet de générer des collisions en quelques secondes et ne résiste pas à une attaque par force brute moderne (GPU/ASIC).
- L'absence de sel signifie que deux utilisateurs avec le même mot de passe produisent le même hash, facilitant les attaques par table arc-en-ciel précalculée.
- La comparaison de hachages sous forme de chaînes (`==`) est en plus vulnérable aux attaques par mesure de temps (timing attack).

## Explication du correctif
- Remplacement de `hashlib.md5()` par `PasswordHasher().hash()` (Argon2id), qui génère un sel aléatoire et embarque les paramètres de coût dans le hash produit.
- Vérification via `password_hasher.verify()`, qui effectue une comparaison constante en temps et lève une exception explicite en cas d'échec.
- Ajout de `check_needs_rehash()` pour renouveler automatiquement les hachages dont le paramétrage est devenu obsolète, à la prochaine connexion réussie.

## Notes résiduelles
- Les mots de passe déjà stockés en MD5 doivent être migrés progressivement : au prochain login réussi avec l'ancien hash, recalculer et stocker un hash Argon2id.
- MD5 peut rester acceptable uniquement pour des checksums non cryptographiques (détection d'erreur de transfert), jamais pour des mots de passe ou des signatures.
