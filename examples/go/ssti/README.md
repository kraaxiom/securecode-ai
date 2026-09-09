# ssti (CWE-1336)

La version vulnerable construit la chaine de template a partir de l'entree utilisateur avant de la parser, permettant l'injection de directives de template (Server-Side Template Injection). La version corrigee fige le template en constante statique et ne passe l'entree utilisateur qu'en tant que donnee de contexte, jamais interpretee comme code de template.
