# Correction : chaque cellule est passée dans neutraliser_formule(), qui
# préfixe d'une apostrophe toute valeur commençant par un caractère
# déclencheur de formule avant écriture dans le fichier exporté.
import csv
from django.http import HttpResponse
from .models import Ticket


def neutraliser_formule(valeur: str) -> str:
    if valeur and valeur[0] in ("=", "+", "-", "@", "\t", "\r"):
        return "'" + valeur
    return valeur


def export_tickets(request):
    response = HttpResponse(content_type="text/csv")
    response["Content-Disposition"] = 'attachment; filename="tickets.csv"'

    writer = csv.writer(response)
    writer.writerow(["Titre", "Description"])
    for ticket in Ticket.objects.all():
        writer.writerow([
            neutraliser_formule(ticket.title),
            neutraliser_formule(ticket.description),
        ])

    return response
