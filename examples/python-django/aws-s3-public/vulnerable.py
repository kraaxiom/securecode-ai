"""
Exemple VULNERABLE — Django + boto3 : bucket S3 exposé publiquement.

CWE-284: Improper Access Control
L'application crée/téléverse des fichiers dans un bucket S3 avec une ACL
publique en lecture. N'importe qui disposant de l'URL de l'objet (ou même
en listant le bucket) peut accéder aux fichiers, y compris des documents
sensibles (factures, pièces d'identité, exports de données utilisateur).
"""

import boto3
from django.conf import settings
from django.http import JsonResponse
from django.views import View

s3_client = boto3.client(
    "s3",
    aws_access_key_id=settings.AWS_ACCESS_KEY_ID,
    aws_secret_access_key=settings.AWS_SECRET_ACCESS_KEY,
    region_name="eu-west-1",
)

BUCKET_NAME = "my-app-bucket"


class UploadInvoiceView(View):
    """Téléverse une facture générée pour un utilisateur."""

    def post(self, request):
        file_obj = request.FILES["invoice"]
        key = f"invoices/{request.user.id}/{file_obj.name}"

        # VULNERABLE : ACL "public-read" -> l'objet est accessible par tout le
        # monde via son URL directe, sans authentification ni autorisation.
        s3_client.upload_fileobj(
            file_obj,
            BUCKET_NAME,
            key,
            ExtraArgs={"ACL": "public-read"},
        )

        # VULNERABLE : l'URL publique permanente est renvoyée telle quelle,
        # aucune expiration, aucune vérification que le demandeur est bien
        # le propriétaire de la facture.
        public_url = f"https://{BUCKET_NAME}.s3.amazonaws.com/{key}"
        return JsonResponse({"url": public_url})


def create_public_bucket():
    """
    VULNERABLE : création d'un bucket puis désactivation explicite du
    Block Public Access, ce qui autorise les ACL et policies publiques.
    """
    s3_client.create_bucket(Bucket=BUCKET_NAME)
    s3_client.put_public_access_block(
        Bucket=BUCKET_NAME,
        PublicAccessBlockConfiguration={
            "BlockPublicAcls": False,
            "IgnorePublicAcls": False,
            "BlockPublicPolicy": False,
            "RestrictPublicBuckets": False,
        },
    )
