# Vulnérable — Insecure Direct Object Reference (CWE-639)
# Le document est récupéré uniquement à partir de l'ID fourni par le client,
# sans aucune clause vérifiant que l'utilisateur courant en est bien le
# propriétaire. Modifier l'ID dans l'URL suffit à accéder au document d'un
# autre utilisateur.

from django.contrib.auth.decorators import login_required
from django.shortcuts import get_object_or_404
from django.http import JsonResponse
from .models import Document


@login_required
def get_document(request, doc_id):
    doc = get_object_or_404(Document, id=doc_id)  # pas de filtrage par propriétaire
    return JsonResponse(doc.to_dict())
