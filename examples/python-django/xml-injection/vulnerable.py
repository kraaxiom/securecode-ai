# Vue Django vulnérable — CWE-91 (XML Injection)
# Le nom fourni par l'utilisateur est inséré par f-string directement
# dans un document XML sans échappement des caractères spéciaux
# ("<", ">", "&"), permettant à un attaquant d'altérer la structure
# du document (ajout de nœuds, modification de la logique métier).

from django.http import HttpResponse


def export_user_xml(request):
    name = request.POST.get("name", "")
    role = request.POST.get("role", "user")

    # Vulnérable : concaténation de chaînes pour construire le XML
    xml = f"<user><name>{name}</name><role>{role}</role></user>"

    return HttpResponse(xml, content_type="application/xml")
