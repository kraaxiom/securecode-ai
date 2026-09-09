"""
Exemple VULNERABLE — Django : webhook sortant configuré en SSLv3.

CWE-326: Inadequate Encryption Strength
Un service Django envoie des notifications de paiement à un webhook
partenaire en forçant SSLv3 pour "compatibilité". SSLv3 est interdit par
la RFC 7568 depuis 2015 : il est vulnérable à l'attaque POODLE, qui
permet à un attaquant en position de man-in-the-middle de déchiffrer des
données via un padding oracle sur le mode CBC.
"""

import ssl

import requests
from requests.adapters import HTTPAdapter
from urllib3.poolmanager import PoolManager


class LegacyWebhookAdapter(HTTPAdapter):
    """Adaptateur requests forçant SSLv3 pour un webhook partenaire ancien."""

    def init_poolmanager(self, *args, **kwargs):
        # VULNERABLE : ssl.PROTOCOL_SSLv3 force la négociation en SSLv3,
        # vulnérable à POODLE (padding oracle sur CBC) et interdit par la
        # RFC 7568.
        context = ssl.SSLContext(ssl.PROTOCOL_SSLv3)
        kwargs["ssl_context"] = context
        self.poolmanager = PoolManager(*args, **kwargs)


class PaymentWebhookNotifier:
    """Notifie un partenaire des événements de paiement via webhook HTTPS."""

    def __init__(self, webhook_url: str):
        self.webhook_url = webhook_url
        self.session = requests.Session()
        # VULNERABLE : les notifications de paiement (données financières)
        # transitent sur un canal chiffré cassé, exposé au déchiffrement
        # via POODLE par un attaquant réseau.
        self.session.mount("https://", LegacyWebhookAdapter())

    def notify_payment_success(self, payload: dict):
        response = self.session.post(self.webhook_url, json=payload, timeout=5)
        response.raise_for_status()
        return response.status_code
