# SQL Injection (CWE-89)

Le handler vulnérable construit la requête SQL par concaténation directe du paramètre `name`, exposant l'application à une modification de la structure de la requête. La correction remplace cette concaténation par une requête préparée `sqlx` utilisant un paramètre lié (`$1` + `.bind()`), ce qui empêche toute donnée utilisateur d'être interprétée comme du SQL et élimine la classe de vulnérabilité CWE-89.
