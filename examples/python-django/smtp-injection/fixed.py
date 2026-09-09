# Vue Django corrigée — CWE-93 (SMTP / Email Header Injection)
# Correction : utilisation de la fonction send_mail de Django (qui
# échappe correctement les en-têtes) et suppression explicite des
# caractères de contrôle CR/LF dans les champs libres avant usage.

from django.core.mail import send_mail
from django.http import HttpResponse


def strip_control_chars(value: str) -> str:
    return value.replace("\r", "").replace("\n", "")


def contact_form(request):
    name = strip_control_chars(request.POST.get("name", ""))
    subject = strip_control_chars(request.POST.get("subject", ""))
    body = request.POST.get("body", "")

    # Sécurisé : send_mail échappe/plie correctement les en-têtes
    send_mail(
        subject=subject,
        message=f"De : {name}\n\n{body}",
        from_email="contact@example.com",
        recipient_list=["dest@example.com"],
    )

    return HttpResponse("Message envoyé")
