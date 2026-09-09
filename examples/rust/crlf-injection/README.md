# CRLF Injection (CWE-93)

Le code vulnérable insère directement le paramètre `next` dans l'en-tête `Location` sans vérifier la présence de `\r`/`\n`, ouvrant la voie à l'injection d'en-têtes supplémentaires. La correction rejette toute valeur contenant des caractères de contrôle et impose un format de chemin local (`starts_with('/')`), retombant sur `/` par défaut sinon. Cela neutralise le vecteur d'injection décrit par CWE-93.
