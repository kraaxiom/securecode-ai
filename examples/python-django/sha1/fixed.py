"""
Exemple CORRIGE — Django : hachage de mot de passe avec Argon2id et
intégrité de fichier avec HMAC-SHA256.

Fix pour CWE-328 (Use of Weak Hash) :
- Les mots de passe utilisent le système de hachage natif de Django avec
  Argon2id (lent, salé, à facteur de coût réglable), au lieu d'un SHA-1
  fait maison.
- L'intégrité des exports sensibles est garantie par un HMAC-SHA256, qui
  n'est pas cassable par attaque de collision.
"""

import hashlib
import hmac

from django.conf import settings
from django.contrib.auth import authenticate
from django.contrib.auth.models import AbstractUser
from django.http import JsonResponse
from django.views import View

# CORRIGE : dans settings.py, s'assurer que Argon2PasswordHasher est en
# tête de PASSWORD_HASHERS, ex:
# PASSWORD_HASHERS = [
#     "django.contrib.auth.hashers.Argon2PasswordHasher",
#     "django.contrib.auth.hashers.PBKDF2PasswordHasher",
# ]


class LegacyUser(AbstractUser):
    """Modèle utilisateur : le mot de passe est géré par Django (champ `password` hérité)."""


class RegisterView(View):
    """Inscription d'un nouvel utilisateur."""

    def post(self, request):
        username = request.POST["username"]
        password = request.POST["password"]

        # CORRIGE : set_password() délègue à Argon2id (salé, coûteux en
        # calcul, résistant au brute force et aux rainbow tables).
        user = LegacyUser(username=username)
        user.set_password(password)
        user.save()
        return JsonResponse({"status": "created"})


class LoginView(View):
    """Vérifie les identifiants via le mécanisme d'authentification Django."""

    def post(self, request):
        username = request.POST["username"]
        password = request.POST["password"]

        # CORRIGE : authenticate() utilise check_password(), qui applique
        # une comparaison à temps constant sur le hachage Argon2id stocké.
        user = authenticate(request, username=username, password=password)

        if user is None:
            return JsonResponse({"status": "invalid"}, status=401)
        return JsonResponse({"status": "ok"})


def sign_export_file(file_content: bytes) -> str:
    """
    CORRIGE : intégrité garantie par un HMAC-SHA256 construit correctement
    (clé secrète via hmac.new, pas de concaténation naïve), résistant aux
    attaques par collision qui touchent SHA-1.
    """
    secret = settings.SECRET_KEY.encode()
    return hmac.new(secret, file_content, hashlib.sha256).hexdigest()
