# Corrigé — Broken Function Level Authorization (CWE-862)
# Ajout d'une permission de rôle explicite (IsAdmin) appliquée uniquement à
# l'action sensible "destroy", en complément de IsAuthenticated. Le contrôle
# est centralisé dans une classe de permission réutilisable, pas dupliqué
# dans chaque vue.

from rest_framework import viewsets
from rest_framework.permissions import IsAuthenticated, BasePermission
from .models import User
from .serializers import UserSerializer


class IsAdmin(BasePermission):
    def has_permission(self, request, view):
        return bool(
            request.user
            and request.user.is_authenticated
            and request.user.role == "admin"
        )


class UserViewSet(viewsets.ModelViewSet):
    """Gère les comptes utilisateurs, y compris leur suppression."""

    queryset = User.objects.all()
    serializer_class = UserSerializer
    permission_classes = [IsAuthenticated]

    def get_permissions(self):
        # Contrôle de rôle explicite pour l'action de suppression.
        if self.action == "destroy":
            return [IsAdmin()]
        return super().get_permissions()

    def destroy(self, request, *args, **kwargs):
        user = self.get_object()
        user.delete()
        return self.finalize_response(
            request, self.get_success_headers({}), *args, **kwargs
        )
