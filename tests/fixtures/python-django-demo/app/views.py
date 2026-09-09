import os
from django.template import Template, Context
from django.http import HttpResponse, FileResponse
from django.views.decorators.csrf import csrf_exempt
import json

# Fixture volontairement vulnérable — NE PAS DÉPLOYER.
# Utilisée pour valider le pipeline scan -> detect -> patch -> test du skill SecureCode AI.

MEDIA_ROOT = '/var/app/media'


def greet(request):
    """
    Vulnérabilité : SSTI (voir knowledge/injections/ssti.md).
    Le nom fourni par l'utilisateur est injecté directement dans le code du template
    avant compilation, permettant l'exécution de code Django Template Language.
    """
    name = request.GET.get('name', 'invité')
    template = Template("Bonjour, " + name + " !")
    return HttpResponse(template.render(Context({})))


def download_file(request):
    """
    Vulnérabilité : Path Traversal (voir knowledge/file-inclusion/path-traversal.md).
    Le chemin fourni par le client est concaténé sans validation, permettant de sortir
    du répertoire MEDIA_ROOT avec des séquences '../'.
    """
    filename = request.GET.get('file')
    path = os.path.join(MEDIA_ROOT, filename)
    return FileResponse(open(path, 'rb'))


@csrf_exempt
def update_profile(request):
    """
    Vulnérabilité : Mass Assignment (voir knowledge/authorization/mass-assignment.md).
    Le corps JSON de la requête est appliqué tel quel sur le modèle utilisateur,
    permettant à un attaquant de modifier des champs sensibles (ex: is_admin).
    """
    data = json.loads(request.body)
    user = request.user
    for key, value in data.items():
        setattr(user, key, value)
    user.save()
    return HttpResponse('OK')
