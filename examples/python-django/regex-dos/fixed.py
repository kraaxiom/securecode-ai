# Vue Django corrigée — CWE-1333 (ReDoS)
# Correction : regex non ambiguë (sans quantificateurs imbriqués) et
# limite de longueur imposée sur l'entrée avant application de la regex,
# garantissant un temps d'exécution constant.

import re

from django.http import JsonResponse

# Sécurisé : pas de groupes imbriqués, correspondance non ambiguë
EMAIL_RE = re.compile(r"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$")


def validate_email(request):
    email = request.GET.get("email", "")

    if len(email) > 254:
        return JsonResponse({"valid": False})

    is_valid = EMAIL_RE.match(email) is not None

    return JsonResponse({"valid": is_valid})
