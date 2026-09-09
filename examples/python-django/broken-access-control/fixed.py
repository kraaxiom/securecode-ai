# Corrigé — Broken Access Control, catégorie générale (CWE-284)
# Ajout d'une vérification d'autorisation centralisée et explicite
# ("deny by default") avant de servir la ressource : l'accès n'est permis
# que si l'utilisateur est propriétaire de la facture ou administrateur.

from django.shortcuts import get_object_or_404
from django.http import FileResponse, HttpResponseForbidden
from django.contrib.auth.decorators import login_required
from .models import Invoice


@login_required
def download_invoice(request, invoice_id):
    invoice = get_object_or_404(Invoice, id=invoice_id)
    if not (invoice.owner_id == request.user.id or request.user.role == "admin"):
        return HttpResponseForbidden("Accès refusé")  # deny by default
    return FileResponse(invoice.file)
