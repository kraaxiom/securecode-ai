"""
Exemple CORRIGE — Django : client HTTP sortant forcé en TLS 1.2 minimum.

Fix pour CWE-326 (Inadequate Encryption Strength) :
- L'adaptateur HTTPS personnalisé est remplacé par un `ssl.SSLContext`
  configurant explicitement `minimum_version = ssl.TLSVersion.TLSv1_2`,
  avec TLS 1.3 privilégié si le partenaire le supporte.
- La vérification du certificat du serveur reste activée (aucune
  désactivation de `verify`), pour empêcher tout interception MITM.
"""

import ssl

import requests
from requests.adapters import HTTPAdapter
from urllib3.poolmanager import PoolManager


class ModernTLSAdapter(HTTPAdapter):
    """Adaptateur requests imposant TLS 1.2 minimum (TLS 1.3 si disponible)."""

    def init_poolmanager(self, *args, **kwargs):
        context = ssl.SSLContext(ssl.PROTOCOL_TLS_CLIENT)
        # CORRIGE : version minimale imposée à TLS 1.2, SSLv2/SSLv3/TLS1.0/1.1
        # sont désormais impossibles à négocier avec ce contexte.
        context.minimum_version = ssl.TLSVersion.TLSv1_2
        context.maximum_version = ssl.TLSVersion.TLSv1_3
        # CORRIGE : vérification du certificat et du nom d'hôte conservées
        # explicitement actives (comportement par défaut de PROTOCOL_TLS_CLIENT).
        context.check_hostname = True
        context.verify_mode = ssl.CERT_REQUIRED

        kwargs["ssl_context"] = context
        self.poolmanager = PoolManager(*args, **kwargs)


class PartnerInvoiceClient:
    """Client HTTP pour synchroniser les factures avec un système partenaire legacy."""

    def __init__(self, base_url: str):
        self.base_url = base_url
        self.session = requests.Session()
        # CORRIGE : toutes les requêtes HTTPS vers ce partenaire négocient
        # désormais au minimum du TLS 1.2, avec vérification du certificat.
        self.session.mount("https://", ModernTLSAdapter())

    def fetch_invoices(self, partner_id: str):
        response = self.session.get(f"{self.base_url}/partners/{partner_id}/invoices")
        response.raise_for_status()
        return response.json()
