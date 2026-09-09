# Time-Based Blind SQL Injection (CWE-89)

Le code vulnérable concatène le paramètre `id` dans une requête SQL exécutée sans retour de données ni erreur détaillée, permettant à un attaquant d'injecter une fonction de pause (ex: `SLEEP`) et de déduire des informations via le délai de réponse. La correction remplace la concaténation par un paramètre lié et configure un timeout d'exécution (`PDO::ATTR_TIMEOUT`), éliminant la classe de vulnérabilité et limitant l'impact résiduel (CWE-89 : Improper Neutralization of Special Elements used in an SQL Command).
