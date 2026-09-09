# Vulnérable — Forced Browsing (CWE-425)
# La page de confirmation de paiement est accessible directement, sans
# vérifier que les étapes précédentes du flux (panier, paiement autorisé)
# ont réellement été complétées côté serveur. Un attaquant peut sauter
# directement à cette étape en devinant/forçant l'URL.

from django.shortcuts import render


def payment_confirmation(request):
    # Aucune vérification de l'état d'avancement réel du flux de paiement.
    return render(request, "confirmation.html")
