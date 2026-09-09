"""
Exemple CORRIGE — Django : chiffrement des données sensibles avec AES-256-GCM.

Fix pour CWE-327 (Use of a Broken or Risky Cryptographic Algorithm) :
- Remplacement de 3DES par AES-256-GCM, un chiffrement authentifié moderne
  avec un bloc de 128 bits, insensible à Sweet32.
- La clé est chargée depuis la configuration Django (issue d'un gestionnaire
  de secrets / variable d'environnement), jamais codée en dur.
- Un nonce aléatoire unique est généré à chaque chiffrement et stocké aux
  côtés du texte chiffré, comme l'exige AES-GCM.
"""

import os

from cryptography.hazmat.primitives.ciphers.aead import AESGCM
from django.conf import settings
from django.http import JsonResponse, HttpResponseBadRequest
from django.views import View

# CORRIGE : clé AES-256 (32 octets) chargée depuis settings, elle-même
# alimentée par une variable d'environnement ou un coffre-fort de secrets.
# Voir settings.py : ENCRYPTION_KEY = base64.b64decode(os.environ["ENCRYPTION_KEY"])
AES_KEY = settings.ENCRYPTION_KEY  # 32 octets, jamais codé en dur ici


class EncryptSocialSecurityView(View):
    """Chiffre le numéro de sécurité sociale d'un utilisateur avant stockage."""

    def post(self, request):
        ssn = request.POST["ssn"].encode("utf-8")

        # CORRIGE : AES-256-GCM — chiffrement authentifié, bloc de 128 bits,
        # aucune vulnérabilité connue de type Sweet32.
        aesgcm = AESGCM(AES_KEY)

        # CORRIGE : nonce aléatoire de 12 octets généré à chaque appel,
        # jamais réutilisé pour la même clé.
        nonce = os.urandom(12)
        ciphertext = aesgcm.encrypt(nonce, ssn, associated_data=None)

        # CORRIGE : le nonce doit être conservé avec le texte chiffré pour
        # permettre le déchiffrement ultérieur ; il n'est pas secret.
        request.user.encrypted_ssn = (nonce + ciphertext).hex()
        request.user.save(update_fields=["encrypted_ssn"])

        return JsonResponse({"status": "encrypted"})


class DecryptSocialSecurityView(View):
    """Déchiffre le numéro de sécurité sociale pour affichage administrateur."""

    def get(self, request):
        raw = bytes.fromhex(request.user.encrypted_ssn)
        nonce, ciphertext = raw[:12], raw[12:]

        aesgcm = AESGCM(AES_KEY)
        try:
            # CORRIGE : AES-GCM vérifie l'intégrité et l'authenticité du
            # texte chiffré ; toute altération fait échouer le déchiffrement
            # au lieu de renvoyer silencieusement des données corrompues.
            plaintext = aesgcm.decrypt(nonce, ciphertext, associated_data=None)
        except Exception:
            return HttpResponseBadRequest("Donnée corrompue ou falsifiée.")

        return JsonResponse({"ssn": plaintext.decode("utf-8")})
