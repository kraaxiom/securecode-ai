# MFA Bypass — Node/Express

`vulnerable.js` émet un token pleinement privilégié dès que le mot de passe est validé, indépendamment de l'activation de la MFA sur le compte (CWE-287, Improper Authentication). La vérification du second facteur devient une simple information côté client, sans effet réel sur l'accès.

`fixed.js` introduit un token intermédiaire à portée limitée (`scope: 'mfa_pending'`) tant que le code MFA n'est pas validé, et ne délivre le token complet qu'après vérification serveur explicite via `/mfa/verify`.
