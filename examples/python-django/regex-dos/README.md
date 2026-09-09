# ReDoS — CWE-1333

La regex `^([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)+$` contient des quantificateurs imbriqués créant une explosion combinatoire de backtracking sur certaines entrées, ce qui peut geler le worker Django (déni de service). La correction remplace la regex par une forme non ambiguë sans groupes imbriqués et impose une limite de longueur (254 caractères) avant tout test, garantissant un temps d'exécution borné. Voir `rules/remediation/regex-dos.md` pour d'autres langages.
