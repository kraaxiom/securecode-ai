# IDOR — Python/Django

`vulnerable.py` récupère un document uniquement par son identifiant transmis par le client, sans vérifier son appartenance à l'utilisateur connecté (CWE-639, Authorization Bypass Through User-Controlled Key).

`fixed.py` ajoute la clause `owner=request.user` directement dans `get_object_or_404`, appliquant le filtrage d'appartenance au niveau de la requête de données elle-même plutôt qu'en post-traitement.
