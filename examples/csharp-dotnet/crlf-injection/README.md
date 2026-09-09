# CRLF Injection (CWE-93)

Le code vulnérable écrit directement la valeur du paramètre `next` dans l'en-tête `Location`, ce qui permet à un attaquant d'injecter des séquences `\r\n` pour ajouter des en-têtes arbitraires ou scinder la réponse HTTP. La correction valide que `next` est un chemin relatif sans caractère de contrôle via une expression régulière stricte, puis utilise la méthode sûre `Redirect()` d'ASP.NET Core, qui encode correctement la valeur.
