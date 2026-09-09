# Forced Browsing — CWE-425

La version vulnérable sert des documents internes via une URL prévisible (`/documents/{filename}`), sans authentification ni vérification d'autorisation, en s'appuyant uniquement sur le fait que le nom exact du fichier n'est censé être connu de personne d'autre (sécurité par l'obscurité). Un attaquant qui devine ou énumère des noms de fichiers plausibles (`rapport-2025.pdf`, `contrat-client-42.pdf`) accède directement au contenu sans jamais avoir été autorisé.

La version corrigée exige une authentification (`@PreAuthorize("isAuthenticated()")`) puis vérifie explicitement, via une liste d'accès stockée en base, que le document demandé est bien autorisé pour l'utilisateur courant — indépendamment du fait qu'il en connaisse ou non le nom exact. Une validation du chemin résolu (`normalize()` + `startsWith`) est ajoutée en complément pour empêcher toute traversée de répertoire.

**Référence** : CWE-425 (Direct Request / Forced Browsing).
