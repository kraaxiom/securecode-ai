## SSRF classique (CWE-918)

Le code vulnérable transmet l'URL fournie par l'utilisateur directement à `fetch()` pour générer un aperçu, sans aucune whitelist ni validation de la destination, permettant à un attaquant de forcer le serveur à interroger des ressources internes. La correction restreint les requêtes sortantes à une whitelist de domaines métier, résout le DNS et rejette toute IP privée/loopback/link-local, désactive les redirections automatiques et applique un timeout.
