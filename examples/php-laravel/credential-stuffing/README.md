# Credential Stuffing

La version corrigée ajoute une détection de vélocité globale par IP (tous comptes confondus), déclenchant une alerte de sécurité en cas de volume anormal, et impose un second facteur après l'authentification réussie. Ces contrôles complètent le simple rate limiting par compte, insuffisant contre des couples identifiant/mot de passe volés testés sur des comptes différents. La faille correspond à CWE-307 (Improper Restriction of Excessive Authentication Attempts), comme indiqué dans `knowledge/auth/credential-stuffing.md`.
