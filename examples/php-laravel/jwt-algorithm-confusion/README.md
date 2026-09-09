# JWT Algorithm Confusion

La version corrigée impose explicitement `RS256` comme unique algorithme accepté lors de la vérification, au lieu de laisser l'algorithme implicite. Cela empêche un attaquant de forger un token signé en HS256 en réutilisant la clé publique RSA de l'application comme secret HMAC. La faille correspond à CWE-347 (Improper Verification of Cryptographic Signature), comme indiqué dans `knowledge/auth/jwt-algorithm-confusion.md`.
