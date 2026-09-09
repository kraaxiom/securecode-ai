# Vue Django corrigée — CWE-89 (Time-Based Blind SQL Injection)
# Correction : requête préparée avec paramètre lié (élimine la classe
# de vulnérabilité) et typage strict de l'identifiant. Un timeout
# d'exécution est en complément configuré côté base de données
# (settings.py : OPTIONS={'options': '-c statement_timeout=5000'}).

from django.db import connection
from django.http import HttpResponseBadRequest, JsonResponse


def order_status(request):
    raw_id = request.GET.get("id", "")

    try:
        order_id = int(raw_id)
    except ValueError:
        return HttpResponseBadRequest("Identifiant invalide")

    with connection.cursor() as cursor:
        # Sécurisé : paramètre lié, timeout appliqué au niveau de la connexion
        cursor.execute("SELECT status FROM orders WHERE id = %s", [order_id])
        row = cursor.fetchone()

    return JsonResponse({"status": "ok" if row else "not_found"})
