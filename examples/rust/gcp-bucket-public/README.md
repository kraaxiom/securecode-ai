# Bucket Google Cloud Storage public (CWE-284)

Le code vulnérable ajoute un binding IAM accordant `roles/storage.objectViewer` au principal `allUsers`, rendant tous les objets du bucket lisibles par quiconque sur Internet sans autorisation. La correction n'accorde jamais l'accès à `allUsers`/`allAuthenticatedUsers` : le rôle de lecture est restreint à un compte de service applicatif nommé, et tout partage ponctuel passe par une URL signée expirant après 15 minutes (`SignedURLOptions`). Élimine la classe de vulnérabilité CWE-284 (Improper Access Control).
