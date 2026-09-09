# nosql-injection (CWE-943)

La version vulnerable decode le corps JSON de la requete directement en `bson.M` et le transmet tel quel comme filtre MongoDB, permettant d'injecter des operateurs (`$ne`, `$gt`...) au lieu de valeurs scalaires pour contourner l'authentification. La version corrigee decode les identifiants dans une struct typee (`string`/`string`), ce qui rejette tout objet JSON injecte, puis reconstruit explicitement le filtre avec des valeurs scalaires uniquement.
