# IDOR — Insecure Direct Object Reference — CWE-639

La version vulnérable récupère et sert une facture uniquement à partir de son identifiant numérique séquentiel transmis dans l'URL, sans jamais vérifier que l'utilisateur connecté en est bien le destinataire. Faire varier l'identifiant (`/invoices/1001`, `/invoices/1002`, ...) suffit à télécharger les factures d'autres clients.

La version corrigée vérifie explicitement, avant toute résolution du fichier, que la facture demandée appartient bien à l'utilisateur courant (`belongsToUser`), et renvoie une réponse 404 identique que la facture existe pour un autre utilisateur ou n'existe pas du tout, afin de ne pas confirmer l'existence d'identifiants valides appartenant à des tiers. Le remplacement d'identifiants séquentiels par des UUID peut réduire la surface d'énumération, mais ne remplace jamais ce contrôle d'autorisation.

**Référence** : CWE-639 (Authorization Bypass Through User-Controlled Key).
