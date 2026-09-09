# CRLF Injection

La version vulnérable affecte directement le paramètre `next` à l'en-tête `Location` de la réponse, ce qui correspond à CWE-93 : une entrée contenant `\r\n` permettrait d'injecter des en-têtes supplémentaires ou de fractionner la réponse. La version corrigée valide `next_url` avec une expression régulière n'acceptant qu'un chemin relatif sans caractères de contrôle, avec repli sur `/` par défaut.
