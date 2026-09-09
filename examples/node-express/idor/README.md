# IDOR — Insecure Direct Object Reference

La version vulnérable récupère un document uniquement à partir de l'ID fourni dans l'URL, sans vérifier que ce document appartient à l'utilisateur connecté, ce qui permet d'accéder aux documents d'un autre utilisateur en modifiant l'identifiant. La version corrigée ajoute une clause `ownerId: req.user.id` directement dans la requête de récupération, garantissant que seul le propriétaire réel peut accéder à son document. Cela correspond à **CWE-639 (Authorization Bypass Through User-Controlled Key)**, catégorie OWASP A01:2021-Broken Access Control.
