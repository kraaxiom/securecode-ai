# Privilege Escalation

La version vulnérable applique directement le rôle envoyé dans `req.body` lors d'une mise à jour, sans vérifier que l'appelant est lui-même autorisé à accorder ce niveau de privilège, permettant une auto-attribution du rôle admin. La version corrigée introduit une vérification `canGrantRole` fondée sur la hiérarchie des rôles avant toute attribution, puis invalide les sessions actives de l'utilisateur ciblé pour empêcher la persistance d'anciens privilèges. Cela correspond à **CWE-269 (Improper Privilege Management)**, catégorie OWASP A01:2021-Broken Access Control.
