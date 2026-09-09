## HTTP Response Splitting (CWE-113)

Le code vulnérable écrit la valeur du paramètre `next` directement dans l'en-tête `Location` sans filtrer les caractères CR/LF, ce qui permet de scinder la réponse HTTP et d'injecter un contenu ou des en-têtes supplémentaires. La correction restreint les cibles de redirection à une liste blanche de chemins connus (`ALLOWED_PATHS`) et utilise `res.redirect()`, qui encode correctement l'en-tête au lieu d'une écriture brute.
