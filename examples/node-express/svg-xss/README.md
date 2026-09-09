# XSS via SVG

`vulnerable.js` stocke et sert un fichier SVG uploadé par l'utilisateur tel quel, sans sanitisation de son contenu XML, depuis la même origine que l'application. `fixed.js` sanitise le contenu SVG à l'upload via une bibliothèque dédiée, puis sert le fichier en téléchargement forcé (`Content-Disposition: attachment`) plutôt qu'en affichage inline. Cette vulnérabilité relève de CWE-79 (Improper Neutralization of Input During Web Page Generation), le format SVG étant un document XML pouvant embarquer des éléments actifs interprétés par le navigateur.
