"""
Fonctionnalité : validateur d'URL de webhook — la destination est validée
une première fois (résolution DNS + vérification d'IP publique), puis la
requête HTTP réelle est effectuée séparément en laissant `requests`
résoudre le nom de domaine une seconde fois.

CWE-918 (SSRF) — variante DNS rebinding : l'attaquant contrôle un domaine
dont la réponse DNS change entre le moment de la validation (résolution
vers une IP publique légitime, TTL très court) et le moment de la requête
HTTP réelle (résolution vers une IP interne). Le décalage temporel entre
les deux résolutions DNS (TOCTOU) permet de contourner un filtre basé
uniquement sur une validation d'IP faite "avant" la requête.
"""
import ipaddress
import socket

import requests
from django.http import JsonResponse
from django.views import View


def _is_public_ip(ip_str: str) -> bool:
    ip = ipaddress.ip_address(ip_str)
    return not (ip.is_private or ip.is_loopback or ip.is_link_local or ip.is_reserved)


class WebhookUrlValidatorView(View):
    """Valide une URL de webhook puis, séparément, envoie une requête de
    test — deux résolutions DNS distinctes, exploitables par rebinding."""

    def post(self, request):
        hostname = request.POST.get("hostname")
        if not hostname:
            return JsonResponse({"error": "hostname manquant"}, status=400)

        # Étape 1 : validation — résolution DNS n°1.
        try:
            ip_str = socket.gethostbyname(hostname)
        except socket.gaierror:
            return JsonResponse({"error": "résolution DNS impossible"}, status=400)

        if not _is_public_ip(ip_str):
            return JsonResponse({"error": "IP interne refusée"}, status=400)

        # VULNÉRABLE : entre cette validation et l'appel ci-dessous, un
        # attaquant contrôlant le DNS du domaine peut changer la réponse
        # (TTL court) pour pointer vers une IP interne. `requests` va
        # résoudre le nom une SECONDE fois (résolution DNS n°2) et se
        # connecter à cette nouvelle adresse, potentiellement interne.
        response = requests.get(f"https://{hostname}/health", timeout=5)

        return JsonResponse({"status_code": response.status_code})
