"""
Exemple VULNERABLE — Django : clés cryptographiques et secrets codés en dur.

CWE-798: Use of Hard-coded Credentials
La clé de signature des tokens et les identifiants d'un service tiers sont
écrits littéralement dans le code source. Quiconque a accès au dépôt Git
(y compris à son historique) ou au binaire déployé peut extraire ces
secrets, annulant toute garantie de confidentialité offerte par les
mécanismes cryptographiques qui les utilisent.
"""

import hmac
import hashlib

from django.http import JsonResponse
from django.views import View

# VULNERABLE : clé de signature HMAC codée en dur dans le code source,
# identique en dev, staging et production, et visible dans l'historique Git.
TOKEN_SIGNING_KEY = "s3cr3t-signing-key-2021-do-not-share"

# VULNERABLE : identifiants d'un fournisseur tiers (paiement) en clair dans
# le code, jamais régénérés depuis la mise en production initiale.
PAYMENT_API_KEY = "sk_live_EXAMPLE_NOT_A_REAL_KEY"


def sign_password_reset_token(user_id: int, expires_at: int) -> str:
    """Génère un token signé pour la réinitialisation de mot de passe."""
    message = f"{user_id}:{expires_at}".encode("utf-8")

    # VULNERABLE : la clé de signature étant publique de facto (présente
    # dans le code source), n'importe qui peut forger des tokens valides
    # pour n'importe quel utilisateur.
    signature = hmac.new(TOKEN_SIGNING_KEY.encode(), message, hashlib.sha256).hexdigest()
    return f"{user_id}:{expires_at}:{signature}"


class ChargeCardView(View):
    """Débite la carte bancaire d'un utilisateur via le fournisseur de paiement."""

    def post(self, request):
        import stripe

        # VULNERABLE : la clé d'API de paiement codée en dur permet à
        # quiconque lit le code source de débiter des cartes ou d'accéder
        # au tableau de bord du compte marchand.
        stripe.api_key = PAYMENT_API_KEY
        charge = stripe.Charge.create(
            amount=int(request.POST["amount_cents"]),
            currency="xof",
            source=request.POST["token"],
        )
        return JsonResponse({"charge_id": charge.id})
