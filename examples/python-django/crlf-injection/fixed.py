# Correction : l'entrée est validée par une expression régulière stricte
# (chemin relatif sans \r\n) et repliée sur une valeur par défaut sûre en
# cas de non-conformité, avant d'utiliser l'API de redirection de Django.
import re
from django.http import HttpResponseRedirect

SAFE_PATH_RE = re.compile(r"^/[^\r\n]*$")


def redirect_view(request):
    next_url = request.GET.get("next", "/")

    if not SAFE_PATH_RE.match(next_url):
        next_url = "/"

    return HttpResponseRedirect(next_url)
