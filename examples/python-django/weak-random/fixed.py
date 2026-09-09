"""
Exemple CORRIGE — Django : génération de tokens avec le module `secrets` (CSPRNG).

Fix pour CWE-338 (Use of Cryptographically Weak Pseudo-Random Number Generator) :
- Tous les tokens, clés d'API et codes OTP sont générés avec le module
  `secrets`, un CSPRNG adossé à la source d'aléa du système d'exploitation.
- Chaque valeur générée offre une entropie d'au moins 128 bits (sauf l'OTP,
  dont la faible longueur est volontaire et compensée par une limitation
  de tentatives et une expiration courte).
"""

import secrets

from django.http import JsonResponse
from django.views import View


class PasswordResetRequestView(View):
    """Génère un token de réinitialisation de mot de passe."""

    def post(self, request):
        email = request.POST.get("email")

        # CORRIGE : `secrets.token_urlsafe` s'appuie sur `os.urandom`,
        # imprévisible même en connaissant des sorties précédentes.
        # 32 octets -> ~256 bits d'entropie avant encodage.
        token = secrets.token_urlsafe(32)

        save_reset_token(email, token)
        send_reset_email(email, token)

        return JsonResponse({"status": "sent"})


class ApiKeyCreateView(View):
    """Génère une nouvelle clé d'API pour l'utilisateur connecté."""

    def post(self, request):
        # CORRIGE : clé d'API générée avec un CSPRNG, entropie suffisante
        # (32 octets hexadécimaux = 256 bits).
        api_key = secrets.token_hex(32)

        save_api_key(request.user, api_key)

        return JsonResponse({"api_key": api_key})


class OtpRequestView(View):
    """Génère un code OTP à 6 chiffres pour la vérification en deux étapes."""

    def post(self, request):
        # CORRIGE : `secrets.choice` tire chaque chiffre depuis un CSPRNG.
        # La faible longueur du code OTP (usage utilisateur) est compensée
        # par une expiration courte et une limitation du nombre de tentatives
        # côté serveur (non représentées ici pour rester concis).
        otp_code = "".join(secrets.choice("0123456789") for _ in range(6))

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
