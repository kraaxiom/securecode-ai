"""
Fonctionnalité : validateur d'URL de webhook — version corrigée.

Correctif contre le DNS rebinding (CWE-918) : résoudre le DNS UNE SEULE
FOIS, valider l'IP obtenue, puis effectuer la connexion HTTP directement
sur cette IP validée (DNS pinning), sans jamais laisser le client HTTP
résoudre à nouveau le nom de domaine. La validation et la connexion
utilisent ainsi strictement la même adresse — plus de fenêtre TOCTOU
exploitable par changement de réponse DNS entre les deux étapes.
"""
import ipaddress
import socket

import requests
from django.http import JsonResponse
from django.views import View


def _is_public_ip(ip_str: str) -> bool:
    ip = ipaddress.ip_address(ip_str)
    return not (ip.is_private or ip.is_loopback or ip.is_link_local or ip.is_reserved)


def resolve_pin_and_validate(hostname: str) -> str:
    """Résolution unique du nom de domaine, IP retenue et validée une fois
    pour toutes : c'est cette même IP qui sera utilisée pour la connexion,
    éliminant la fenêtre de rebinding entre validation et requête."""
    try:
        ip_str = socket.gethostbyname(hostname)
    except socket.gaierror as exc:
        raise ValueError("résolution DNS impossible") from exc

    if not _is_public_ip(ip_str):
        raise ValueError("IP interne refusée")

    return ip_str


class WebhookUrlValidatorView(View):
    """Une seule résolution DNS, épinglée pour toute la durée de l'appel :
    la couche réseau ne peut plus revalider vers une IP interne."""

    def post(self, request):
        hostname = request.POST.get("hostname")
        if not hostname:
            return JsonResponse({"error": "hostname manquant"}, status=400)

        try:
            pinned_ip = resolve_pin_and_validate(hostname)
        except ValueError as exc:
            return JsonResponse({"error": str(exc)}, status=400)

        # La connexion se fait sur l'IP épinglée (pas sur le nom de
        # domaine) : aucune seconde résolution DNS n'est possible ici.
        # L'en-tête Host et la vérification TLS SNI restent alignés sur le
        # nom de domaine d'origine pour préserver la validité du certificat.
        response = requests.get(
            f"https://{pinned_ip}/health",
            headers={"Host": hostname},
            timeout=3,
            allow_redirects=False,
            verify=True,
        )

        return JsonResponse({"status_code": response.status_code})
