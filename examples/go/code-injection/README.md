# code-injection (CWE-94)

La version vulnerable evalue dynamiquement une expression fournie par l'utilisateur via un moteur d'evaluation, permettant l'execution de code arbitraire. La version corrigee remplace l'evaluation dynamique par une logique metier explicite (switch sur une allowlist d'operations), supprimant tout interpreteur expose a l'utilisateur.
