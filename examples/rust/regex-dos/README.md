# ReDoS — Regular Expression Denial of Service (CWE-1333)

Le motif vulnérable emploie des quantificateurs imbriqués (`(a+)+`) sur une entrée non bornée, un schéma classiquement vecteur de backtracking catastrophique. La correction réécrit la regex sous une forme non ambiguë et impose une limite de longueur sur l'entrée avant tout appel au moteur, conformément au principe de défense en profondeur du guide de remédiation, ce qui neutralise CWE-1333 même si le moteur sous-jacent change.
