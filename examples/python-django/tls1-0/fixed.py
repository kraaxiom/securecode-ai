"""
Exemple CORRIGE — Django : client de paiement forcé en TLS 1.2 minimum.

Fix pour CWE-326 (Inadequate Encryption Strength) :
- Le plafond TLS 1.0 est remplacé par `minimum_version =
  ssl.TLSVersion.TLSv1_2`, avec TLS 1.3 privilégié, conformément à la
  RFC 8996 et aux exigences PCI-DSS pour les flux de paiement.
"""

import ssl

import requests
from requests.adapters import HTTPAdapter
from urllib3.poolmanager import PoolManager


class ModernPaymentGatewayAdapter(HTTPAdapter):
    """Adaptateur requests imposant TLS 1.2 minimum (TLS 1.3 si disponible)."""

    def init_poolmanager(self, *args, **kwargs):
        context = ssl.SSLContext(ssl.PROTOCOL_TLS_CLIENT)
        # CORRIGE : minimum relevé à TLS 1.2, conforme PCI-DSS et à la
        # dépréciation de TLS 1.0/1.1 par la RFC 8996.
        context.minimum_version = ssl.TLSVersion.TLSv1_2
        context.maximum_version = ssl.TLSVersion.TLSv1_3
        # CORRIGE : vérification du certificat et du nom d'hôte conservées
        # explicitement actives.
        context.check_hostname = True
        context.verify_mode = ssl.CERT_REQUIRED

        kwargs["ssl_context"] = context
        self.poolmanager = PoolManager(*args, **kwargs)


class PaymentGatewayClient:
    """Client HTTP vers la passerelle de paiement partenaire."""

    def __init__(self, base_url: str):
        self.base_url = base_url
        self.session = requests.Session()
        # CORRIGE : les données de paiement transitent désormais sur un
        # canal TLS 1.2+ avec suites de chiffrement authentifiées.
        self.session.mount("https://", ModernPaymentGatewayAdapter())

    def charge(self, token: str, amount_cents: int):
        response = self.session.post(
            f"{self.base_url}/charges",
            json={"token": token, "amount": amount_cents},
            timeout=10,
        )
        response.raise_for_status()
        return response.json()
