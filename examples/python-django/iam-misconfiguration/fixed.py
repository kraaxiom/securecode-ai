"""
Exemple CORRIGE — Django + boto3 : policy IAM à moindre privilège.

Fix pour CWE-269 (Improper Privilege Management) :
- La policy IAM du compte de service est restreinte aux seules actions
  nécessaires à sa tâche réelle (écriture d'objets dans un préfixe précis
  d'un bucket S3 donné), suivant le principe du moindre privilège.
- La ressource est explicitement scoped via un ARN précis plutôt qu'un
  wildcard global.
"""

import json

import boto3
from django.conf import settings

iam_client = boto3.client(
    "iam",
    aws_access_key_id=settings.AWS_ACCESS_KEY_ID,
    aws_secret_access_key=settings.AWS_SECRET_ACCESS_KEY,
    region_name="eu-west-1",
)

BUCKET_NAME = "my-app-documents-bucket"


def provision_document_upload_service_user(user_name: str):
    """
    Crée un utilisateur IAM de service destiné uniquement à téléverser des
    justificatifs clients, avec une policy strictement limitée à ce besoin.
    """
    iam_client.create_user(UserName=user_name)

    # CORRIGE : la policy n'autorise que les actions S3 réellement requises
    # (dépôt/lecture d'objets) et uniquement sur le préfixe et le bucket
    # concernés — pas d'accès en écriture au reste du bucket, ni à d'autres
    # services AWS (IAM, EC2, facturation, etc.).
    least_privilege_policy = {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Action": ["s3:PutObject", "s3:GetObject"],
                "Resource": f"arn:aws:s3:::{BUCKET_NAME}/uploads/*",
                "Condition": {
                    "StringEquals": {"aws:RequestedRegion": "eu-west-1"}
                },
            }
        ],
    }

    iam_client.put_user_policy(
        UserName=user_name,
        PolicyName="document-upload-policy",
        PolicyDocument=json.dumps(least_privilege_policy),
    )

    # CORRIGE : à privilégier lorsque possible — utiliser des identités
    # temporaires (rôle assumé via STS) plutôt qu'une clé d'accès statique
    # à durée de vie illimitée. Conservé ici en clé statique uniquement
    # pour illustrer le scoping de la policy ; en production, préférer
    # `sts.assume_role` ou l'identité fédérée du workload.
    access_key = iam_client.create_access_key(UserName=user_name)
    return access_key["AccessKey"]
