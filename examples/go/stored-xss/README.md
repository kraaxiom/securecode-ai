# stored-xss (CWE-79)

La version vulnerable affiche les commentaires stockes en base tels quels via `fmt.Fprintf`, sans echappement : un script injecte une seule fois s'execute pour chaque visiteur qui consulte la page. La version corrigee utilise `html/template`, qui echappe chaque commentaire selon le contexte HTML au moment de l'affichage, independamment de ce qui a ete stocke.
