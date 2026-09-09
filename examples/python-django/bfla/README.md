# BFLA — Python/Django

`vulnerable.py` expose une action de suppression d'utilisateur protégée uniquement par `IsAuthenticated`, sans vérification du rôle de l'appelant (CWE-862, Missing Authorization). N'importe quel utilisateur connecté peut donc supprimer n'importe quel compte.

`fixed.py` ajoute une classe de permission `IsAdmin` appliquée spécifiquement à l'action `destroy` via `get_permissions`, centralisant le contrôle de rôle plutôt que de le disperser. Seuls les utilisateurs ayant le rôle `admin` peuvent désormais exécuter cette fonction sensible.
