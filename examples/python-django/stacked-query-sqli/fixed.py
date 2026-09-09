# Vue Django corrigée — CWE-89 (Stacked Query SQL Injection)
# Correction : requête préparée avec paramètres liés pour toutes les
# valeurs, et typage strict de l'identifiant. Le paramètre lié empêche
# toute injection de point-virgule suivi d'une instruction distincte.

from django.db import connection
from django.http import HttpResponse, HttpResponseBadRequest


def update_username(request):
    new_name = request.POST.get("name", "")
    raw_user_id = request.POST.get("user_id", "1")

    try:
        user_id = int(raw_user_id)
    except ValueError:
        return HttpResponseBadRequest("Identifiant invalide")

    with connection.cursor() as cursor:
        # Sécurisé : paramètres liés, aucune concaténation
        cursor.execute(
            "UPDATE users SET name = %s WHERE id = %s", [new_name, user_id]
        )

    return HttpResponse("Utilisateur mis à jour")
