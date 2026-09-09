"""
Exemple VULNERABLE — Django : client HTTP sortant forcé en SSLv2.

CWE-326: Inadequate Encryption Strength
Un service Django appelle une API partenaire legacy et force le client
`requests` à négocier en SSLv2 via un adaptateur HTTPS personnalisé.
SSLv2 est interdit par la RFC 6176 depuis 2011 : négociation non
authentifiée, MAC faible, et vulnérable à l'attaque DROWN qui permet de
déchiffrer le trafic TLS d'un serveur partageant la même clé RSA.
"""

import ssl

import requests
from requests.adapters import HTTPAdapter
from urllib3.poolmanager import PoolManager


class SSLv2Adapter(HTTPAdapter):
    """Adaptateur requests forçant le protocole SSLv2 pour un partenaire legacy."""

    def init_poolmanager(self, *args, **kwargs):
        # VULNERABLE : ssl.PROTOCOL_SSLv2 force la négociation en SSLv2,
        # un protocole cassé structurellement (pas d'authentification du
        # handshake, MAC faible, vulnérable à DROWN).
        context = ssl.SSLContext(ssl.PROTOCOL_SSLv2)
        kwargs["ssl_context"] = context
        self.poolmanager = PoolManager(*args, **kwargs)


class PartnerInvoiceClient:
    """Client HTTP pour synchroniser les factures avec un système partenaire legacy."""

    def __init__(self, base_url: str):
        self.base_url = base_url
        self.session = requests.Session()
        # VULNERABLE : toutes les requêtes HTTPS vers ce partenaire passent
        # par un canal chiffré en SSLv2, interceptable et déchiffrable.
        self.session.mount("https://", SSLv2Adapter())

    def fetch_invoices(self, partner_id: str):
        response = self.session.get(f"{self.base_url}/partners/{partner_id}/invoices")
        response.raise_for_status()
        return response.json()
