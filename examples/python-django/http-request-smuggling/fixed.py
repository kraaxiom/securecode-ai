# Correction : middleware défensif qui rejette explicitement toute
# requête présentant simultanément Content-Length et Transfer-Encoding,
# en complément de la configuration côté proxy/serveur front (HTTP/2 de
# bout en bout recommandé au niveau infrastructure).
from django.http import JsonResponse, HttpResponseBadRequest
from django.views.decorators.csrf import csrf_exempt


def reject_ambiguous_requests(get_response):
    def middleware(request):
        if "CONTENT_LENGTH" in request.META and "HTTP_TRANSFER_ENCODING" in request.META:
            return HttpResponseBadRequest("Requête ambiguë refusée")
        return get_response(request)

    return middleware


@csrf_exempt
def api_endpoint(request):
    payload = request.body
    return JsonResponse({"received_bytes": len(payload)})
