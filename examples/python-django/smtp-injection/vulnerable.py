# Vue Django vulnérable — CWE-93 (SMTP / Email Header Injection)
# Les champs "name" et "subject" fournis par l'utilisateur sont
# concaténés directement dans les en-têtes du message avant l'envoi.
# Un attaquant peut injecter des séquences CR/LF pour ajouter des
# en-têtes arbitraires (Cc, Bcc) ou un second corps de message.

import smtplib

from django.http import HttpResponse


def contact_form(request):
    name = request.POST.get("name", "")
    subject = request.POST.get("subject", "")
    body = request.POST.get("body", "")

    # Vulnérable : en-têtes construits par concaténation manuelle
    message = f"From: contact@example.com\r\nReply-To: {name}\r\nSubject: {subject}\r\n\r\n{body}"

    with smtplib.SMTP("localhost") as server:
        server.sendmail("contact@example.com", "dest@example.com", message)

    return HttpResponse("Message envoyé")
