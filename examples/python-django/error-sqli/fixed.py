# Correction : requête paramétrée avec identifiant casté en entier, et
# gestion d'exception centralisée qui journalise le détail côté serveur
# sans jamais renvoyer le message natif du driver SQL au client.
import logging
from django.http import JsonResponse
from django.db import connection, DatabaseError

logger = logging.getLogger(__name__)


def get_order(request):
    try:
        order_id = int(request.GET.get("id"))
    except (TypeError, ValueError):
        return JsonResponse({"error": "Identifiant invalide"}, status=400)

    try:
        with connection.cursor() as cursor:
            cursor.execute(
                "SELECT * FROM shop_order WHERE id = %s", [order_id]
            )
            row = cursor.fetchone()
        return JsonResponse({"order": row})
    except DatabaseError as e:
        logger.error("Erreur SQL sur get_order: %s", e)
        return JsonResponse({"error": "Une erreur est survenue."}, status=500)
