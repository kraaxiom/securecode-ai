# Faille : injection de langage de template (CWE-917, proche EL injection).
# Le nom fourni par l'utilisateur est concaténé dans la SOURCE du template
# Django avant compilation, permettant l'exécution de tags/variables
# arbitraires du moteur de template.
from django.http import HttpResponse
from django.template import Template, Context


def greet(request):
    nom = request.GET.get("nom", "")

    template_source = "<p>Bonjour " + nom + "</p>"
    template = Template(template_source)

    return HttpResponse(template.render(Context({})))
