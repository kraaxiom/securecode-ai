# Vulnérable — Broken Object Level Authorization (CWE-639)
# Le queryset retourne toutes les commandes sans filtrage par propriétaire.
# Un utilisateur authentifié peut donc consulter/modifier la commande de
# n'importe quel autre utilisateur simplement en changeant l'ID dans l'URL.

from rest_framework import viewsets
from rest_framework.permissions import IsAuthenticated
from .models import Order
from .serializers import OrderSerializer


class OrderViewSet(viewsets.ModelViewSet):
    """API de gestion des commandes."""

    queryset = Order.objects.all()  # non filtré par utilisateur
    serializer_class = OrderSerializer
    permission_classes = [IsAuthenticated]
