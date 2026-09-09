# SQL Injection générique (CWE-89)

Le contrôleur vulnérable concatène directement le paramètre `name` dans une requête SQL brute via `DB::select()`, permettant à un attaquant d'altérer la logique de la requête. La correction remplace la concaténation par un paramètre lié (`?`), délégant l'échappement au driver PDO sous-jacent et éliminant toute la classe de vulnérabilité (CWE-89 : Improper Neutralization of Special Elements used in an SQL Command).
