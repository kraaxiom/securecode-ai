# BOLA — Python/Django

`vulnerable.py` définit un `OrderViewSet` dont le queryset (`Order.objects.all()`) n'est jamais filtré par utilisateur (CWE-639, Authorization Bypass Through User-Controlled Key). Toute commande est accessible à tout utilisateur authentifié via son ID.

`fixed.py` remplace le queryset statique par un `get_queryset()` filtrant systématiquement sur `user=self.request.user`, appliquant le contrôle d'appartenance directement dans la requête de données plutôt qu'en post-traitement.
