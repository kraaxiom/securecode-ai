# Correction : QueryDict.getlist() récupère toutes les occurrences du
# paramètre. Toute duplication d'un paramètre censé être unique est
# explicitement rejetée avant d'être utilisée dans la logique métier.
from django.http import JsonResponse, HttpResponseBadRequest


def assign_role(request):
    values = request.GET.getlist("role")

    if len(values) > 1:
        return HttpResponseBadRequest("Paramètre dupliqué non autorisé")

    role = values[0] if values else "user"

    user = request.user
    user.role = role
    user.save()

    return JsonResponse({"role": role})
