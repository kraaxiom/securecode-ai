# Faille : DOM-based XSS (CWE-79) — la vue Django renvoie un template qui
# écrit une donnée contrôlable par l'attaquant (paramètre d'URL) directement
# dans le DOM via innerHTML côté client, en dehors de tout échappement du
# moteur de templates serveur. Le rendu HTML est produit dynamiquement en JS.

from django.shortcuts import render


def search_page(request):
    """Page de recherche : le terme est repris tel quel dans un script inline."""
    query = request.GET.get("q", "")

    # Le paramètre est injecté brut dans un bloc <script> du template, pour
    # être ensuite écrit dans le DOM via innerHTML côté navigateur :
    # <div id="result"></div>
    # <script>
    #   document.getElementById('result').innerHTML = "{{ query|safe }}";
    # </script>
    return render(request, "search/results.html", {"query": query})
