# mass-assignment (CWE-915)

La version vulnerable deserialise directement le JSON attaquant dans le modele persiste complet, y compris les champs sensibles `Role` et `IsAdmin`, permettant a un utilisateur de s'auto-promouvoir administrateur en ajoutant simplement ces champs a sa requete. La version corrigee introduit un DTO d'entree explicite (`UpdateProfileRequest`) qui ne liste que les champs autorises a etre modifies, et copie ces champs un par un vers le modele. C'est la remediation standard recommandee contre CWE-915.
