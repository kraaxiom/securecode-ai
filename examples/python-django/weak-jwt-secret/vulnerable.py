"""
Exemple VULNERABLE — Django : émission de JWT avec un secret faible.

CWE-1391: Use of Weak Credentials
La vue signe des tokens JWT (HS256) avec un secret court, codé en dur dans
le code source, et la vérification côté serveur accepte n'importe quel
algorithme annoncé par le token (y compris `alg: none`). Un attaquant peut
forger des tokens valides par force brute hors ligne ou en exploitant la
confusion d'algorithme.
"""

import jwt
from django.http import JsonResponse
from django.views import View

# VULNERABLE : secret codé en dur, court (moins de 32 octets) et présent
# dans le code source versionné — trivialement récupérable et sujet au
# bruteforce hors ligne.
JWT_SECRET = "changeme123"


class LoginView(View):
    """Authentifie l'utilisateur et émet un token JWT."""

    def post(self, request):
        username = request.POST.get("username")
        password = request.POST.get("password")

        user = authenticate_user(username, password)
        if user is None:
            return JsonResponse({"error": "invalid credentials"}, status=401)

        # VULNERABLE : secret faible, aucune expiration courte définie.
        payload = {"user_id": user.id, "username": user.username}
        token = jwt.encode(payload, JWT_SECRET, algorithm="HS256")

        return JsonResponse({"token": token})


class ProtectedView(View):
    """Vérifie le token JWT fourni par le client."""

    def get(self, request):
        token = request.headers.get("Authorization", "").replace("Bearer ", "")

        try:
            # VULNERABLE : aucun paramètre `algorithms` explicite fourni à
            # decode() dans certaines versions/wrappers, ou liste incluant
            # "none" -> un attaquant peut envoyer un token non signé
            # (alg: none) et le faire accepter comme valide.
            payload = jwt.decode(
                token,
                JWT_SECRET,
                algorithms=["HS256", "none"],
                options={"verify_signature": True},
            )
        except jwt.InvalidTokenError:
            return JsonResponse({"error": "invalid token"}, status=401)

        return JsonResponse({"user_id": payload["user_id"]})


def authenticate_user(username, password):
    """Stub d'authentification pour l'exemple."""
    raise NotImplementedError
