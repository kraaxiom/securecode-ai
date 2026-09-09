"""
Exemple VULNERABLE — Django : client de paiement configuré avec un
plafond TLS 1.0.

CWE-326: Inadequate Encryption Strength
Un service Django appelle une passerelle de paiement en limitant
explicitement la version TLS à 1.0. TLS 1.0/1.1 sont officiellement
dépréciés par la RFC 8996 (2021), vulnérables à l'attaque BEAST, ne
supportent pas les suites de chiffrement authentifiées modernes, et sont
exclus des référentiels de conformité comme PCI-DSS depuis 2018.
"""

import ssl

import requests
from requests.adapters import HTTPAdapter
from urllib3.poolmanager import PoolManager


class LegacyPaymentGatewayAdapter(HTTPAdapter):
    """Adaptateur requests limitant la négociation TLS à la version 1.0."""

    def init_poolmanager(self, *args, **kwargs):
        context = ssl.SSLContext(ssl.PROTOCOL_TLS_CLIENT)
        # VULNERABLE : minimum ET maximum figés sur TLS 1.0, alors que ce
        # protocole est déprécié (RFC 8996), vulnérable à BEAST et non
        # conforme PCI-DSS pour un flux de paiement.
        context.minimum_version = ssl.TLSVersion.TLSv1
        context.maximum_version = ssl.TLSVersion.TLSv1

        kwargs["ssl_context"] = context
        self.poolmanager = PoolManager(*args, **kwargs)


class PaymentGatewayClient:
    """Client HTTP vers la passerelle de paiement partenaire."""

    def __init__(self, base_url: str):
        self.base_url = base_url
        self.session = requests.Session()
        # VULNERABLE : les données de carte tokenisées et les confirmations
        # de transaction transitent via un canal TLS 1.0 obsolète.
        self.session.mount("https://", LegacyPaymentGatewayAdapter())

    def charge(self, token: str, amount_cents: int):
        response = self.session.post(
            f"{self.base_url}/charges",
            json={"token": token, "amount": amount_cents},
            timeout=10,
        )
        response.raise_for_status()
        return response.json()
