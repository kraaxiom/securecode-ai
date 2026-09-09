# JWT Algorithm Confusion — Node/Express

`vulnerable.js` appelle `jwt.verify(token, publicKey)` sans restreindre l'algorithme attendu, laissant la bibliothèque se fier au champ `alg` du token (CWE-347, Improper Verification of Cryptographic Signature). Un attaquant peut alors signer un token en HS256 en utilisant la clé publique RSA comme secret partagé, ce qui contourne la vérification.

`fixed.js` impose explicitement `algorithms: ['RS256']` lors de l'appel à `jwt.verify()`, ce qui exclut structurellement toute famille d'algorithme différente et empêche la réutilisation de la clé publique comme secret HMAC.
