# Vue Django corrigée — CWE-78 (OS Command Injection)
# Correction : validation stricte de "host" par liste blanche (regex),
# puis exécution via une API à tableau d'arguments (shell=False), qui
# évite toute interprétation shell des méta-caractères.

import re
import subprocess

from django.http import HttpResponse, HttpResponseBadRequest

HOST_RE = re.compile(r"^[a-zA-Z0-9.\-]+$")


def ping_host(request):
    host = request.GET.get("host", "")

    if not HOST_RE.match(host):
        return HttpResponseBadRequest("Hôte invalide")

    # Sécurisé : tableau d'arguments, pas d'interprétation shell
    result = subprocess.run(
        ["ping", "-c", "4", host],
        shell=False,
        capture_output=True,
        text=True,
    )

    return HttpResponse(result.stdout, content_type="text/plain")
