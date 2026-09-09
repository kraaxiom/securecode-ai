# Mutation XSS (mXSS) — CWE-79

La version vulnérable sanitise la bio utilisateur avec une liste blanche Jsoup trop permissive (balises `svg`, `math`, `noscript`, attribut `style` autorisé), puis réinjecte le résultat côté client via `innerHTML` pour l'aperçu instantané. Le problème du mutation XSS est que le HTML jugé "propre" au moment de la sanitisation peut être réinterprété différemment par le moteur de parsing du navigateur lors d'une seconde passe (ex: `innerHTML` sur du contenu déjà sanitisé) : des structures ambiguës peuvent "muter" en un payload actif qui n'existait pas sous cette forme initialement.

La version corrigée réduit la liste blanche au strict nécessaire (`Safelist.basic()`, sans attributs de style ni balises conteneurs ambiguës) et surtout applique une seconde sanitisation, dédiée au DOM, juste avant toute insertion côté client via `DOMPurify.sanitize(html, { SANITIZE_DOM: true })` — la sanitisation serveur seule n'est jamais suffisante lorsque le contenu retraverse un pipeline de parsing HTML.

**Référence** : CWE-79 (Improper Neutralization of Input During Web Page Generation), variante mutation XSS.
