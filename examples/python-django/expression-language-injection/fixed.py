# Correction : le template est une chaîne statique définie par le
# développeur. La donnée utilisateur est transmise uniquement comme
# variable de contexte, jamais insérée dans la source du template évaluée.
from django.http import HttpResponse
from django.template import Template, Context


def greet(request):
    nom = request.GET.get("nom", "")

    template_source = "<p>Bonjour {{ nom }}</p>"
    template = Template(template_source)

    return HttpResponse(template.render(Context({"nom": nom})))
