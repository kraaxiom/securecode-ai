# Faille : pollution de paramètres HTTP (CWE-235).
# request.GET.get("role") ne renvoie silencieusement que la première
# valeur en cas de paramètre dupliqué (?role=user&role=admin), masquant
# une éventuelle tentative de contournement d'un contrôle en amont (WAF).
from django.http import JsonResponse


def assign_role(request):
    role = request.GET.get("role", "user")

    # La logique métier fait confiance à une seule valeur sans vérifier
    # si le paramètre a été envoyé plusieurs fois.
    user = request.user
    user.role = role
    user.save()

    return JsonResponse({"role": role})
