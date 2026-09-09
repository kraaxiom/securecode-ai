# Secret JWT faible — Python/Django

`vulnerable.py` signe des tokens JWT (HS256) avec un secret court et codé en dur (`"changeme123"`) directement dans le code source, et la vérification accepte une liste d'algorithmes incluant `none` (CWE-1391, Use of Weak Credentials). Un attaquant disposant du code source ou capable de bruteforcer le secret hors ligne peut forger des tokens valides pour n'importe quel utilisateur.

`fixed.py` charge le secret depuis une variable d'environnement (générée hors du code avec un CSPRNG, au moins 256 bits), vérifie sa longueur au démarrage, restreint strictement l'algorithme accepté à `HS256` lors du décodage, et ajoute une expiration courte (`exp`) sur chaque token émis.

## Pourquoi c'est dangereux
- Un secret HS256 court ou prévisible peut être retrouvé par force brute hors ligne, permettant de forger des tokens pour n'importe quel compte.
- Accepter `alg: none` ou une liste d'algorithmes non contrôlée permet à un attaquant de fournir un token non signé et de le faire accepter comme valide.
- Un secret codé en dur dans le code source versionné est exposé à quiconque a accès au dépôt (y compris d'anciens contributeurs, forks publics, etc.).

## Explication du correctif
- Le secret est chargé via `os.environ["JWT_SECRET"]` et généré hors du code applicatif avec une entropie suffisante (≥ 256 bits).
- Une vérification de longueur minimale échoue au démarrage si le secret est trop faible.
- `jwt.decode()` reçoit un paramètre `algorithms` restreint à une liste fermée (`["HS256"]`), empêchant toute confusion d'algorithme ou acceptation de `none`.
- Ajout d'une expiration courte (`exp`) pour limiter la durée de vie exploitable d'un token compromis.

## Notes résiduelles
- Prévoir une procédure de rotation du secret (avec courte période de recouvrement) en cas de compromission suspectée.
- Pour une architecture multi-services, privilégier une signature asymétrique (RS256/ES256) afin que seuls les émetteurs détiennent la clé privée.
