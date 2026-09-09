# Vulnérable — Broken Function Level Authorization (CWE-862)
# Le endpoint de suppression d'utilisateur ne vérifie que l'authentification
# (IsAuthenticated), sans jamais contrôler que l'appelant possède le rôle
# admin requis pour effectuer cette action sensible. Un utilisateur standard
# authentifié peut donc appeler directement cette action.

from rest_framework import viewsets
from rest_framework.permissions import IsAuthenticated
from .models import User
from .serializers import UserSerializer


class UserViewSet(viewsets.ModelViewSet):
    """Gère les comptes utilisateurs, y compris leur suppression."""

    queryset = User.objects.all()
    serializer_class = UserSerializer
    permission_classes = [IsAuthenticated]

    def destroy(self, request, *args, **kwargs):
        # Aucune vérification de rôle : seule l'authentification est requise.
        user = self.get_object()
        user.delete()
        return self.finalize_response(
            request, self.get_success_headers({}), *args, **kwargs
        )
