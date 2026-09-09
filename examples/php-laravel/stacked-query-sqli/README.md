# Stacked Query SQL Injection (CWE-89)

Le code vulnérable concatène le paramètre `name` dans une requête exécutée avec l'émulation de préparation PDO activée, ce qui permet à un attaquant d'ajouter une instruction SQL complète après un point-virgule (ex: `DROP TABLE`). La correction utilise un paramètre lié et désactive explicitement `PDO::ATTR_EMULATE_PREPARES`, empêchant le driver d'exécuter plusieurs instructions dans un seul appel (CWE-89 : Improper Neutralization of Special Elements used in an SQL Command).
