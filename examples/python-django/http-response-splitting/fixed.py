# Correction : la destination de redirection est validée contre une
# liste blanche de chemins internes connus, et rejette explicitement
# tout schéma/hôte externe (redirection ouverte) via urlparse.
from urllib.parse import urlparse
from django.http import HttpResponseRedirect

ALLOWED_PATHS = {"/dashboard", "/profile", "/account"}


def redirect_after_login(request):
    next_url = request.GET.get("next", "/dashboard")
    parsed = urlparse(next_url)

    if parsed.scheme or parsed.netloc or next_url not in ALLOWED_PATHS:
        next_url = "/dashboard"

    return HttpResponseRedirect(next_url)
