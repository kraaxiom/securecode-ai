# JWT Algorithm Confusion — CWE-347

La version vulnérable vérifie la signature d'un JWT en se fiant implicitement à l'algorithme déclaré dans son propre header, sans jamais l'imposer côté serveur. Un attaquant en possession de la clé publique RS256 de l'application (souvent publiquement accessible) peut forger un token en HS256 en utilisant cette clé publique comme secret HMAC : le serveur, qui ne restreint pas l'algorithme accepté, valide alors une signature forgée par l'attaquant lui-même.

La version corrigée impose explicitement `RS256` comme unique algorithme accepté, indépendamment de ce que déclare le header du token entrant, avec une vérification redondante après le parsing en défense en profondeur. Tout token utilisant un autre algorithme (HS256 notamment) est rejeté avant que sa signature ne soit considérée comme valide.

**Référence** : CWE-347 (Improper Verification of Cryptographic Signature).
