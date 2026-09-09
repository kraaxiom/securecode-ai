# SQL Injection — CWE-89

La vue vulnérable insère le paramètre `name` par f-string dans une requête SQL exécutée via `connection.cursor()`, permettant à un attaquant d'altérer la logique de la requête (contournement de filtre, extraction de données). La correction remplace la concaténation par un paramètre lié (`%s` + liste de valeurs), que le driver échappe et traite indépendamment du texte SQL. Voir `rules/remediation/sql-injection.md` pour d'autres langages.
