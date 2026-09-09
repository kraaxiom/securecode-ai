# BFLA — Broken Function Level Authorization

Le code vulnérable protège la suppression d'un utilisateur uniquement par le middleware `auth`, sans vérifier que l'appelant dispose du rôle administrateur requis pour cette fonction. La version corrigée ajoute une `UserPolicy` centralisée (`can:delete`) qui refuse l'action à tout utilisateur qui n'est pas admin, indépendamment de la visibilité de la route dans l'interface. Cette faille correspond à CWE-862 (Missing Authorization), tel qu'indiqué dans `knowledge/authorization/bfla.md`.
