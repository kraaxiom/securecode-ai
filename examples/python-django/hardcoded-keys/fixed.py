"""
Exemple CORRIGE — Django : chargement des secrets depuis l'environnement.

Fix pour CWE-798 (Use of Hard-coded Credentials) :
- Aucun secret n'est écrit dans le code source.
- La clé de signature HMAC et la clé d'API de paiement sont chargées
  depuis les variables d'environnement (elles-mêmes alimentées par un
  gestionnaire de secrets en production), avec échec explicite si absentes.
- Les secrets historiquement exposés doivent être révoqués et régénérés
  côté fournisseur, pas seulement retirés du code.
"""

import hmac
import hashlib
import os

from django.http import JsonResponse, HttpResponseServerError
from django.views import View


def _get_required_env(name: str) -> str:
    """Charge une variable d'environnement obligatoire, échoue explicitement sinon."""
    value = os.environ.get(name)
    if not value:
        # CORRIGE : échec explicite au démarrage/à l'appel plutôt qu'une
        # valeur par défaut codée en dur qui masquerait le problème.
        raise RuntimeError(f"Variable d'environnement manquante : {name}")
    return value


# CORRIGE : clé de signature chargée depuis l'environnement, distincte par
# environnement (dev/staging/prod), jamais commitée dans le dépôt.
TOKEN_SIGNING_KEY = _get_required_env("TOKEN_SIGNING_KEY")

# CORRIGE : clé d'API de paiement chargée depuis l'environnement, injectée
# de façon sécurisée par le gestionnaire de secrets (Vault, AWS Secrets
# Manager, etc.) au déploiement.
PAYMENT_API_KEY = _get_required_env("STRIPE_SECRET_KEY")


def sign_password_reset_token(user_id: int, expires_at: int) -> str:
    """Génère un token signé pour la réinitialisation de mot de passe."""
    message = f"{user_id}:{expires_at}".encode("utf-8")

    # CORRIGE : la clé de signature n'est jamais présente dans le code
    # source ni dans l'historique Git, elle est propre à chaque environnement.
    signature = hmac.new(TOKEN_SIGNING_KEY.encode(), message, hashlib.sha256).hexdigest()
    return f"{user_id}:{expires_at}:{signature}"


class ChargeCardView(View):
    """Débite la carte bancaire d'un utilisateur via le fournisseur de paiement."""

    def post(self, request):
        import stripe

        try:
            # CORRIGE : clé d'API chargée dynamiquement depuis
            # l'environnement, jamais visible dans le code source.
            stripe.api_key = PAYMENT_API_KEY
            charge = stripe.Charge.create(
                amount=int(request.POST["amount_cents"]),
                currency="xof",
                source=request.POST["token"],
            )
        except Exception:
            return HttpResponseServerError("Erreur lors du paiement.")

        return JsonResponse({"charge_id": charge.id})
