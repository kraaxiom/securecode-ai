## Header Injection (CWE-113)

Le code vulnérable insère le paramètre `filename` directement dans l'en-tête `Content-Disposition` sans supprimer les caractères `\r`/`\n`, permettant l'injection d'en-têtes supplémentaires ou la modification de la réponse. La correction supprime les caractères de contrôle, applique `path.basename()` puis valide le nom de fichier avec une liste blanche stricte de caractères autorisés avant de l'insérer dans l'en-tête.
