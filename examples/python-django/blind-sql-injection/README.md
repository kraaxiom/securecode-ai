# Blind SQL Injection

La version vulnérable concatène le paramètre `username` directement dans la requête SQL brute exécutée via `connection.cursor()`, ce qui correspond à CWE-89 (Improper Neutralization of Special Elements used in an SQL Command). Même sans affichage direct du résultat, un attaquant peut déduire des informations en observant le booléen `exists` ou le temps de réponse. La version corrigée utilise une requête paramétrée (`%s` avec liste de paramètres liés), éliminant toute concaténation de chaîne dans la requête SQL.
