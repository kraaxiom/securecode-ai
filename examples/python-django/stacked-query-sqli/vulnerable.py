# Vue Django vulnérable — CWE-89 (Stacked Query SQL Injection)
# Le nom fourni par l'utilisateur est concaténé dans une requête SQL
# exécutée via un curseur brut. Selon le driver/backend utilisé
# (multi-statement activé), un attaquant peut ajouter un point-virgule
# suivi d'une instruction SQL distincte (UPDATE, INSERT, DROP...).

from django.db import connection
from django.http import HttpResponse


def update_username(request):
    new_name = request.POST.get("name", "")
    user_id = request.POST.get("user_id", "1")

    with connection.cursor() as cursor:
        # Vulnérable : concaténation directe, exécution multi-instructions possible
        cursor.execute(
            f"UPDATE users SET name = '{new_name}' WHERE id = {user_id}"
        )

    return HttpResponse("Utilisateur mis à jour")
