# Stored XSS (CWE-79)

Les commentaires persistés en base étaient enveloppés dans `mark_safe()` avant affichage, contournant l'échappement automatique de Django et permettant à un payload stocké une seule fois de s'exécuter chez chaque visiteur de la page. Le correctif retire `mark_safe()` et laisse le template afficher `{{ comment.body }}` directement, ce qui applique l'échappement HTML par défaut à chaque rendu.
