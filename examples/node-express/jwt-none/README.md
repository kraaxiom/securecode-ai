# JWT alg:none — Node/Express

`vulnerable.js` utilise `jwt.decode()`, qui ne vérifie aucune signature, pour extraire les claims d'un token utilisé dans une décision d'autorisation (CWE-347, Improper Verification of Cryptographic Signature). Un attaquant peut forger un token `alg: none` avec une signature vide et un claim `role: admin` arbitraire, sans connaître aucun secret.

`fixed.js` remplace `jwt.decode()` par `jwt.verify()` avec une liste fermée d'algorithmes (`algorithms: ['HS256']`), qui exclut structurellement `none` et rejette tout token dont la signature n'est pas valide.
