# Corrigé — Privilege Escalation (CWE-269)
# Vérification explicite que l'appelant est autorisé à accorder le rôle
# demandé (can_grant_role) avant toute modification, puis invalidation des
# sessions actives de la cible pour éviter la persistance d'anciens
# privilèges après le changement de rôle.

from rest_framework.views import APIView
from rest_framework.response import Response
from django.contrib.sessions.models import Session
from .models import User


class UserRoleUpdateView(APIView):
    def put(self, request, user_id):
        requested_role = request.data["role"]
        if not request.user.can_grant_role(requested_role):
            return Response({"error": "Attribution non autorisée"}, status=403)
        user = User.objects.get(id=user_id)
        user.role = requested_role
        user.save()
        Session.objects.filter(user=user).delete()  # invalide les sessions actives
        return Response({"ok": True})
