# blind-xss (CWE-79)

La version vulnerable affiche le message d'un ticket de support soumis par un utilisateur externe tel quel dans le tableau de bord interne, sans echappement : le script injecte s'execute dans le contexte de session de l'administrateur qui consulte le ticket, sans que l'attaquant n'observe directement le resultat. La version corrigee utilise `html/template`, dont l'echappement contextuel automatique neutralise tout balisage avant insertion dans la page.
