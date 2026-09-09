"""
Fonctionnalité : validation asynchrone d'URL de webhook (job "fire and
forget") — la réponse détaillée n'est jamais renvoyée au client, seul un
statut générique "en cours" est retourné.

CWE-918 (SSRF) — variante aveugle (Blind SSRF) : le serveur émet bien une
requête sortante vers une URL contrôlée par l'utilisateur, mais aucune
information sur le résultat (statut, contenu, erreur) n'est exposée dans
la réponse observable par l'appelant. L'exploitation repose alors sur des
canaux indirects (délai de traitement, callback DNS/HTTP hors bande) plutôt
que sur la lecture directe du contenu récupéré — ce qui rend la faille plus
difficile à détecter mais tout aussi dangereuse pour un scan du réseau
interne.
"""
import requests
from django.http import JsonResponse
from django.views import View


def register_webhook_task(webhook_url: str) -> None:
    """Tâche exécutée en arrière-plan (worker) : vérifie que l'URL de
    webhook répond, sans jamais exposer le résultat au client."""
    try:
        # VULNÉRABLE : aucune validation de destination, et le résultat
        # (y compris les erreurs) n'est ni renvoyé ni journalisé de façon
        # exploitable pour la détection.
        requests.get(webhook_url, timeout=5)
    except requests.RequestException:
        pass  # échec silencieux : aucune trace exploitable


class WebhookRegistrationView(View):
    """Enregistre une URL de webhook fournie par l'utilisateur et déclenche
    une vérification asynchrone (le client ne voit jamais le résultat)."""

    def post(self, request):
        webhook_url = request.POST.get("webhook_url")
        if not webhook_url:
            return JsonResponse({"error": "webhook_url manquante"}, status=400)

        # Traitement "fire and forget" : la requête part en tâche de fond,
        # aucune information sur son résultat n'est jamais exposée.
        register_webhook_task(webhook_url)

        return JsonResponse({"status": "en cours de vérification"})
