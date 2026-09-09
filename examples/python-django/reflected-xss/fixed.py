# Correctif : Reflected XSS (CWE-79) — suppression du filtre |safe ; le
# terme de recherche est affiché avec l'échappement automatique par défaut
# du moteur de templates Django.

from django.shortcuts import render


def search_view(request):
    """Vue de recherche : le terme saisi est réaffiché dans la page de résultats."""
    term = request.GET.get("q", "")

    # Le template utilise désormais {{ term }} (sans |safe) :
    # <h1>Résultats pour : {{ term }}</h1>
    # Django échappe automatiquement les caractères spéciaux HTML.
    return render(request, "search/results.html", {"term": term})
