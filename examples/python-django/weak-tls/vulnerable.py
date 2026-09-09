"""
Exemple VULNERABLE — Django : appels sortants avec vérification TLS désactivée.

CWE-326: Inadequate Encryption Strength
Le service applicatif effectue des appels HTTPS vers une API de paiement
tierce en désactivant la vérification du certificat serveur et en
autorisant des suites de chiffrement obsolètes. Le canal est alors
vulnérable à une interception/falsification (MITM), même si le protocole
utilisé est nominalement TLS.
"""

import ssl

import requests
from django.conf import settings
from django.http import JsonResponse
from django.views import View

PAYMENT_API_URL = "https://payments.example-provider.com/v1/charge"


class ChargePaymentView(View):
    """Déclenche un paiement auprès du prestataire tiers."""

    def post(self, request):
        payload = {
            "amount": request.POST.get("amount"),
            "currency": "XOF",
            "customer_id": request.user.id,
        }

        # VULNERABLE : `verify=False` désactive intégralement la
        # vérification du certificat du serveur distant. N'importe quel
        # intermédiaire (proxy malveillant, réseau compromis) peut alors
        # intercepter ou falsifier la requête/réponse sans être détecté.
        response = requests.post(
            PAYMENT_API_URL,
            json=payload,
            headers={"Authorization": f"Bearer {settings.PAYMENT_API_KEY}"},
            verify=False,
            timeout=10,
        )

        return JsonResponse(response.json(), status=response.status_code)


def build_legacy_ssl_context():
    """
    VULNERABLE : contexte SSL personnalisé autorisant des suites de
    chiffrement faibles (RC4, DES, sans forward secrecy) et une version
    minimale de TLS trop permissive.
    """
    context = ssl.SSLContext(ssl.PROTOCOL_TLS_CLIENT)
    context.check_hostname = False
    context.verify_mode = ssl.CERT_NONE

    # VULNERABLE : suites de chiffrement obsolètes explicitement autorisées.
    context.set_ciphers("RC4:DES-CBC3-SHA:AES128-SHA")

    return context
