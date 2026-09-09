# Corrigé — Insecure Direct Object Reference (CWE-639)
# La requête inclut désormais une clause de filtrage sur le propriétaire
# (owner=request.user) directement dans get_object_or_404, garantissant
# qu'un utilisateur ne peut accéder qu'à ses propres documents.

from django.contrib.auth.decorators import login_required
from django.shortcuts import get_object_or_404
from django.http import JsonResponse
from .models import Document


@login_required
def get_document(request, doc_id):
    doc = get_object_or_404(Document, id=doc_id, owner=request.user)
    return JsonResponse(doc.to_dict())
