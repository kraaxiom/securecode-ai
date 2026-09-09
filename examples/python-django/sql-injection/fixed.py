# Vue Django corrigée — CWE-89 (SQL Injection générique)
# Correction : requête préparée avec paramètre lié via l'API du curseur
# Django, qui neutralise toute tentative d'injection dans la valeur.

from django.db import connection
from django.http import JsonResponse


def search_users(request):
    name = request.GET.get("name", "")

    with connection.cursor() as cursor:
        # Sécurisé : paramètre lié, jamais inséré dans le texte de la requête
        cursor.execute("SELECT id, name FROM users WHERE name = %s", [name])
        rows = cursor.fetchall()

    return JsonResponse({"results": rows}, safe=False)
