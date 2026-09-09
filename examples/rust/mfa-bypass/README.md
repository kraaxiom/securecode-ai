# MFA Bypass

Le handler `login` émettait un token pleinement privilégié dès validation du mot de passe, sans jamais vérifier le second facteur côté serveur, rendant la MFA purement facultative (CWE-287). La version corrigée introduit un token intermédiaire à portée restreinte (`partial_token`) et n'émet le token complet que via `mfa_verify`, après validation serveur explicite du code, empêchant tout accès privilégié sans second facteur validé.
