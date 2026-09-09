"""
Exemple VULNERABLE — Django + boto3 : policy IAM avec permissions excessives.

CWE-269: Improper Privilege Management
Lors du provisionnement d'un utilisateur de service pour une tâche
applicative précise (upload de justificatifs dans un bucket S3 donné),
le code attache une policy IAM inline accordant `"Action": "*"` sur
`"Resource": "*"` — un accès administrateur complet sur le compte AWS —
alors que la tâche ne nécessite que d'écrire dans un unique bucket.
"""

import boto3
from django.conf import settings

iam_client = boto3.client(
    "iam",
    aws_access_key_id=settings.AWS_ACCESS_KEY_ID,
    aws_secret_access_key=settings.AWS_SECRET_ACCESS_KEY,
    region_name="eu-west-1",
)


def provision_document_upload_service_user(user_name: str):
    """
    Crée un utilisateur IAM de service destiné uniquement à téléverser des
    justificatifs clients dans le bucket S3 de l'application.
    """
    iam_client.create_user(UserName=user_name)

    # VULNERABLE : la policy accorde toutes les actions sur toutes les
    # ressources du compte AWS (S3, EC2, IAM, facturation, etc.) alors que
    # ce compte de service ne devrait pouvoir qu'écrire dans un bucket précis.
    # En cas de fuite des identifiants de ce compte, l'attaquant obtient un
    # contrôle administrateur complet sur l'infrastructure cloud.
    overly_broad_policy = {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Action": "*",
                "Resource": "*",
            }
        ],
    }

    iam_client.put_user_policy(
        UserName=user_name,
        PolicyName="document-upload-policy",
        PolicyDocument=str(overly_broad_policy).replace("'", '"'),
    )

    access_key = iam_client.create_access_key(UserName=user_name)
    return access_key["AccessKey"]
