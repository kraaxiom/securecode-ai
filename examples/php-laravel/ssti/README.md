# Server-Side Template Injection (CWE-1336)

Le code vulnérable concatène le paramètre `name` directement dans le texte du template avant sa compilation par Twig, permettant au moteur d'interpréter une syntaxe de template injectée par l'utilisateur. La correction fige le template en un fichier/source statique et ne transmet plus `name` que comme variable de contexte échappée par Twig, séparant strictement structure de template et données utilisateur (CWE-1336 : Improper Neutralization of Special Elements Used in a Template Engine).
