"""
Exemple CORRIGE — Django : émission de JWT avec un secret fort et vérification stricte.

Fix pour CWE-1391 (Use of Weak Credentials) :
- Le secret provient d'une variable d'environnement / gestionnaire de secrets,
  avec au moins 256 bits d'entropie générés via un CSPRNG.
- La vérification épingle explicitement l'algorithme attendu (HS256 uniquement)
  et rejette tout token annonçant `alg: none` ou un algorithme non prévu.
- Une expiration courte est appliquée sur chaque token émis.
"""

import os

import jwt
from django.http import JsonResponse
from django.views import View

# CORRIGE : secret chargé depuis l'environnement (généré hors du code avec
# secrets.token_urlsafe(64) ou équivalent, stocké dans un gestionnaire de
# secrets — jamais commité). L'application échoue explicitement si absent.
JWT_SECRET = os.environ["JWT_SECRET"]
if len(JWT_SECRET.encode()) < 32:
    raise RuntimeError("JWT_SECRET doit contenir au moins 256 bits d'entropie")

ALLOWED_ALGORITHMS = ["HS256"]


class LoginView(View):
    """Authentifie l'utilisateur et émet un token JWT avec expiration courte."""

    def post(self, request):
        username = request.POST.get("username")
        password = request.POST.get("password")

        user = authenticate_user(username, password)
        if user is None:
            return JsonResponse({"error": "invalid credentials"}, status=401)

        import time

        # CORRIGE : secret fort + expiration courte (`exp`) intégrée au
        # payload pour limiter la fenêtre d'exploitation d'un token volé.
        payload = {
            "user_id": user.id,
            "username": user.username,
            "exp": int(time.time()) + 15 * 60,
        }
        token = jwt.encode(payload, JWT_SECRET, algorithm="HS256")

        return JsonResponse({"token": token})


class ProtectedView(View):
    """Vérifie le token JWT en épinglant strictement l'algorithme attendu."""

    def get(self, request):
        token = request.headers.get("Authorization", "").replace("Bearer ", "")

        try:
            # CORRIGE : `algorithms` restreint explicitement à HS256 — tout
            # token annonçant `alg: none` ou un autre algorithme est rejeté
            # par la bibliothèque avant même de vérifier la signature.
            payload = jwt.decode(token, JWT_SECRET, algorithms=ALLOWED_ALGORITHMS)
        except jwt.InvalidTokenError:
            return JsonResponse({"error": "invalid token"}, status=401)

        return JsonResponse({"user_id": payload["user_id"]})


def authenticate_user(username, password):
    """Stub d'authentification pour l'exemple."""
    raise NotImplementedError
