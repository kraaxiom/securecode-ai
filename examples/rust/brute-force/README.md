# Brute Force

Le handler `login` tentait l'authentification sans aucune limitation du nombre d'essais par compte ou par IP, permettant un brute force automatisé (CWE-307). La version corrigée introduit un `AttemptTracker` qui verrouille temporairement la combinaison IP+compte après 5 échecs sur une fenêtre de 15 minutes, réinitialise le compteur en cas de succès, et conserve un message d'erreur générique ne révélant pas l'existence du compte.
