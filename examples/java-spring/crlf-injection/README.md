# CRLF Injection (CWE-93)

Le code vulnérable écrit directement le paramètre `next` dans l'en-tête `Location` via `HttpServletResponse.setHeader`, ce qui permet à un attaquant d'injecter des séquences `\r\n` pour ajouter des en-têtes arbitraires ou scinder la réponse HTTP. La version corrigée valide que `next` est un chemin relatif sans caractère de contrôle avant utilisation, et s'appuie sur `HttpHeaders`/`ResponseEntity` de Spring pour construire la redirection, éliminant tout risque d'injection dans les en-têtes.
