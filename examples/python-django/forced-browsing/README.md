# Forced Browsing — Python/Django

`vulnerable.py` affiche la page de confirmation de paiement sans vérifier que les étapes précédentes du flux ont réellement été complétées côté serveur (CWE-425, Direct Request). Il suffit de connaître l'URL pour y accéder directement.

`fixed.py` vérifie l'état d'avancement réel du flux via `request.session` avant d'autoriser l'accès à l'étape de confirmation, redirigeant vers le début du parcours si le paiement n'a pas été authentifié comme autorisé.
