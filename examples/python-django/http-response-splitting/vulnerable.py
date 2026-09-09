# Faille : fractionnement de réponse HTTP (CWE-113).
# La valeur "next" est écrite directement dans l'en-tête Location sans
# validation, permettant potentiellement l'injection d'une seconde
# réponse HTTP si des caractères CR/LF passent le filtrage du runtime.
from django.http import HttpResponse


def redirect_after_login(request):
    next_url = request.GET.get("next", "/dashboard")

    response = HttpResponse(status=302)
    response["Location"] = next_url

    return response
