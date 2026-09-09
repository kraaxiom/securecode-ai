"""
Exemple CORRIGE — Django : hachage de mots de passe avec Argon2id.

Fix pour CWE-328 (Use of Weak Hash) :
- Remplacement de MD5 par Argon2id, algorithme de hachage de mot de passe
  dédié, avec sel automatique intégré et facteur de coût configurable.
- La vérification utilise une comparaison sécurisée fournie par la
  bibliothèque (résistante aux attaques temporelles) au lieu d'une
  comparaison de chaînes manuelle.
- Prévoit un plan de migration progressive pour les hachages MD5 existants.
"""

from argon2 import PasswordHasher
from argon2.exceptions import VerifyMismatchError, InvalidHash
from django.http import JsonResponse
from django.views import View

# CORRIGE : Argon2id est le choix recommandé par l'OWASP pour le hachage de
# mots de passe. Le sel est généré et géré automatiquement par la bibliothèque.
password_hasher = PasswordHasher()


class SignupView(View):
    """Inscrit un nouvel utilisateur en stockant son mot de passe haché."""

    def post(self, request):
        username = request.POST["username"]
        password = request.POST["password"]

        # CORRIGE : Argon2id génère un sel aléatoire unique par mot de passe
        # et applique un facteur de coût (temps/mémoire/parallélisme) rendant
        # la force brute massivement plus coûteuse.
        password_hash = password_hasher.hash(password)

        from myapp.models import User

        User.objects.create(username=username, password_hash=password_hash)
        return JsonResponse({"status": "created"})


class LoginView(View):
    """Authentifie un utilisateur en vérifiant le hachage Argon2id."""

    def post(self, request):
        username = request.POST["username"]
        password = request.POST["password"]

        from myapp.models import User

        user = User.objects.filter(username=username).first()
        if user is None:
            return JsonResponse({"error": "invalid_credentials"}, status=401)

        try:
            # CORRIGE : password_hasher.verify() effectue une comparaison
            # sécurisée et lève une exception explicite en cas d'échec,
            # au lieu d'une comparaison de chaînes vulnérable aux attaques
            # temporelles.
            password_hasher.verify(user.password_hash, password)
        except (VerifyMismatchError, InvalidHash):
            return JsonResponse({"error": "invalid_credentials"}, status=401)

        # CORRIGE : migration progressive — si le hash stocké utilise encore
        # un ancien paramétrage (ou provient d'une migration MD5 -> Argon2id
        # avec re-hachage au prochain login), on le renouvelle ici.
        if password_hasher.check_needs_rehash(user.password_hash):
            user.password_hash = password_hasher.hash(password)
            user.save(update_fields=["password_hash"])

        return JsonResponse({"status": "authenticated"})
