# Faille : injection NoSQL (CWE-943).
# Le corps JSON de la requête est transmis tel quel comme filtre de
# requête MongoDB. Un attaquant peut envoyer un objet opérateur (ex:
# {"$ne": null}) à la place d'une chaîne pour contourner l'authentification.
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

    user = collection.find_one({"username": username, "password": password})

    return JsonResponse({"authenticated": user is not None})
