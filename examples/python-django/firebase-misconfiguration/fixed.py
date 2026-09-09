"""
Exemple CORRIGE — Django + Firebase Admin SDK : accès contrôlé.

Fix pour CWE-284 (Improper Access Control) :
- La clé de service Firebase est chargée depuis un gestionnaire de secrets
  (jamais codée en dur ni versionnée).
- Le backend applique un contrôle d'autorisation explicite avant toute
  lecture/écriture.
- Les règles Firestore associées (firestore.rules, déployées séparément)
  doivent être restreintes à l'auth + la propriété du document — voir
  extrait de référence en commentaire ci-dessous.
"""

import json
import os

import firebase_admin
from firebase_admin import credentials, firestore
from django.http import JsonResponse, HttpResponseForbidden
from django.views import View

# CORRIGE : clé de service chargée depuis un gestionnaire de secrets
# (ex: AWS Secrets Manager) au démarrage, injectée via variable
# d'environnement — jamais codée en dur ni committée.
FIREBASE_SERVICE_ACCOUNT = json.loads(os.environ["FIREBASE_SERVICE_ACCOUNT_JSON"])

cred = credentials.Certificate(FIREBASE_SERVICE_ACCOUNT)
firebase_admin.initialize_app(cred)
db = firestore.client()

# CORRIGE : règles Firestore de référence à déployer (firestore.rules) :
#
# rules_version = '2';
# service cloud.firestore {
#   match /databases/{database}/documents {
#     match /users/{userId}/documents/{docId} {
#       allow read, write: if request.auth != null
#                           && request.auth.uid == userId;
#     }
#   }
# }
#
# -> plus aucun accès anonyme ni accès croisé entre utilisateurs.


class UserProfileView(View):
    """Expose le profil utilisateur stocké dans Firestore, avec contrôle d'accès."""

    def get(self, request, user_id):
        # CORRIGE : contrôle d'autorisation explicite côté backend, en
        # complément des règles Firestore restrictives.
        if not request.user.is_authenticated or str(request.user.id) != str(user_id):
            return HttpResponseForbidden("Accès refusé.")

        doc = db.collection("users").document(user_id).get()
        return JsonResponse(doc.to_dict() or {})
