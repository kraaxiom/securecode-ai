"""
Exemple CORRIGE — Django : chiffrement de pièces jointes avec AES-256-GCM.

Fix pour CWE-327 (Use of a Broken or Risky Cryptographic Algorithm) :
- Remplacement de RC4 par AES-256-GCM, chiffrement authentifié sans biais
  statistique connu, conforme aux recommandations actuelles (RFC 7465
  interdit RC4 dans TLS depuis 2015).
- La clé est chargée depuis la configuration Django (issue d'un gestionnaire
  de secrets), jamais codée en dur.
- Un nonce aléatoire unique est généré à chaque chiffrement.
"""

import os

from cryptography.hazmat.primitives.ciphers.aead import AESGCM
from django.conf import settings
from django.http import JsonResponse, HttpResponseBadRequest
from django.views import View

# CORRIGE : clé AES-256 (32 octets) chargée depuis settings, elle-même
# alimentée par une variable d'environnement ou un coffre-fort de secrets.
AES_KEY = settings.ENCRYPTION_KEY  # 32 octets, jamais codé en dur ici


class UploadAttachmentView(View):
    """Chiffre et stocke une pièce jointe sensible (contrat, relevé, etc.)."""

    def post(self, request):
        file_obj = request.FILES["attachment"]
        raw_data = file_obj.read()

        # CORRIGE : AES-256-GCM, chiffrement authentifié sans biais
        # statistique connu, recommandé par l'OWASP pour le chiffrement
        # symétrique applicatif.
        aesgcm = AESGCM(AES_KEY)

        # CORRIGE : nonce aléatoire de 12 octets généré à chaque fichier,
        # jamais réutilisé pour la même clé.
        nonce = os.urandom(12)
        ciphertext = aesgcm.encrypt(nonce, raw_data, associated_data=None)

        storage_path = f"attachments/{request.user.id}/{file_obj.name}.enc"
        with open(storage_path, "wb") as f:
            # CORRIGE : le nonce est préfixé au texte chiffré ; il n'est pas
            # secret mais doit être unique et conservé pour le déchiffrement.
            f.write(nonce + ciphertext)

        return JsonResponse({"status": "stored", "path": storage_path})


class DownloadAttachmentView(View):
    """Déchiffre une pièce jointe stockée pour la renvoyer à l'utilisateur."""

    def get(self, request, storage_path):
        with open(storage_path, "rb") as f:
            raw = f.read()
        nonce, ciphertext = raw[:12], raw[12:]

        aesgcm = AESGCM(AES_KEY)
        try:
            # CORRIGE : AES-GCM vérifie l'intégrité et l'authenticité du
            # fichier chiffré ; toute altération fait échouer le
            # déchiffrement au lieu de renvoyer des données corrompues.
            plaintext = aesgcm.decrypt(nonce, ciphertext, associated_data=None)
        except Exception:
            return HttpResponseBadRequest("Fichier corrompu ou falsifié.")

        return JsonResponse({"content_base64": plaintext.hex()})
