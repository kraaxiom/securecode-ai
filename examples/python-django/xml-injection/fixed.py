# Vue Django corrigée — CWE-91 (XML Injection)
# Correction : le document XML est construit via lxml.etree, qui
# échappe automatiquement le contenu des nœuds et attributs, plutôt
# que par concaténation de chaînes.

from django.http import HttpResponse
from lxml import etree


def export_user_xml(request):
    name = request.POST.get("name", "")
    role = request.POST.get("role", "user")

    # Sécurisé : échappement automatique par lxml
    user_el = etree.Element("user")
    name_el = etree.SubElement(user_el, "name")
    name_el.text = name
    role_el = etree.SubElement(user_el, "role")
    role_el.text = role

    xml = etree.tostring(user_el)

    return HttpResponse(xml, content_type="application/xml")
