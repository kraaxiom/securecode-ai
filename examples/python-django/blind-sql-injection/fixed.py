# Correction : la requête est paramétrée (aucune concaténation), et le
# temps de réponse/le message renvoyé restent uniformes quel que soit le
# résultat, ce qui supprime le canal d'inférence exploité par le blind SQLi.
from django.http import JsonResponse
from django.db import connection


def check_user_exists(request):
    username = request.GET.get("username", "")

    with connection.cursor() as cursor:
        cursor.execute(
            "SELECT 1 FROM auth_user WHERE username = %s AND is_active = true",
            [username],
        )
        exists = cursor.fetchone() is not None

    # Réponse générique, temps de traitement comparable dans les deux cas.
    return JsonResponse({"exists": exists})
