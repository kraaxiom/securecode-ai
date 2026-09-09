# reflected-xss (CWE-79)

La version vulnerable reflete le parametre de recherche `q` tel quel dans la reponse HTML via `fmt.Fprintf`, permettant l'execution d'un script injecte via un lien controle par l'attaquant. La version corrigee utilise `html/template`, dont l'echappement contextuel automatique neutralise tout balisage avant l'ecriture dans la reponse.
