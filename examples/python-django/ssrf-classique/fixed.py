"""
Fonctionnalité : aperçu d'URL (webhook / preview de lien) — version corrigée.

Correctifs appliqués contre le CWE-918 (SSRF) :
  1. Whitelist stricte de schémas autorisés (http/https uniquement).
  2. Résolution DNS explicite du domaine, PUIS validation de l'IP obtenue
     contre les plages privées/loopback/link-local/multicast/metadata
     AVANT d'émettre la requête (évite un TOCTOU basique).
  3. Connexion effectuée directement sur l'IP validée (DNS pinning) plutôt
     que de laisser `requests` résoudre à nouveau le nom au moment de la
     connexion : protège contre le DNS rebinding (cf. dns-rebinding.md).
  4. Redirections HTTP désactivées (`allow_redirects=False`) : une
     redirection vers une cible interne ne doit jamais être suivie
     automatiquement.
  5. Timeout court pour limiter l'impact d'un scan de ports interne.
"""
import ipaddress
import socket
from urllib.parse import urlparse

import requests
from django.http import JsonResponse
from django.views import View

ALLOWED_SCHEMES = {"http", "https"}

# Domaines métier explicitement autorisés pour la fonctionnalité de preview.
ALLOWED_DOMAINS = {
    "example.com",
    "cdn.example.com",
}


def _is_public_ip(ip_str: str) -> bool:
    """Rejette toute IP privée, loopback, link-local (dont 169.254.169.254,
    l'adresse des services de métadonnées cloud), multicast ou réservée."""
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
    """Résout le nom d'hôte et retourne la première IP publique valide,
    ou lève ValueError. Résoudre une seule fois puis épingler l'IP évite
    le DNS rebinding (double résolution DNS/HTTP)."""
    try:
        infos = socket.getaddrinfo(hostname, None)
    except socket.gaierror as exc:
        raise ValueError(f"Résolution DNS impossible pour {hostname}") from exc

    for family, _, _, _, sockaddr in infos:
        ip_str = sockaddr[0]
        if _is_public_ip(ip_str):
            return ip_str

    raise ValueError(f"Aucune IP publique valide pour {hostname}")


class UrlPreviewView(View):
    """Version durcie : whitelist de domaines + résolution/validation IP
    + pinning + blocage des redirections."""

    def post(self, request):
        target_url = request.POST.get("url")
        if not target_url:
            return JsonResponse({"error": "url manquante"}, status=400)

        parsed = urlparse(target_url)

        if parsed.scheme not in ALLOWED_SCHEMES:
            return JsonResponse({"error": "schéma non autorisé"}, status=400)

        if parsed.hostname not in ALLOWED_DOMAINS:
            return JsonResponse({"error": "domaine non autorisé"}, status=400)

        try:
            resolved_ip = resolve_and_validate(parsed.hostname)
        except ValueError:
            return JsonResponse({"error": "destination invalide"}, status=400)

        # On force la connexion sur l'IP résolue et validée (pinning),
        # tout en conservant l'en-tête Host attendu par le serveur cible.
        pinned_url = target_url.replace(parsed.hostname, resolved_ip, 1)

        try:
            response = requests.get(
                pinned_url,
                headers={"Host": parsed.hostname},
                timeout=3,
                allow_redirects=False,  # pas de suivi aveugle de redirection
                verify=True,
            )
        except requests.RequestException:
            return JsonResponse({"error": "échec de la requête"}, status=502)

        return JsonResponse({
            "status_code": response.status_code,
            "content_type": response.headers.get("Content-Type"),
            "excerpt": response.text[:500],
        })
