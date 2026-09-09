# Faille : injection SQL basée sur les erreurs (CWE-89).
# L'identifiant de commande est concaténé dans la requête SQL, et le
# message d'exception brut du driver est renvoyé au client en cas d'échec,
# ce qui permet d'exfiltrer des données via des erreurs SQL provoquées.
from django.http import JsonResponse
from django.db import connection, DatabaseError


def get_order(request):
    order_id = request.GET.get("id")

    try:
        with connection.cursor() as cursor:
            cursor.execute(f"SELECT * FROM shop_order WHERE id = {order_id}")
            row = cursor.fetchone()
        return JsonResponse({"order": row})
    except DatabaseError as e:
        return JsonResponse({"error": str(e)}, status=500)
