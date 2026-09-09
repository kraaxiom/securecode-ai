"""
Exemple VULNERABLE — Django + azure-storage-blob : conteneur Blob public.

CWE-284: Improper Access Control
Le conteneur est créé avec un accès public anonyme en lecture au niveau
"container", ce qui permet à quiconque de lister et télécharger tous les
blobs sans authentification.
"""

from azure.storage.blob import BlobServiceClient, PublicAccess
from django.conf import settings
from django.http import JsonResponse
from django.views import View

blob_service_client = BlobServiceClient.from_connection_string(
    settings.AZURE_STORAGE_CONNECTION_STRING
)

CONTAINER_NAME = "myfiles"


def create_public_container():
    """VULNERABLE : conteneur créé avec accès public anonyme au niveau container."""
    container_client = blob_service_client.get_container_client(CONTAINER_NAME)
    # VULNERABLE : public_access="container" autorise la lecture ET le
    # listing anonyme de tous les blobs du conteneur.
    container_client.create_container(public_access=PublicAccess.CONTAINER)


class UploadDocumentView(View):
    """Téléverse un document utilisateur dans le conteneur public."""

    def post(self, request):
        file_obj = request.FILES["document"]
        blob_name = f"documents/{request.user.id}/{file_obj.name}"

        container_client = blob_service_client.get_container_client(CONTAINER_NAME)
        # VULNERABLE : le blob hérite de l'accès public du conteneur,
        # accessible directement via son URL, sans expiration ni contrôle.
        container_client.upload_blob(name=blob_name, data=file_obj, overwrite=True)

        blob_url = (
            f"{blob_service_client.url}{CONTAINER_NAME}/{blob_name}"
        )
        return JsonResponse({"url": blob_url})
