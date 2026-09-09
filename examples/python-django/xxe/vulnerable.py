# Vue Django vulnérable — CWE-611 (XML External Entity Injection)
# Le fichier XML uploadé par l'utilisateur est analysé avec la
# configuration par défaut de lxml, qui ne désactive pas explicitement
# le traitement des DTD et des entités externes. Un attaquant peut
# déclarer une entité pointant vers un fichier local ou une URL interne.

from django.http import HttpResponse
from lxml import etree


def import_catalog(request):
    uploaded_file = request.FILES["catalog"]

    # Vulnérable : parseur par défaut, DTD et entités externes non désactivées
    tree = etree.parse(uploaded_file)
    root = tree.getroot()

    return HttpResponse(f"Catalogue importé : {len(root)} éléments")
