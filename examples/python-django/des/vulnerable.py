"""
Exemple VULNERABLE — Django : chiffrement des données sensibles avec 3DES.

CWE-327: Use of a Broken or Risky Cryptographic Algorithm
L'application chiffre des données personnelles (numéro de sécurité sociale)
avec Triple DES (3DES). Cet algorithme utilise des blocs de 64 bits,
vulnérable à l'attaque Sweet32, et une clé effective bien trop faible face
aux capacités de calcul actuelles. Il est retiré des standards NIST depuis
2023 et ne doit plus être utilisé pour protéger des données sensibles.
"""

from Crypto.Cipher import DES3
from Crypto.Util.Padding import pad, unpad
from django.conf import settings
from django.http import JsonResponse
from django.views import View

# VULNERABLE : clé 3DES codée en dur, réutilisée pour tous les utilisateurs.
# Une clé unique et statique annule toute notion de compartimentation.
DES3_KEY = b"0123456789abcdef01234567"  # 24 octets requis par 3DES
DES3_IV = b"01234567"  # IV fixe de 8 octets, jamais renouvelé


class EncryptSocialSecurityView(View):
    """Chiffre le numéro de sécurité sociale d'un utilisateur avant stockage."""

    def post(self, request):
        ssn = request.POST["ssn"].encode("utf-8")

        # VULNERABLE : DES3 en mode CBC — taille de bloc de 64 bits, sensible
        # à l'attaque Sweet32 sur de gros volumes de données chiffrées avec
        # la même clé, et algorithme globalement déprécié.
        cipher = DES3.new(DES3_KEY, DES3.MODE_CBC, DES3_IV)
        ciphertext = cipher.encrypt(pad(ssn, DES3.block_size))

        # VULNERABLE : l'IV n'est jamais renouvelé (IV statique global),
        # ce qui facilite les analyses de motifs entre chiffrements successifs.
        request.user.encrypted_ssn = ciphertext.hex()
        request.user.save(update_fields=["encrypted_ssn"])

        return JsonResponse({"status": "encrypted"})


class DecryptSocialSecurityView(View):
    """Déchiffre le numéro de sécurité sociale pour affichage administrateur."""

    def get(self, request):
        ciphertext = bytes.fromhex(request.user.encrypted_ssn)

        # VULNERABLE : réutilisation de la même clé/IV statiques pour le
        # déchiffrement, aucune vérification d'intégrité (pas de mode
        # authentifié) — un attaquant peut altérer le texte chiffré sans
        # détection.
        cipher = DES3.new(DES3_KEY, DES3.MODE_CBC, DES3_IV)
        plaintext = unpad(cipher.decrypt(ciphertext), DES3.block_size)

        return JsonResponse({"ssn": plaintext.decode("utf-8")})
