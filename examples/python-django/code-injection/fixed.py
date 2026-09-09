# Correction : suppression totale de eval(). Les opérations autorisées
# sont déclarées dans une liste blanche de fonctions, l'entrée utilisateur
# ne sert plus qu'à sélectionner une opération, jamais à fournir du code.
from django.http import JsonResponse, HttpResponseBadRequest
from django.views.decorators.csrf import csrf_exempt

ALLOWED_OPERATIONS = {
    "add": lambda a, b: a + b,
    "sub": lambda a, b: a - b,
    "mul": lambda a, b: a * b,
}


@csrf_exempt
def compute_formula(request):
    op = request.POST.get("op")
    if op not in ALLOWED_OPERATIONS:
        return HttpResponseBadRequest("Opération non autorisée")

    try:
        a = float(request.POST.get("a"))
        b = float(request.POST.get("b"))
    except (TypeError, ValueError):
        return HttpResponseBadRequest("Paramètres invalides")

    result = ALLOWED_OPERATIONS[op](a, b)

    return JsonResponse({"result": result})
