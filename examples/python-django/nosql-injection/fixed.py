# Correction : chaque champ attendu est validé par type (str strict)
# avant d'être utilisé dans le filtre MongoDB, ce qui rejette tout objet
# ou opérateur (`$ne`, `$gt`) injecté à la place d'une valeur scalaire.
import json
from django.http import JsonResponse, HttpResponseBadRequest
from django.views.decorators.csrf import csrf_exempt
from pymongo import MongoClient

client = MongoClient("mongodb://localhost:27017")
collection = client.app_db.users


@csrf_exempt
def login(request):
    try:
        body = json.loads(request.body)
    except ValueError:
        return HttpResponseBadRequest("JSON invalide")

    username = body.get("username")
    password = body.get("password")

    if not isinstance(username, str) or not isinstance(password, str):
        return HttpResponseBadRequest("Format invalide")

    user = collection.find_one({"username": username, "password": password})

    return JsonResponse({"authenticated": user is not None})
