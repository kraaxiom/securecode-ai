# ReDoS — Regular Expression Denial of Service (CWE-1333)

La regex de validation d'email dans le code vulnérable contient des quantificateurs imbriqués (`([a-zA-Z0-9]+)+`) appliqués sans limite de longueur, ce qui peut provoquer un temps de calcul exponentiel sur une entrée pathologique et bloquer le processus PHP. La correction remplace la regex par un motif non ambigu, borne la longueur de l'entrée à 254 caractères avant traitement, et fixe une limite de backtracking PCRE (CWE-1333 : Inefficient Regular Expression Complexity).
