# Correction : suppression des caractères de contrôle \r\n et réduction
# du nom de fichier à son composant de base (pas de chemin), avant
# insertion dans l'en-tête, en défense en profondeur.
import os
import re
from django.http import HttpResponse


def download_report(request):
    filename = request.GET.get("filename", "report.pdf")
    filename = re.sub(r"[\r\n]", "", filename)
    filename = os.path.basename(filename)

    response = HttpResponse(b"%PDF-1.4 ...", content_type="application/pdf")
    response["Content-Disposition"] = f'attachment; filename="{filename}"'

    return response
