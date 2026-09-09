# XSS réfléchi

`vulnerable.js` concatène directement le paramètre de requête `q` dans la réponse HTML renvoyée par le serveur, sans encodage contextuel. `fixed.js` fait passer la même donnée par un moteur de templates EJS dont l'échappement automatique (`<%= %>`) reste actif, neutralisant tout balisage HTML injecté dans le paramètre avant affichage. Il s'agit de CWE-79 (Improper Neutralization of Input During Web Page Generation), la variante réfléchie se caractérisant par une exécution limitée à la réponse immédiate, sans persistance côté serveur.
