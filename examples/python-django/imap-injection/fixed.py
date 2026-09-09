# Correction : le critère est nettoyé des guillemets et des caractères
# de contrôle, limité en longueur, et transmis comme argument séparé à
# search() plutôt que d'être interpolé dans une chaîne de commande IMAP.
import imaplib
from django.http import JsonResponse


def search_mailbox(request):
    criteria = request.GET.get("q", "")
    criteria = criteria.replace('"', "").replace("\r", "").replace("\n", "")[:200]

    mail = imaplib.IMAP4_SSL("imap.example.com")
    mail.login(request.user.email, request.session["imap_password"])
    mail.select("INBOX")

    typ, data = mail.search(None, "SUBJECT", criteria)

    return JsonResponse({"ids": data[0].decode().split()})
