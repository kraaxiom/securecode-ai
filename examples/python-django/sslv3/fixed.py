"""
Exemple CORRIGE — Django : webhook sortant forcé en TLS 1.2 minimum.

Fix pour CWE-326 (Inadequate Encryption Strength) :
- Le contexte SSLv3 est remplacé par un `ssl.SSLContext` avec
  `minimum_version = ssl.TLSVersion.TLSv1_2` et suites de chiffrement
  authentifiées (AEAD), rendant l'attaque POODLE inapplicable.
- La vérification du certificat du webhook partenaire reste active.
"""

import ssl

import requests
from requests.adapters import HTTPAdapter
from urllib3.poolmanager import PoolManager


class ModernTLSAdapter(HTTPAdapter):
    """Adaptateur requests imposant TLS 1.2 minimum (TLS 1.3 si disponible)."""

    def init_poolmanager(self, *args, **kwargs):
        context = ssl.SSLContext(ssl.PROTOCOL_TLS_CLIENT)
        # CORRIGE : version minimale imposée à TLS 1.2 -> SSLv3 (et le
        # padding oracle CBC exploité par POODLE) n'est plus négociable.
        context.minimum_version = ssl.TLSVersion.TLSv1_2
        context.maximum_version = ssl.TLSVersion.TLSv1_3
        # CORRIGE : vérification du certificat et du nom d'hôte conservées
        # explicitement actives.
        context.check_hostname = True
        context.verify_mode = ssl.CERT_REQUIRED

        kwargs["ssl_context"] = context
        self.poolmanager = PoolManager(*args, **kwargs)


class PaymentWebhookNotifier:
    """Notifie un partenaire des événements de paiement via webhook HTTPS."""

    def __init__(self, webhook_url: str):
        self.webhook_url = webhook_url
        self.session = requests.Session()
        # CORRIGE : les notifications de paiement transitent désormais sur
        # un canal TLS 1.2+ avec suites de chiffrement authentifiées.
        self.session.mount("https://", ModernTLSAdapter())

    def notify_payment_success(self, payload: dict):
        response = self.session.post(self.webhook_url, json=payload, timeout=5)
        response.raise_for_status()
        return response.status_code
