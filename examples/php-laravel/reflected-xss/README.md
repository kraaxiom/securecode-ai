# Reflected XSS (CWE-79)

Le paramètre de recherche `q` était réaffiché dans la vue Blade via la syntaxe `{!! !!}`, désactivant explicitement l'échappement automatique et permettant l'exécution de HTML/JS injecté via un lien piégé. La correction remplace `{!! $term !!}` par `{{ $term }}`, réactivant l'échappement contextuel automatique de Blade sans toucher au contrôleur. Cela élimine le vecteur de réflexion immédiate décrit par le CWE-79 (Improper Neutralization of Input During Web Page Generation) propre au XSS réfléchi.
