# Vue Django corrigée — CWE-643 (XPath Injection)
# Correction : utilisation de variables XPath liées (supportées par
# lxml), qui séparent l'expression XPath statique des valeurs
# utilisateur. À terme, XPath ne devrait pas servir de mécanisme
# d'authentification (préférer une base avec hachage de mot de passe).

from django.http import JsonResponse
from lxml import etree

USERS_XML_PATH = "/var/app/data/users.xml"


def xml_login(request):
    user = request.POST.get("user", "")
    password = request.POST.get("password", "")

    tree = etree.parse(USERS_XML_PATH)

    # Sécurisé : variables XPath liées, pas de concaténation
    nodes = tree.xpath(
        "//user[username=$u and password=$p]",
        u=user,
        p=password,
    )

    return JsonResponse({"authenticated": len(nodes) > 0})
