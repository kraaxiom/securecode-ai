"""
Exemple VULNERABLE — Django : hachage de mot de passe et vérification
d'intégrité avec SHA-1.

CWE-328: Use of Weak Hash
SHA-1 est cryptographiquement cassé pour la résistance aux collisions
(attaque "SHAttered", 2017). L'utiliser pour stocker des mots de passe ou
pour garantir l'intégrité d'un fichier sensible permet à un attaquant de
forger des collisions ou de casser les hachages par force brute massive
(SHA-1 est rapide, donc peu coûteux à attaquer).
"""

import hashlib

from django.conf import settings
from django.contrib.auth.models import AbstractUser
from django.db import models
from django.http import JsonResponse
from django.views import View


class LegacyUser(AbstractUser):
    """Modèle utilisateur avec un stockage de mot de passe fait maison."""

    # VULNERABLE : le mot de passe n'est jamais stocké via le système de
    # hachage de Django (PASSWORD_HASHERS) mais via un SHA-1 fait maison.
    password_hash = models.CharField(max_length=40)


class RegisterView(View):
    """Inscription d'un nouvel utilisateur."""

    def post(self, request):
        username = request.POST["username"]
        password = request.POST["password"]

        # VULNERABLE : SHA-1 est rapide et non salé ici -> vulnérable aux
        # attaques par rainbow table et par force brute massive (GPU).
        password_hash = hashlib.sha1(password.encode()).hexdigest()

        LegacyUser.objects.create(username=username, password_hash=password_hash)
        return JsonResponse({"status": "created"})


class LoginView(View):
    """Vérifie les identifiants en recalculant le SHA-1 du mot de passe fourni."""

    def post(self, request):
        username = request.POST["username"]
        password = request.POST["password"]

        # VULNERABLE : comparaison directe de hachages SHA-1, sans sel ni
        # facteur de coût -> aucune protection contre le brute force hors
        # ligne en cas de fuite de la base de données.
        candidate_hash = hashlib.sha1(password.encode()).hexdigest()
        user = LegacyUser.objects.filter(
            username=username, password_hash=candidate_hash
        ).first()

        if user is None:
            return JsonResponse({"status": "invalid"}, status=401)
        return JsonResponse({"status": "ok"})


def sign_export_file(file_content: bytes) -> str:
    """
    VULNERABLE : utilisation de SHA-1 pour garantir l'intégrité d'un export
    de données sensibles (ex: export RGPD envoyé à un utilisateur). Un
    attaquant capable de générer une collision SHA-1 peut produire un
    fichier différent partageant la même empreinte.
    """
    secret = settings.SECRET_KEY.encode()
    # VULNERABLE : concaténation naïve secret+contenu au lieu d'un HMAC,
    # combinée à un algorithme de hachage affaibli.
    return hashlib.sha1(secret + file_content).hexdigest()
