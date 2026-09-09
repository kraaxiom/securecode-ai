# SSRF classique — CWE-918

La version vulnérable transmet directement l'URL fournie par l'utilisateur à `RestTemplate` sans aucun contrôle, permettant à un attaquant de forcer le serveur à effectuer des requêtes vers des ressources internes (services d'administration, ports non exposés publiquement, etc.).

La version corrigée impose une liste blanche stricte de schémas (`https` uniquement) et d'hôtes autorisés, puis résout le DNS et vérifie que l'adresse IP obtenue n'appartient pas aux plages privées, loopback ou link-local avant d'effectuer la requête sortante. La whitelist d'hôtes reste le contrôle principal ; la vérification d'IP protège contre le DNS rebinding basique.

**Référence** : CWE-918 (Server-Side Request Forgery).
