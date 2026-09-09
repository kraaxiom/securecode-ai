# Faille : Reflected XSS (CWE-79) — le paramètre de requête est réinjecté
# dans la réponse HTML en désactivant explicitement l'échappement
# automatique du moteur de templates Django via le filtre |safe.

from django.shortcuts import render


def search_view(request):
    """Vue de recherche : le terme saisi est réaffiché dans la page de résultats."""
    term = request.GET.get("q", "")

    # Le template utilise {{ term|safe }} pour afficher :
    # <h1>Résultats pour : {{ term|safe }}</h1>
    # ce qui désactive l'échappement HTML normalement automatique de Django
    # sur une donnée directement issue de la requête utilisateur.
    return render(request, "search/results.html", {"term": term})
