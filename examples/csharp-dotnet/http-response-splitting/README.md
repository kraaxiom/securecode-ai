# HTTP Response Splitting (CWE-113)

La version vulnérable écrit le paramètre `next` directement dans l'en-tête `Location`, permettant à un attaquant d'injecter des séquences CR/LF pour fractionner la réponse HTTP et insérer du contenu arbitraire perçu par un cache ou un autre utilisateur. La correction restreint les cibles de redirection à une liste blanche (`AllowedPaths`) et utilise la méthode `Redirect()` d'ASP.NET Core, qui encode correctement la valeur.
