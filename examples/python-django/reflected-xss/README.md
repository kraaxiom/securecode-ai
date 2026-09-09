# Reflected XSS (CWE-79)

La vue de recherche réaffichait le paramètre `q` dans le template avec le filtre `|safe`, désactivant volontairement l'échappement automatique de Django sur une donnée directement contrôlée par l'attaquant via l'URL. Le correctif retire simplement `|safe` : Django échappe alors la valeur par défaut, neutralisant toute tentative d'injection HTML/JS réfléchie dans la page de résultats.
