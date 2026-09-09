# Header Injection (CWE-113)

Le code vulnérable insère le paramètre `filename` directement dans l'en-tête `Content-Disposition`, ce qui permet d'injecter des séquences `\r\n` pour ajouter des en-têtes non prévus. La correction supprime les caractères de contrôle, réduit le nom au nom de base via `Path.GetFileName`, puis délègue la construction de l'en-tête à la surcharge sûre `File(bytes, contentType, fileName)` d'ASP.NET Core.
