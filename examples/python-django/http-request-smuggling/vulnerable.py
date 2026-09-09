# Faille : contrebande de requêtes HTTP (CWE-444).
# La vue applicative fait confiance à la requête reçue sans vérifier la
# cohérence entre Content-Length et Transfer-Encoding, alors que
# l'application est déployée derrière un proxy/reverse-proxy en chaîne.
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt


@csrf_exempt
def api_endpoint(request):
    # Aucune vérification de cohérence des en-têtes de délimitation de
    # requête avant traitement du corps.
    payload = request.body
    return JsonResponse({"received_bytes": len(payload)})
