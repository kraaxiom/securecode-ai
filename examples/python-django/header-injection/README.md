# Header Injection

La version vulnérable insère `filename` directement dans l'en-tête `Content-Disposition` par f-string, ce qui correspond à CWE-113 : une valeur contenant `\r\n` pourrait injecter des en-têtes supplémentaires. La version corrigée supprime les caractères de contrôle et applique `os.path.basename()` pour ne retenir qu'un nom de fichier simple, avant insertion dans l'en-tête.
