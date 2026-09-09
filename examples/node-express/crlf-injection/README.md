## CRLF Injection (CWE-93)

Le code vulnérable écrit le paramètre `next` directement dans l'en-tête HTTP `Location` sans filtrer les séquences `\r\n`, permettant à un attaquant d'injecter des en-têtes supplémentaires ou de scinder la réponse. La correction valide la valeur avec une expression régulière stricte n'autorisant qu'un chemin relatif sans caractère de contrôle, puis utilise `res.redirect()` pour une gestion sûre de l'en-tête.
