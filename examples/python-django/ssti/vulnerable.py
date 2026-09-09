# Vue Django vulnérable — CWE-1336 (Server-Side Template Injection)
# Le nom fourni par l'utilisateur est concaténé au texte même du
# template avant compilation par le moteur (via Engine.from_string),
# au lieu d'être passé comme variable de contexte. Le moteur interprète
# alors la syntaxe injectée comme du code de template.

from django.http import HttpResponse
from django.template import Engine

engine = Engine.get_default()


def welcome(request):
    name = request.GET.get("name", "")

    # Vulnérable : le texte du template est construit dynamiquement
    template_source = "Bonjour " + name + ", bienvenue !"
    template = engine.from_string(template_source)

    return HttpResponse(template.render())
