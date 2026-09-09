# Correction : l'hôte est validé comme adresse IP avant tout usage, puis
# la commande est exécutée sans shell, avec des arguments passés en liste
# distincte — impossible d'injecter des métacaractères shell.
import ipaddress
import subprocess
from django.http import JsonResponse, HttpResponseBadRequest


def ping_host(request):
    host = request.GET.get("host", "")

    try:
        ipaddress.ip_address(host)
    except ValueError:
        return HttpResponseBadRequest("Adresse IP invalide")

    output = subprocess.check_output(
        ["ping", "-c", "3", host], shell=False
    )

    return JsonResponse({"output": output.decode()})
