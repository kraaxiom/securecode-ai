# Vue Django vulnérable — CWE-78 (OS Command Injection)
# Le paramètre "host" fourni par l'utilisateur est concaténé directement
# dans une commande shell exécutée via subprocess.run(shell=True). Un
# attaquant peut injecter des méta-caractères shell (";", "|", "&&", "`")
# pour exécuter des commandes arbitraires sur le serveur.

import subprocess

from django.http import HttpResponse


def ping_host(request):
    host = request.GET.get("host", "")

    # Vulnérable : interpolation dans une chaîne shell
    result = subprocess.run(
        f"ping -c 4 {host}",
        shell=True,
        capture_output=True,
        text=True,
    )

    return HttpResponse(result.stdout, content_type="text/plain")
