# Faille : injection de commande (CWE-78).
# L'hôte fourni par l'utilisateur est concaténé dans une commande shell
# exécutée via subprocess avec shell=True.
import subprocess
from django.http import JsonResponse


def ping_host(request):
    host = request.GET.get("host", "")

    output = subprocess.check_output(
        f"ping -c 3 {host}", shell=True
    )

    return JsonResponse({"output": output.decode()})
