# Vue Django vulnérable — CWE-89 (SQL Injection générique)
# Le paramètre "name" est inséré par f-string directement dans la
# requête SQL exécutée via un curseur brut, sans liaison de paramètre.
# Un attaquant peut altérer la logique de la requête.

from django.db import connection
from django.http import JsonResponse


def search_users(request):
    name = request.GET.get("name", "")

    with connection.cursor() as cursor:
        # Vulnérable : concaténation de la valeur utilisateur dans le SQL
        cursor.execute(f"SELECT id, name FROM users WHERE name = '{name}'")
        rows = cursor.fetchall()

    return JsonResponse({"results": rows}, safe=False)
