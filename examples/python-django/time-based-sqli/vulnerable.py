# Vue Django vulnérable — CWE-89 (Time-Based Blind SQL Injection)
# L'identifiant de commande est concaténé dans la requête SQL. Comme
# l'application ne renvoie ni données détaillées ni erreur (réponse
# générique), un attaquant peut exploiter une injection aveugle en
# observant le délai de réponse (ex: fonction de pause conditionnelle).

from django.db import connection
from django.http import JsonResponse


def order_status(request):
    order_id = request.GET.get("id", "")

    with connection.cursor() as cursor:
        # Vulnérable : concaténation, aucune limite de temps configurée
        cursor.execute(f"SELECT status FROM orders WHERE id = {order_id}")
        row = cursor.fetchone()

    return JsonResponse({"status": "ok" if row else "not_found"})
