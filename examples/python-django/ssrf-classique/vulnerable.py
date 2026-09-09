"""
Fonctionnalité : aperçu d'URL (webhook / preview de lien) — vue Django.

CWE-918 (Server-Side Request Forgery) : l'URL fournie par l'utilisateur est
transmise telle quelle à `requests.get()`, sans aucune restriction de
destination (whitelist de domaines) ni filtrage des plages d'adresses
privées/loopback/link-local. Un attaquant peut ainsi forcer le serveur à
émettre des requêtes vers des ressources internes normalement inaccessibles
depuis l'extérieur (services d'administration, bases de données, API de
métadonnées cloud, etc.), et récupérer la réponse via cette fonctionnalité
de preview.
"""
import requests
from django.http import JsonResponse
from django.views import View


class UrlPreviewView(View):
    """Récupère le contenu d'une URL fournie par l'utilisateur pour en
    générer un aperçu (ex: aperçu de lien partagé dans un message)."""

    def post(self, request):
        target_url = request.POST.get("url")
        if not target_url:
            return JsonResponse({"error": "url manquante"}, status=400)

        # VULNÉRABLE : aucune validation de l'hôte/IP de destination.
        # L'utilisateur contrôle intégralement `target_url`, y compris
        # des cibles comme http://127.0.0.1/admin ou http://10.0.0.5:6379/.
        response = requests.get(target_url, timeout=5)

        return JsonResponse({
            "status_code": response.status_code,
            "content_type": response.headers.get("Content-Type"),
            "excerpt": response.text[:500],
        })
