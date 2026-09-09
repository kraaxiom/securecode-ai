"""
Exemple VULNERABLE — Django : génération de tokens avec le module `random`.

CWE-338: Use of Cryptographically Weak Pseudo-Random Number Generator (PRNG)
La vue génère des tokens de réinitialisation de mot de passe, des clés d'API
et des codes OTP avec le module standard `random`, un PRNG statistique
prévisible et non conçu pour un usage sécuritaire. Un attaquant capable
d'observer une partie des sorties peut en déduire l'état interne et prédire
les valeurs futures ou passées.
"""

import random
import string

from django.http import JsonResponse
from django.views import View


class PasswordResetRequestView(View):
    """Génère un token de réinitialisation de mot de passe."""

    def post(self, request):
        email = request.POST.get("email")

        # VULNERABLE : `random` (module standard, PRNG Mersenne Twister)
        # est prévisible et ne doit jamais servir à générer un secret de
        # sécurité comme un token de réinitialisation.
        token = "".join(random.choices(string.ascii_letters + string.digits, k=20))

        save_reset_token(email, token)
        send_reset_email(email, token)

        return JsonResponse({"status": "sent"})


class ApiKeyCreateView(View):
    """Génère une nouvelle clé d'API pour l'utilisateur connecté."""

    def post(self, request):
        # VULNERABLE : même problème pour une clé d'API — `random.random()`
        # n'offre aucune garantie cryptographique d'imprévisibilité.
        api_key = str(random.random())[2:] + str(random.randint(100000, 999999))

        save_api_key(request.user, api_key)

        return JsonResponse({"api_key": api_key})


class OtpRequestView(View):
    """Génère un code OTP à 6 chiffres pour la vérification en deux étapes."""

    def post(self, request):
        # VULNERABLE : code OTP généré avec `random.randint`, prévisible si
        # la graine ou une partie de la séquence est connue de l'attaquant.
        otp_code = str(random.randint(0, 999999)).zfill(6)

        save_otp(request.user, otp_code)
        send_otp_sms(request.user, otp_code)

        return JsonResponse({"status": "sent"})


def save_reset_token(email, token):
    raise NotImplementedError


def send_reset_email(email, token):
    raise NotImplementedError


def save_api_key(user, api_key):
    raise NotImplementedError


def save_otp(user, otp_code):
    raise NotImplementedError


def send_otp_sms(user, otp_code):
    raise NotImplementedError
