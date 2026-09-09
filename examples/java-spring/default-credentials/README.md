# Default Credentials — CWE-1392

La version vulnérable crée automatiquement un compte administrateur avec un identifiant et un mot de passe fixes (`admin` / `admin123`) codés en dur au démarrage de l'application, sans jamais forcer leur changement. Ces valeurs, souvent documentées publiquement (README, image Docker, documentation d'installation), sont triviales à deviner ou à retrouver, et constituent une des causes les plus fréquentes de compromission d'applications exposées.

La version corrigée génère un mot de passe aléatoire cryptographiquement sûr à la première initialisation, l'affiche une seule fois dans les logs de démarrage pour récupération par l'opérateur, le stocke haché (BCrypt), et positionne un indicateur `mustChangePassword` obligeant son renouvellement dès la première connexion. Aucune valeur fixe n'est jamais codée en dur ni exposée dans un dépôt de code.

**Référence** : CWE-1392 (Use of Default Credentials).
