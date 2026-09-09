# Correctif : DOM-based XSS (CWE-79) — on ne réinjecte plus la donnée dans
# un contexte JS via |safe. On la transmet échappée par Django et le script
# client utilise textContent (API texte) plutôt qu'innerHTML comme sink.

from django.shortcuts import render


def search_page(request):
    """Page de recherche : le terme est affiché sans passer par un sink HTML."""
    query = request.GET.get("q", "")

    # Le template n'utilise plus |safe : {{ query }} est échappé par Django.
    # Côté client, le script correspondant utilise textContent au lieu
    # d'innerHTML :
    # <div id="result" data-query="{{ query }}"></div>
    # <script>
    #   document.getElementById('result').textContent =
    #       document.getElementById('result').dataset.query;
    # </script>
    return render(request, "search/results.html", {"query": query})
