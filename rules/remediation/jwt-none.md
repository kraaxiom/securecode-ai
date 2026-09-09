# Remédiation — JWT `alg: none`

## Principe
Exclure explicitement l'algorithme `none` de toute vérification JWT et ne jamais utiliser une fonction de simple décodage (sans vérification de signature) pour une décision d'authentification ou d'autorisation.

## PHP (firebase/php-jwt)
```php
// Avant — vulnérable : décodage sans vérifier la liste d'algorithmes acceptés
$decoded = JWT::decode($token, new Key($secret, '')); // algo vide, potentiellement 'none' accepté

// Après — sécurisé : algorithme fort explicitement imposé
$decoded = JWT::decode($token, new Key($secret, 'HS256')); // 'none' impossible à passer
```

## JS / Node.js (jsonwebtoken)
```js
// Avant — vulnérable : jwt.decode() ne vérifie aucune signature
const payload = jwt.decode(token); // AUCUNE vérification cryptographique !
if (payload.role === 'admin') { /* ... */ }

// Après — sécurisé : jwt.verify() avec algorithmes explicitement whitelistés
const payload = jwt.verify(token, secretOrPublicKey, {
  algorithms: ['HS256'], // 'none' est structurellement exclu de cette liste
});
if (payload.role === 'admin') { /* ... */ }
```

## Python (PyJWT)
```python
# Avant — vulnérable : algorithms non fourni, PyJWT peut accepter 'none' selon la version/config
payload = jwt.decode(token, options={"verify_signature": False})  # décodage sans vérification !

# Après — sécurisé : vérification de signature active, algorithme explicite
payload = jwt.decode(token, secret_or_public_key, algorithms=["HS256"])
```

## Checklist de vérification post-patch
- [ ] Aucune décision d'authentification/autorisation ne s'appuie sur un décodage JWT sans vérification de signature.
- [ ] La liste d'algorithmes acceptés est explicite et ne contient jamais `none`.
- [ ] Un test de non-régression confirme qu'un token `alg: none` avec signature vide est rejeté.
- [ ] La bibliothèque JWT est à jour (les versions récentes rejettent `none` par défaut, mais la configuration doit rester explicite).
