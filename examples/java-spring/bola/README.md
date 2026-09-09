# BOLA — Broken Object Level Authorization — CWE-639

La version vulnérable récupère et met à jour une commande uniquement à partir de son identifiant transmis par le client, sans jamais vérifier que l'utilisateur authentifié en est bien le propriétaire. Un utilisateur peut donc accéder aux commandes d'autrui, voire les modifier, en changeant simplement l'ID dans l'URL.

La version corrigée ajoute une condition d'appartenance (`ownerId`) directement dans la requête de récupération des données (`findByIdAndOwner`, `updateIfOwner`), garantissant qu'aucun objet appartenant à un autre utilisateur ne peut jamais être lu ni modifié, quelle que soit la valeur de l'ID fourni.

**Référence** : CWE-639 (Authorization Bypass Through User-Controlled Key), référencé comme API1:2023 Broken Object Level Authorization dans l'OWASP API Security Top 10.
