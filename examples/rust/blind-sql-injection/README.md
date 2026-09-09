# Blind SQL Injection (CWE-89)

Le code vulnérable concatène l'entrée utilisateur dans une requête SQL, même lorsque seul un booléen `exists` est renvoyé, ce qui permet une exfiltration progressive par inférence de comportement. La correction remplace la concaténation par une requête préparée `sqlx` avec paramètre lié (`$1`), éliminant tout risque d'altération de la logique SQL. La réponse reste uniforme dans les deux cas afin de ne pas offrir de canal temporel ou de contenu exploitable, conformément à CWE-89.
