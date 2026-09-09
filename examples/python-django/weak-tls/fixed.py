"""
Exemple CORRIGE — Django : appels sortants avec vérification TLS active et suites modernes.

Fix pour CWE-326 (Inadequate Encryption Strength) :
- La vérification du certificat serveur reste activée (`verify=True`,
  comportement par défaut de `requests`), y compris en environnement de test.
- Le contexte SSL personnalisé restreint les suites de chiffrement aux
  algorithmes AEAD modernes avec forward secrecy et impose TLS 1.2 minimum.
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

        # CORRIGE : vérification du certificat activée (comportement par
        # défaut de `requests` — `verify=True` explicite ici pour la clarté).
        # Utilise le magasin de CA système ; possibilité de fournir un
        # bundle explicite via `verify="/path/to/ca-bundle.pem"` si requis.
        response = requests.post(
            PAYMENT_API_URL,
            json=payload,
            headers={"Authorization": f"Bearer {settings.PAYMENT_API_KEY}"},
            verify=True,
            timeout=10,
        )

        return JsonResponse(response.json(), status=response.status_code)


def build_modern_ssl_context():
    """
    CORRIGE : contexte SSL restreint aux suites de chiffrement modernes
    (AEAD, forward secrecy) avec TLS 1.2 comme version minimale et
    vérification de certificat/hostname activée.
    """
    context = ssl.SSLContext(ssl.PROTOCOL_TLS_CLIENT)
    context.check_hostname = True
    context.verify_mode = ssl.CERT_REQUIRED
    context.minimum_version = ssl.TLSVersion.TLSv1_2

    # CORRIGE : suites de chiffrement limitées aux algorithmes AEAD modernes
    # avec échange de clé éphémère (forward secrecy).
    context.set_ciphers("ECDHE+AESGCM:ECDHE+CHACHA20")

    return context
