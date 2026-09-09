# Remédiation — JWT Algorithm Confusion

## Principe
Toujours imposer côté serveur l'algorithme de vérification attendu, sans jamais le déduire du champ `alg` du token lui-même. Ne jamais réutiliser une clé publique RSA comme secret HMAC potentiel.

## PHP (firebase/php-jwt)
```php
// Avant — vulnérable : l'algorithme n'est pas restreint explicitement
use Firebase\JWT\JWT;

$decoded = JWT::decode($token, new Key($publicKey, null)); // algo déduit du header

// Après — sécurisé : algorithme figé et explicite
use Firebase\JWT\JWT;
use Firebase\JWT\Key;

$decoded = JWT::decode($token, new Key($publicKey, 'RS256')); // uniquement RS256 accepté
```

## JS / Node.js (jsonwebtoken)
```js
// Avant — vulnérable : pas de restriction d'algorithme, header fait foi
const payload = jwt.verify(token, publicKey);

// Après — sécurisé : liste fermée d'algorithmes, clé dédiée à la vérification
const payload = jwt.verify(token, publicKey, {
  algorithms: ['RS256'], // exclut explicitement HS256 et toute autre famille
});
```

## Python (PyJWT)
```python
# Avant — vulnérable : algorithme non restreint
payload = jwt.decode(token, public_key)

# Après — sécurisé : algorithme explicitement imposé
payload = jwt.decode(token, public_key, algorithms=["RS256"])
```

## Checklist de vérification post-patch
- [ ] Chaque appel de vérification JWT spécifie explicitement l'algorithme (ou une liste fermée), sans jamais le déduire du token.
- [ ] Aucune clé publique RSA n'est réutilisable comme secret HMAC dans le code de vérification.
- [ ] Un test confirme qu'un token signé en HS256 avec la clé publique RSA comme secret est rejeté.
- [ ] La bibliothèque JWT utilisée est à jour et maintenue.
- [ ] Les clés de signature et de vérification sont distinctes par algorithme et stockées séparément.
