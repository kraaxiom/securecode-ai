# Vulnérable — Broken Access Control, catégorie générale (CWE-284)
# Aucune vérification d'autorisation n'est effectuée avant de servir le
# fichier de facture : la vue se contente de récupérer l'objet par ID et de
# le retourner, sans contrôler que l'utilisateur courant y a réellement droit.

from django.shortcuts import get_object_or_404
from django.http import FileResponse
from django.contrib.auth.decorators import login_required
from .models import Invoice


@login_required
def download_invoice(request, invoice_id):
    invoice = Invoice.objects.get(id=invoice_id)  # aucun contrôle d'autorisation
    return FileResponse(invoice.file)
