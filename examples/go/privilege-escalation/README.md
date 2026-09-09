# privilege-escalation (CWE-269)

La version vulnerable applique le role cible envoye par le client sans verifier que l'appelant a le privilege d'attribuer des roles ni restreindre les roles attribuables, permettant a un utilisateur standard de s'auto-promouvoir administrateur. La version corrigee exige le privilege `roles:manage`, restreint les roles attribuables via cet endpoint a une liste blanche (excluant `admin`), et interdit a un utilisateur de modifier son propre role. C'est la remediation standard recommandee contre CWE-269.
