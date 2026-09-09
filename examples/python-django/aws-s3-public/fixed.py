"""
Exemple CORRIGE — Django + boto3 : bucket S3 privé avec accès temporaire.

Fix pour CWE-284 (Improper Access Control) :
- Le bucket reste privé (Block Public Access activé, pas d'ACL publique).
- L'accès en lecture est accordé au cas par cas via une URL pré-signée à
  durée de vie courte, générée uniquement pour l'utilisateur propriétaire
  de la ressource (contrôle d'autorisation explicite côté serveur).
"""

import boto3
from django.conf import settings
from django.http import JsonResponse, HttpResponseForbidden
from django.views import View

s3_client = boto3.client(
    "s3",
    aws_access_key_id=settings.AWS_ACCESS_KEY_ID,
    aws_secret_access_key=settings.AWS_SECRET_ACCESS_KEY,
    region_name="eu-west-1",
)

BUCKET_NAME = "my-app-bucket"


class UploadInvoiceView(View):
    """Téléverse une facture générée pour un utilisateur, en gardant le bucket privé."""

    def post(self, request):
        file_obj = request.FILES["invoice"]
        key = f"invoices/{request.user.id}/{file_obj.name}"

        # CORRIGE : aucune ACL publique. L'objet reste privé par défaut,
        # protégé par le Block Public Access du bucket.
        s3_client.upload_fileobj(file_obj, BUCKET_NAME, key)

        return JsonResponse({"key": key})


class InvoiceDownloadUrlView(View):
    """Génère une URL pré-signée temporaire, réservée au propriétaire de la facture."""

    def get(self, request, invoice_key):
        # CORRIGE : contrôle d'autorisation explicite avant toute génération
        # d'URL — seul le propriétaire de la ressource peut y accéder.
        expected_prefix = f"invoices/{request.user.id}/"
        if not invoice_key.startswith(expected_prefix):
            return HttpResponseForbidden("Accès refusé.")

        # CORRIGE : URL pré-signée valable 15 minutes au lieu d'un accès
        # public permanent.
        presigned_url = s3_client.generate_presigned_url(
            "get_object",
            Params={"Bucket": BUCKET_NAME, "Key": invoice_key},
            ExpiresIn=900,
        )
        return JsonResponse({"url": presigned_url})


def create_private_bucket():
    """
    CORRIGE : création d'un bucket avec Block Public Access activé
    intégralement et chiffrement au repos (SSE-KMS).
    """
    s3_client.create_bucket(Bucket=BUCKET_NAME)
    s3_client.put_public_access_block(
        Bucket=BUCKET_NAME,
        PublicAccessBlockConfiguration={
            "BlockPublicAcls": True,
            "IgnorePublicAcls": True,
            "BlockPublicPolicy": True,
            "RestrictPublicBuckets": True,
        },
    )
    s3_client.put_bucket_encryption(
        Bucket=BUCKET_NAME,
        ServerSideEncryptionConfiguration={
            "Rules": [{"ApplyServerSideEncryptionByDefault": {"SSEAlgorithm": "aws:kms"}}]
        },
    )
