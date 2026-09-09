# BOLA — Broken Object Level Authorization

Le code vulnérable récupère une commande uniquement par son ID (`findOrFail`), sans vérifier que l'appelant en est bien le propriétaire, ce qui permet d'accéder aux commandes d'autres utilisateurs en changeant l'ID. La version corrigée ajoute une clause `where('user_id', auth()->id())` directement dans la requête, filtrant l'appartenance au niveau de la donnée elle-même plutôt qu'en post-traitement. Cette faille correspond à CWE-639 (Authorization Bypass Through User-Controlled Key), tel qu'indiqué dans `knowledge/authorization/bola.md`.
