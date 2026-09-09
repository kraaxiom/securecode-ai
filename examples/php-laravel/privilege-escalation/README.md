# Privilege Escalation

Le code vulnérable applique n'importe quel rôle transmis par le client sans vérifier que l'appelant a lui-même le droit d'accorder ce niveau de privilège, permettant une escalade verticale vers 'admin'. La version corrigée exige que l'appelant passe la vérification `canGrantRole()` avant toute attribution, et invalide les tokens existants de l'utilisateur modifié pour empêcher la persistance d'anciens privilèges. Cette faille correspond à CWE-269 (Improper Privilege Management), tel qu'indiqué dans `knowledge/authorization/privilege-escalation.md`.
