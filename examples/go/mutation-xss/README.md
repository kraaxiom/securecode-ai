# mutation-xss (CWE-79)

La version vulnerable "sanitise" le HTML utilisateur avec une expression reguliere qui supprime les balises `<script>` par correspondance textuelle : cette approche ignore la grammaire reelle du HTML et les regles de reparsing du navigateur, laissant passer des payloads qui se recomposent en vecteur actif une fois reinterpretes par le DOM. La version corrigee utilise `bluemonday`, qui s'appuie sur un vrai parseur HTML (arbre DOM conforme) avant d'appliquer une politique d'allowlist stricte.
