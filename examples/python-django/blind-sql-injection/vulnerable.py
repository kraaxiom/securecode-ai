# Faille : injection SQL aveugle (CWE-89).
# La valeur "username" est concaténée directement dans la requête SQL.
# Le résultat n'est pas affiché mais le booléen "exists" révèle une
# information exploitable (déduction caractère par caractère).
from django.http import JsonResponse
from django.db import connection


def check_user_exists(request):
    username = request.GET.get("username", "")

    with connection.cursor() as cursor:
        cursor.execute(
            "SELECT 1 FROM auth_user WHERE username = '%s' AND is_active = true"
            % username
        )
        exists = cursor.fetchone() is not None

    return JsonResponse({"exists": exists})
