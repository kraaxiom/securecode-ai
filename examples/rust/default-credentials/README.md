# Default Credentials

Le handler `provision_admin` créait un compte administrateur avec un mot de passe fixe et documenté publiquement, réutilisable tel quel en production (CWE-1392). La version corrigée génère un mot de passe temporaire aléatoire via `rand`, force son changement via le flag `must_change_password`, et le transmet uniquement par un canal hors-bande sécurisé, jamais en clair dans les journaux.
