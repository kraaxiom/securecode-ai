# Time-Based Blind SQL Injection — CWE-89

Le code vulnérable concatène `id` dans une requête SQL sans renvoyer d'erreur détaillée, ce qui expose l'application à une injection aveugle exploitée via le délai de réponse (ex: fonctions de pause). La correction type strictement l'identifiant et utilise un paramètre lié, éliminant la classe de vulnérabilité ; un timeout d'exécution côté base de données est recommandé en complément. Voir `rules/remediation/time-based-sqli.md` pour d'autres langages.
