# Blind SQL Injection (CWE-89)

Le code vulnérable concatène directement le paramètre `username` dans la requête SQL, ce qui permet une injection SQL aveugle même si seul un booléen `exists` est renvoyé au client. La version corrigée utilise une `SqlCommand` paramétrée avec `@Username` lié via `Parameters.AddWithValue`, empêchant toute modification de la structure de la requête. Cela élimine le canal d'inférence booléen tout en conservant un comportement fonctionnel identique.
