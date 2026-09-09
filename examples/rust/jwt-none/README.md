# JWT alg: none

Le handler `admin_data` décodait le token en lisant directement le payload base64 sans jamais vérifier la signature, acceptant de fait un token `alg: none` forgé sans secret (CWE-347). La version corrigée utilise `jsonwebtoken::decode` avec `Validation::new(Algorithm::HS256)`, qui impose une vérification de signature active et exclut structurellement l'algorithme `none`.
