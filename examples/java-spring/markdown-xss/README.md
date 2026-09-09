# Markdown XSS — CWE-79

La version vulnérable convertit le Markdown soumis par l'utilisateur en HTML (via CommonMark) et l'insère tel quel dans la page (`th:utext`). Comme le Markdown autorise nativement le HTML brut en ligne, un commentaire contenant `<img src=x onerror=alert(document.cookie)>` ou `<script>` est fidèlement restitué et exécuté dans le navigateur des autres utilisateurs — le rendu Markdown n'offre aucune garantie de sécurité par lui-même.

La version corrigée fait passer le HTML généré par le moteur Markdown dans une liste blanche stricte de balises et attributs via l'OWASP Java HTML Sanitizer : seules quelques balises de mise en forme (`p`, `strong`, `a`, `ul`…) sont conservées, les protocoles d'URL sont restreints à `https`, et `rel="nofollow"` est forcé sur les liens. Toute balise ou attribut porteur de script (`<script>`, `on*`, `javascript:`) est supprimé avant l'affichage.

**Référence** : CWE-79 (Improper Neutralization of Input During Web Page Generation).
