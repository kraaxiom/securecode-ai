# Vue Django corrigée — CWE-1336 (Server-Side Template Injection)
# Correction : le template est une chaîne statique fixée par le
# développeur ; la donnée utilisateur transite uniquement via le
# contexte de rendu ({{ name }}), jamais dans la structure du template.

from django.http import HttpResponse
from django.template import Context, Engine

engine = Engine.get_default()

# Sécurisé : template statique, jamais construit à partir d'une entrée utilisateur
_TEMPLATE = engine.from_string("Bonjour {{ name }}, bienvenue !")


def welcome(request):
    name = request.GET.get("name", "")

    return HttpResponse(_TEMPLATE.render(Context({"name": name})))
