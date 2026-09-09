# Vulnérable — Privilege Escalation (CWE-269)
# Le rôle de l'utilisateur cible est modifié directement à partir de la
# valeur fournie par le client, sans vérifier que l'appelant a lui-même le
# droit d'accorder ce niveau de privilège. Un utilisateur standard pourrait
# ainsi s'attribuer ou attribuer à un tiers le rôle 'admin'.

from rest_framework.views import APIView
from rest_framework.response import Response
from .models import User


class UserRoleUpdateView(APIView):
    def put(self, request, user_id):
        user = User.objects.get(id=user_id)
        user.role = request.data["role"]  # aucune vérification du droit d'attribution
        user.save()
        return Response({"ok": True})
