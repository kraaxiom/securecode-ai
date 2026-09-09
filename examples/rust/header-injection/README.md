# Header Injection (CWE-113)

Le code vulnérable insère le nom de fichier utilisateur brut dans l'en-tête `Content-Disposition` sans supprimer les caractères `\r`/`\n`, ce qui autorise l'injection d'en-têtes supplémentaires. La correction filtre ces caractères de contrôle puis n'extrait que le nom de base du chemin (`Path::file_name`) avant de reconstruire l'en-tête, appliquant la validation par liste blanche recommandée pour CWE-113.
