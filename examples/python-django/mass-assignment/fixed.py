# Corrigé — Mass Assignment (CWE-915)
# Un serializer d'entrée dédié restreint explicitement les champs
# modifiables par le client à une liste blanche ('name', 'email'). Les
# attributs sensibles ('role', 'is_staff') sont exclus et ne peuvent plus
# transiter par cette route.

from rest_framework import serializers, generics
from .models import User


class UserUpdateSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ["name", "email"]  # 'role' et 'is_staff' explicitement exclus


class UserUpdateView(generics.UpdateAPIView):
    queryset = User.objects.all()
    serializer_class = UserUpdateSerializer
