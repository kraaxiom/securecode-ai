# Correction : toute cellule dont le premier caractère est un déclencheur
# de formule (=, +, -, @, tabulation, retour chariot) est préfixée par une
# apostrophe, neutralisant son interprétation comme formule par le tableur.
import csv
from django.http import HttpResponse
from .models import Feedback

FORMULA_TRIGGERS = ("=", "+", "-", "@", "\t", "\r")


def sanitize_csv_cell(value: str) -> str:
    if value and value[0] in FORMULA_TRIGGERS:
        return "'" + value
    return value


def export_feedback_csv(request):
    response = HttpResponse(content_type="text/csv")
    response["Content-Disposition"] = 'attachment; filename="feedback.csv"'

    writer = csv.writer(response)
    for feedback in Feedback.objects.all():
        writer.writerow([
            sanitize_csv_cell(feedback.author_name),
            sanitize_csv_cell(feedback.comment),
        ])

    return response
