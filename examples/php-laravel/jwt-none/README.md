# JWT alg: none

La version corrigée impose explicitement l'algorithme `HS256` lors de la vérification du token, excluant structurellement l'acceptation d'un token signé avec `alg: none`, contrairement à la version vulnérable qui laissait l'algorithme vide. La faille correspond à CWE-347 (Improper Verification of Cryptographic Signature), comme indiqué dans `knowledge/auth/jwt-none.md`.
