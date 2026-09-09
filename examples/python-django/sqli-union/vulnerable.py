# Vue Django vulnérable — CWE-89 (UNION-based SQL Injection)
# L'identifiant "product_id" est concaténé directement dans la requête
# SQL sans typage ni liaison de paramètre, permettant à un attaquant
# d'ajouter une clause UNION SELECT pour extraire des données d'autres
# tables.

from django.db import connection
from django.http import JsonResponse


def product_detail(request):
    product_id = request.GET.get("id", "")

    with connection.cursor() as cursor:
        # Vulnérable : id non casté, concaténé tel quel dans le SQL
        cursor.execute(
            f"SELECT id, name, price FROM products WHERE id = {product_id}"
        )
        row = cursor.fetchone()

    return JsonResponse({"product": row}, safe=False)
