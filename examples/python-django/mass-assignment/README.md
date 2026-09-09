# Mass Assignment — Python/Django

`vulnerable.py` utilise un `UserSerializer` avec `fields = "__all__"`, exposant en écriture des champs sensibles comme `role` ou `is_staff` (CWE-915, Improperly Controlled Modification of Dynamically-Determined Object Attributes). Un attaquant peut injecter ces champs dans une requête de mise à jour de profil.

`fixed.py` introduit un `UserUpdateSerializer` séparé, restreint explicitement à `["name", "email"]`, empêchant toute modification des attributs sensibles via cette route, conformément au principe de liste blanche.
