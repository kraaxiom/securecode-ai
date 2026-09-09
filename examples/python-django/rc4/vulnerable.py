"""
Exemple VULNERABLE — Django : chiffrement de pièces jointes avec RC4.

CWE-327: Use of a Broken or Risky Cryptographic Algorithm
L'application chiffre le contenu de pièces jointes sensibles avec RC4 avant
de les stocker. RC4 est un chiffrement par flux comportant des biais
statistiques connus et exploitables dans son flux de sortie (attaques
pratiques démontrées contre WEP et contre les cookies chiffrés en TLS-RC4).
Il est interdit dans TLS depuis la RFC 7465 (2015).
"""

from Crypto.Cipher import ARC4
from django.conf import settings
from django.http import JsonResponse
from django.views import View

# VULNERABLE : clé RC4 codée en dur, identique pour tous les documents
# chiffrés par l'application.
RC4_KEY = b"legacy-static-key-2018"


class UploadAttachmentView(View):
    """Chiffre et stocke une pièce jointe sensible (contrat, relevé, etc.)."""

    def post(self, request):
        file_obj = request.FILES["attachment"]
        raw_data = file_obj.read()

        # VULNERABLE : RC4 est un chiffrement par flux cassé, avec des biais
        # statistiques dans les premiers octets du flux de sortie qui
        # facilitent la récupération partielle du texte en clair.
        cipher = ARC4.new(RC4_KEY)
        ciphertext = cipher.encrypt(raw_data)

        storage_path = f"attachments/{request.user.id}/{file_obj.name}.enc"
        with open(storage_path, "wb") as f:
            f.write(ciphertext)

        return JsonResponse({"status": "stored", "path": storage_path})


class DownloadAttachmentView(View):
    """Déchiffre une pièce jointe stockée pour la renvoyer à l'utilisateur."""

    def get(self, request, storage_path):
        with open(storage_path, "rb") as f:
            ciphertext = f.read()

        # VULNERABLE : réutilisation de la même clé RC4 statique pour tous
        # les fichiers, sans nonce ni vecteur d'initialisation — deux
        # fichiers chiffrés avec la même clé produisent des flux corrélés.
        cipher = ARC4.new(RC4_KEY)
        plaintext = cipher.decrypt(ciphertext)

        return JsonResponse({"content_base64": plaintext.hex()})
