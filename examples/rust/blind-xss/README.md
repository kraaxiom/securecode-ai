# Blind XSS

Le handler `view_ticket` réaffichait le message d'un ticket de support soumis par un utilisateur externe directement dans le HTML du back-office admin, sans encodage de sortie (CWE-79). La version corrigée applique un échappement HTML systématique sur cette donnée, y compris pour un affichage interne, car une donnée externe reste non fiable quel que soit l'écran qui la consulte. Cela empêche l'exécution d'un script dans le contexte de l'administrateur, cible privilégiée du XSS aveugle.
