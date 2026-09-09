# Markdown XSS (CWE-79)

Le rendu Markdown vulnérable activait l'extension `md_in_html`, laissant passer tout HTML brut présent dans le contenu utilisateur, puis marquait le résultat comme sûr sans sanitisation. Le correctif retire ce passthrough HTML et fait passer le HTML généré par `bleach.clean()` avec une liste blanche stricte de balises, d'attributs et de schémas d'URL autorisés, avant de le marquer sûr pour l'affichage.
