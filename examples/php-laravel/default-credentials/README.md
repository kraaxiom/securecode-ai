# Default Credentials

La version corrigée remplace le mot de passe admin fixe et documenté par un mot de passe temporaire généré aléatoirement à chaque provisioning, associé à un flag `must_change_password` bloquant l'accès tant qu'il n'a pas été renouvelé. Le secret n'est jamais journalisé en clair et transite par un canal hors-bande sécurisé. La faille correspond à CWE-1392 (Use of Default Credentials), comme indiqué dans `knowledge/auth/default-credentials.md`.
