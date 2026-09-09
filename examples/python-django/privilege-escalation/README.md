# Privilege Escalation — Python/Django

`vulnerable.py` met à jour le rôle d'un utilisateur directement depuis la valeur envoyée par le client, sans vérifier que l'appelant est lui-même autorisé à accorder ce niveau de privilège (CWE-269, Improper Privilege Management).

`fixed.py` ajoute une vérification `can_grant_role` avant toute attribution de rôle et invalide les sessions actives de l'utilisateur cible après le changement, empêchant la persistance d'anciens privilèges et l'auto-attribution non autorisée.
