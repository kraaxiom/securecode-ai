# Faille : injection CRLF (CWE-93).
# La valeur "next" provenant de l'utilisateur est écrite telle quelle dans
# l'en-tête Location, sans filtrage des caractères \r\n.
from django.http import HttpResponse


def redirect_view(request):
    next_url = request.GET.get("next", "/")

    response = HttpResponse(status=302)
    response["Location"] = next_url

    return response
