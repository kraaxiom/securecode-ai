# Vue Django vulnérable — CWE-1333 (ReDoS)
# La validation d'email utilise une regex avec des quantificateurs
# imbriqués ("(...)+" contenant lui-même un "+") appliquée à une entrée
# non bornée en taille. Une chaîne pathologique peut provoquer un temps
# de calcul exponentiel et bloquer le worker.

import re

from django.http import JsonResponse

# Vulnérable : groupes répétés imbriqués, ambiguïté de correspondance
EMAIL_RE = re.compile(r"^([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)+$")


def validate_email(request):
    email = request.GET.get("email", "")

    is_valid = EMAIL_RE.match(email) is not None

    return JsonResponse({"valid": is_valid})
