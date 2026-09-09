"""
Fonctionnalité : validation asynchrone d'URL de webhook — version corrigée.

Les contrôles anti-SSRF sont identiques à une SSRF classique (whitelist,
résolution/validation d'IP, pinning, blocage des redirections) : l'absence
de retour visible au client ne dispense en rien de ces contrôles, elle
rend simplement la détection côté attaquant plus difficile — pas la faille
moins dangereuse. On ajoute ici une journalisation structurée des requêtes
sortantes, indispensable puisque le résultat n'est jamais exposé au client
(CWE-918).
"""
import ipaddress
import logging
import socket
from urllib.parse import urlparse

import requests
from django.http import JsonResponse
from django.views import View

logger = logging.getLogger("outbound_requests")

ALLOWED_SCHEMES = {"http", "https"}
ALLOWED_DOMAINS = {"example.com", "hooks.example.com"}


def _is_public_ip(ip_str: str) -> bool:
    ip = ipaddress.ip_address(ip_str)
    return not (
        ip.is_private
        or ip.is_loopback
        or ip.is_link_local
        or ip.is_multicast
        or ip.is_reserved
        or ip.is_unspecified
    )


def resolve_and_validate(hostname: str) -> str:
    """Résolution unique + validation de l'IP, pour épingler la connexion
    et empêcher un contournement par DNS rebinding."""
    try:
        infos = socket.getaddrinfo(hostname, None)
    except socket.gaierror as exc:
        raise ValueError(f"Résolution DNS impossible pour {hostname}") from exc

    for _, _, _, _, sockaddr in infos:
        if _is_public_ip(sockaddr[0]):
            return sockaddr[0]

    raise ValueError(f"Aucune IP publique valide pour {hostname}")


def register_webhook_task(webhook_url: str) -> None:
    """Tâche de fond durcie : mêmes contrôles qu'une SSRF classique, plus
    une journalisation systématique car le résultat n'est jamais renvoyé
    au client — c'est le seul canal de détection disponible."""
    parsed = urlparse(webhook_url)

    if parsed.scheme not in ALLOWED_SCHEMES or parsed.hostname not in ALLOWED_DOMAINS:
        logger.warning("webhook rejeté (domaine/schéma non autorisé): %s", webhook_url)
        return

    try:
        resolved_ip = resolve_and_validate(parsed.hostname)
    except ValueError:
        logger.warning("webhook rejeté (IP interne/invalide): %s", webhook_url)
        return

    pinned_url = webhook_url.replace(parsed.hostname, resolved_ip, 1)

    try:
        requests.get(
            pinned_url,
            headers={"Host": parsed.hostname},
            timeout=3,
            allow_redirects=False,
        )
        logger.info("webhook vérifié: %s (%s)", parsed.hostname, resolved_ip)
    except requests.RequestException as exc:
        logger.warning("échec vérification webhook %s: %s", parsed.hostname, exc)


class WebhookRegistrationView(View):
    """Version durcie : les contrôles de destination s'appliquent avant
    même de planifier la tâche de fond."""

    def post(self, request):
        webhook_url = request.POST.get("webhook_url")
        if not webhook_url:
            return JsonResponse({"error": "webhook_url manquante"}, status=400)

        register_webhook_task(webhook_url)

        return JsonResponse({"status": "en cours de vérification"})
