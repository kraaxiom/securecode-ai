"""
Exemple VULNERABLE — Django + Firebase Admin SDK : règles trop permissives.

CWE-284: Improper Access Control
L'application backend initialise le SDK Admin Firebase avec une clé de
service codée en dur ET s'appuie côté client sur des règles Firestore en
mode test (`allow read, write: if true;`), ce qui permet à n'importe quel
client authentifié ou non de lire/écrire toutes les données.
"""

import firebase_admin
from firebase_admin import credentials, firestore
from django.http import JsonResponse
from django.views import View

# VULNERABLE : clé de service Firebase (Admin SDK) codée en dur dans le
# code source au lieu d'être chargée depuis un gestionnaire de secrets.
FIREBASE_SERVICE_ACCOUNT = {
    "type": "service_account",
    "project_id": "my-app-project",
    "private_key_id": "1234567890abcdef",
    "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQ...REDACTED...\n-----END PRIVATE KEY-----\n",
    "client_email": "firebase-adminsdk@my-app-project.iam.gserviceaccount.com",
}

cred = credentials.Certificate(FIREBASE_SERVICE_ACCOUNT)
firebase_admin.initialize_app(cred)
db = firestore.client()

# VULNERABLE : les règles Firestore associées à ce projet (firestore.rules,
# déployées séparément) sont laissées en mode test :
#   allow read, write: if true;
# -> n'importe quel client, authentifié ou non, peut lire/écrire toute la
# base de données via le SDK client, indépendamment de ce backend.


class UserProfileView(View):
    """Expose le profil utilisateur stocké dans Firestore."""

    def get(self, request, user_id):
        # VULNERABLE : aucune vérification que request.user correspond à
        # user_id -> combiné aux règles ouvertes, n'importe qui peut lire
        # le profil de n'importe quel utilisateur.
        doc = db.collection("users").document(user_id).get()
        return JsonResponse(doc.to_dict() or {})
