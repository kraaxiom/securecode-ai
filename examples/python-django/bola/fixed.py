# Corrigé — Broken Object Level Authorization (CWE-639)
# Le queryset est filtré par le propriétaire directement dans la requête de
# données (get_queryset), garantissant qu'un utilisateur ne peut jamais
# récupérer, modifier ou supprimer une commande qui ne lui appartient pas.

from rest_framework import viewsets
from rest_framework.permissions import IsAuthenticated
from .models import Order
from .serializers import OrderSerializer


class OrderViewSet(viewsets.ModelViewSet):
    """API de gestion des commandes."""

    serializer_class = OrderSerializer
    permission_classes = [IsAuthenticated]

    def get_queryset(self):
        # Filtrage d'appartenance appliqué au niveau de la requête elle-même.
        return Order.objects.filter(user=self.request.user)
