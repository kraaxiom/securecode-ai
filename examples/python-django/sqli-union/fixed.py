# Vue Django corrigée — CWE-89 (UNION-based SQL Injection)
# Correction : casting strict de l'identifiant en entier et requête
# préparée avec paramètre lié, ce qui rend impossible l'ajout d'une
# clause UNION SELECT.

from django.db import connection
from django.http import HttpResponseBadRequest, JsonResponse


def product_detail(request):
    raw_id = request.GET.get("id", "")

    try:
        product_id = int(raw_id)
    except ValueError:
        return HttpResponseBadRequest("Identifiant invalide")

    with connection.cursor() as cursor:
        # Sécurisé : valeur typée + paramètre lié
        cursor.execute(
            "SELECT id, name, price FROM products WHERE id = %s", [product_id]
        )
        row = cursor.fetchone()

    return JsonResponse({"product": row}, safe=False)
