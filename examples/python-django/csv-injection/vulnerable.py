# Faille : injection CSV / formule (CWE-1236).
# Les champs utilisateur (nom, commentaire) sont écrits tels quels dans le
# fichier CSV exporté, sans neutraliser les caractères déclencheurs de
# formule (=, +, -, @) lorsqu'ils sont ouverts dans un tableur.
import csv
from django.http import HttpResponse
from .models import Feedback


def export_feedback_csv(request):
    response = HttpResponse(content_type="text/csv")
    response["Content-Disposition"] = 'attachment; filename="feedback.csv"'

    writer = csv.writer(response)
    for feedback in Feedback.objects.all():
        writer.writerow([feedback.author_name, feedback.comment])

    return response
