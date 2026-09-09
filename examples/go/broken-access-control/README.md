# broken-access-control (CWE-284)

La version vulnerable ne fait reposer la protection de l'endpoint d'administration que sur l'affichage cote client, sans aucun controle serveur, ce qui le rend accessible directement a n'importe quel appelant. La version corrigee applique une verification explicite du role cote serveur selon une logique deny-by-default, independamment de ce que montre l'interface. C'est la remediation standard recommandee contre CWE-284.
