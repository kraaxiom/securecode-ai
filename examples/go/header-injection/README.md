# header-injection (CWE-113)

La version vulnerable copie une valeur utilisateur non filtree dans un en-tete de reponse, exposant l'application a l'injection d'en-tetes via CR/LF. La version corrigee restreint la valeur a une allowlist de codes de langue connus, eliminant toute possibilite d'injection.
