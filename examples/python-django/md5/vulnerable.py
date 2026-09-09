"""
Exemple VULNERABLE — Django : hachage de mots de passe avec MD5.

CWE-328: Use of Weak Hash
L'application hache les mots de passe utilisateurs avec MD5 avant stockage
en base de données. MD5 est cassé depuis 2004 (collisions générables en
quelques secondes) et ne dispose d'aucun facteur de coût configurable : un
attaquant disposant de la base peut casser la quasi-totalité des mots de
passe par force brute/tables arc-en-ciel en un temps très court.
"""

import hashlib

from django.contrib.auth.models import AbstractUser
from django.http import JsonResponse
from django.views import View


class SignupView(View):
    """Inscrit un nouvel utilisateur en stockant son mot de passe haché."""

    def post(self, request):
        username = request.POST["username"]
        password = request.POST["password"]

        # VULNERABLE : MD5 ne comporte aucun sel ni facteur de coût. Deux
        # utilisateurs avec le même mot de passe produisent le même hash,
        # ce qui facilite les attaques par table arc-en-ciel précalculée.
        password_hash = hashlib.md5(password.encode("utf-8")).hexdigest()

        from myapp.models import User

        User.objects.create(username=username, password_hash=password_hash)
        return JsonResponse({"status": "created"})


class LoginView(View):
    """Authentifie un utilisateur en comparant les hachages MD5."""

    def post(self, request):
        username = request.POST["username"]
        password = request.POST["password"]

        from myapp.models import User

        user = User.objects.filter(username=username).first()
        if user is None:
            return JsonResponse({"error": "invalid_credentials"}, status=401)

        # VULNERABLE : comparaison directe de hachages MD5, cassable par
        # force brute GPU en quelques heures pour des mots de passe courants,
        # et vulnérable aux collisions.
        candidate_hash = hashlib.md5(password.encode("utf-8")).hexdigest()
        if candidate_hash != user.password_hash:
            return JsonResponse({"error": "invalid_credentials"}, status=401)

        return JsonResponse({"status": "authenticated"})
