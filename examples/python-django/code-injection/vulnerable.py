# Faille : injection de code (CWE-94).
# La formule fournie par l'utilisateur est évaluée directement avec eval(),
# donnant à l'attaquant un contrôle total sur l'interpréteur Python.
from django.http import JsonResponse, HttpResponseBadRequest
from django.views.decorators.csrf import csrf_exempt


@csrf_exempt
def compute_formula(request):
    formula = request.POST.get("formula")
    if not formula:
        return HttpResponseBadRequest("formula manquante")

    result = eval(formula)

    return JsonResponse({"result": result})
