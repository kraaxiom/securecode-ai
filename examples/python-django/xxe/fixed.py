# Vue Django corrigée — CWE-611 (XML External Entity Injection)
# Correction : parseur lxml configuré explicitement pour désactiver la
# résolution d'entités, le chargement de DTD externe et l'accès réseau,
# ce qui neutralise les attaques XXE (divulgation de fichier, SSRF,
# "billion laughs").

from django.http import HttpResponse
from lxml import etree

# Sécurisé : DTD et entités externes désactivées, pas d'accès réseau
_SAFE_PARSER = etree.XMLParser(
    resolve_entities=False,
    no_network=True,
    dtd_validation=False,
    load_dtd=False,
)


def import_catalog(request):
    uploaded_file = request.FILES["catalog"]

    tree = etree.parse(uploaded_file, parser=_SAFE_PARSER)
    root = tree.getroot()

    return HttpResponse(f"Catalogue importé : {len(root)} éléments")
