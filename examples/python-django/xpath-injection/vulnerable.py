# Vue Django vulnérable — CWE-643 (XPath Injection)
# Les identifiants "user" et "password" fournis par le client sont
# concaténés directement dans une expression XPath utilisée pour
# authentifier l'utilisateur contre un document XML. Un attaquant peut
# modifier la logique de sélection de nœuds (contournement d'auth).

from django.http import JsonResponse
from lxml import etree

USERS_XML_PATH = "/var/app/data/users.xml"


def xml_login(request):
    user = request.POST.get("user", "")
    password = request.POST.get("password", "")

    tree = etree.parse(USERS_XML_PATH)

    # Vulnérable : concaténation directe dans l'expression XPath
    expr = f"//user[username='{user}' and password='{password}']"
    nodes = tree.xpath(expr)

    return JsonResponse({"authenticated": len(nodes) > 0})
