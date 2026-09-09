# Faille : injection de formule dans un export tableur (CWE-1236).
# Les commentaires utilisateur sont écrits sans neutralisation dans le
# fichier CSV exporté vers Excel/LibreOffice.
import csv
from django.http import HttpResponse
from .models import Ticket


def export_tickets(request):
    response = HttpResponse(content_type="text/csv")
    response["Content-Disposition"] = 'attachment; filename="tickets.csv"'

    writer = csv.writer(response)
    writer.writerow(["Titre", "Description"])
    for ticket in Ticket.objects.all():
        writer.writerow([ticket.title, ticket.description])

    return response
