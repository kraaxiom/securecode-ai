# Vulnérable — Mass Assignment (CWE-915)
# Le serializer expose tous les champs du modèle User, y compris 'role' et
# 'is_staff'. Un attaquant peut ainsi injecter ces champs dans le corps de
# la requête de mise à jour de profil pour s'auto-attribuer des privilèges.

from rest_framework import serializers, generics
from .models import User


class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = "__all__"  # inclut 'role', 'is_staff', etc.


class UserUpdateView(generics.UpdateAPIView):
    queryset = User.objects.all()
    serializer_class = UserSerializer
