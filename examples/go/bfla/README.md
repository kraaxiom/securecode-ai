# bfla (CWE-862)

La version vulnerable controle uniquement que l'appelant est authentifie, sans jamais verifier qu'il dispose de la fonction/permission necessaire pour supprimer un utilisateur, ce qui permet a un utilisateur standard d'appeler une fonction reservee aux administrateurs (Broken Function Level Authorization). La version corrigee ajoute une verification explicite de la permission requise (`users:delete`) avant d'executer l'action. C'est la remediation standard recommandee contre CWE-862.
