# Header Injection (CWE-113)

Le code vulnérable concatène directement le paramètre `filename` dans les en-têtes `Content-Disposition` et `X-Requested-File`, permettant à un attaquant d'injecter des séquences CR/LF pour ajouter des en-têtes HTTP arbitraires à la réponse. La version corrigée supprime systématiquement les caractères `\r` et `\n` puis applique une validation par liste blanche de caractères (`^[\w.-]+$`), rejetant toute valeur invalide avec un code 400 et éliminant tout vecteur d'injection d'en-tête tout en conservant le téléchargement légitime.
