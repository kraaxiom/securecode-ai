# Faille : injection d'en-tête HTTP (CWE-113).
# Le nom de fichier fourni par l'utilisateur est concaténé directement
# dans l'en-tête Content-Disposition, sans filtrage des caractères \r\n.
from django.http import HttpResponse


def download_report(request):
    filename = request.GET.get("filename", "report.pdf")

    response = HttpResponse(b"%PDF-1.4 ...", content_type="application/pdf")
    response["Content-Disposition"] = f"attachment; filename={filename}"

    return response
