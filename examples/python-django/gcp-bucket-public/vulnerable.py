"""
Exemple VULNERABLE — Django + google-cloud-storage : bucket GCS public.

CWE-284: Improper Access Control
Le binding IAM du bucket accorde le rôle "objectViewer" à `allUsers`,
rendant tous les objets du bucket lisibles publiquement et anonymement.
"""

from django.conf import settings
from django.http import JsonResponse
from django.views import View
from google.cloud import storage

storage_client = storage.Client.from_service_account_json(
    settings.GCP_SERVICE_ACCOUNT_JSON_PATH
)

BUCKET_NAME = "my-app-bucket"


def create_public_bucket():
    """VULNERABLE : accorde un accès public en lecture à tout le bucket."""
    bucket = storage_client.bucket(BUCKET_NAME)
    policy = bucket.get_iam_policy(requested_policy_version=3)
    # VULNERABLE : "allUsers" -> n'importe qui sur Internet peut lister et
    # télécharger tous les objets du bucket, sans authentification.
    policy.bindings.append(
        {"role": "roles/storage.objectViewer", "members": {"allUsers"}}
    )
    bucket.set_iam_policy(policy)


class UploadExportView(View):
    """Téléverse un export de données utilisateur dans le bucket public."""

    def post(self, request):
        file_obj = request.FILES["export"]
        blob_name = f"exports/{request.user.id}/{file_obj.name}"

        bucket = storage_client.bucket(BUCKET_NAME)
        blob = bucket.blob(blob_name)
        blob.upload_from_file(file_obj)

        # VULNERABLE : URL publique permanente renvoyée directement, sans
        # expiration ni vérification du demandeur.
        public_url = f"https://storage.googleapis.com/{BUCKET_NAME}/{blob_name}"
        return JsonResponse({"url": public_url})
