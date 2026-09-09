# Corrigé — Forced Browsing (CWE-425)
# La vue vérifie côté serveur, via la session, que le paiement a réellement
# été autorisé avant d'afficher la page de confirmation. L'accès direct en
# sautant les étapes précédentes est désormais refusé.

from django.shortcuts import render, redirect


def payment_confirmation(request):
    checkout = request.session.get("checkout")
    if not checkout or checkout.get("status") != "payment_authorized":
        # Refuse l'accès direct à l'étape finale sans avoir complété le flux.
        return redirect("checkout_start")
    return render(request, "confirmation.html")
