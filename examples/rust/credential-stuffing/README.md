# Credential Stuffing

Le handler `login` ne corrélait pas le volume de tentatives par source et n'imposait pas la MFA, laissant un attaquant tester en masse des couples identifiant/mot de passe issus de fuites externes (CWE-307). La version corrigée ajoute un `VelocityTracker` qui bloque une IP dépassant un seuil de tentatives global tous comptes confondus, et exige une validation MFA après authentification réussie lorsque celle-ci est activée sur le compte.
