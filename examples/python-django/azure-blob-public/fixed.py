"""
Exemple CORRIGE — Django + azure-storage-blob : conteneur privé + SAS.

Fix pour CWE-284 (Improper Access Control) :
- Le conteneur est créé sans accès public (private par défaut).
- L'accès en lecture est accordé via un jeton SAS à durée de vie courte,
  scoped en lecture seule et généré uniquement pour le propriétaire de la
  ressource.
"""

from datetime import datetime, timedelta, timezone

from azure.storage.blob import (
    BlobServiceClient,
    BlobSasPermissions,
    generate_blob_sas,
)
from django.conf import settings
from django.http import JsonResponse, HttpResponseForbidden
from django.views import View

blob_service_client = BlobServiceClient.from_connection_string(
    settings.AZURE_STORAGE_CONNECTION_STRING
)

CONTAINER_NAME = "myfiles"


def create_private_container():
    """CORRIGE : conteneur créé sans accès public (private par défaut)."""
    container_client = blob_service_client.get_container_client(CONTAINER_NAME)
    # CORRIGE : aucun paramètre public_access -> conteneur privé.
    container_client.create_container()


class UploadDocumentView(View):
    """Téléverse un document utilisateur dans le conteneur privé."""

    def post(self, request):
        file_obj = request.FILES["document"]
        blob_name = f"documents/{request.user.id}/{file_obj.name}"

        container_client = blob_service_client.get_container_client(CONTAINER_NAME)
        container_client.upload_blob(name=blob_name, data=file_obj, overwrite=True)

        return JsonResponse({"blob_name": blob_name})


class DocumentDownloadUrlView(View):
    """Génère une URL SAS temporaire réservée au propriétaire du document."""

    def get(self, request, blob_name):
        # CORRIGE : contrôle d'autorisation explicite avant toute génération
        # de jeton d'accès.
        expected_prefix = f"documents/{request.user.id}/"
        if not blob_name.startswith(expected_prefix):
            return HttpResponseForbidden("Accès refusé.")

        account_name = blob_service_client.account_name
        account_key = settings.AZURE_STORAGE_ACCOUNT_KEY

        # CORRIGE : SAS en lecture seule, expirant dans 1 heure, au lieu
        # d'un accès public permanent.
        sas_token = generate_blob_sas(
            account_name=account_name,
            container_name=CONTAINER_NAME,
            blob_name=blob_name,
            account_key=account_key,
            permission=BlobSasPermissions(read=True),
            expiry=datetime.now(timezone.utc) + timedelta(hours=1),
        )

        url = f"{blob_service_client.url}{CONTAINER_NAME}/{blob_name}?{sas_token}"
        return JsonResponse({"url": url})
