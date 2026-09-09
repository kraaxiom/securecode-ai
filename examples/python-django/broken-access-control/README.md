# Broken Access Control — Python/Django

`vulnerable.py` sert un fichier de facture après authentification uniquement, sans aucune vérification d'autorisation propre à la ressource (CWE-284, Improper Access Control). Tout utilisateur connecté peut télécharger la facture de n'importe qui.

`fixed.py` ajoute un contrôle explicite d'appartenance ou de rôle admin avant de servir le fichier, appliquant le principe "deny by default" : l'accès est refusé sauf autorisation explicite et vérifiable côté serveur.
