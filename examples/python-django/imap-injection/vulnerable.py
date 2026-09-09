# Faille : injection IMAP (CWE-93).
# Le critère de recherche fourni par l'utilisateur est concaténé
# directement dans la commande IMAP SEARCH, permettant d'injecter des
# guillemets ou des mots-clés IMAP supplémentaires.
import imaplib
from django.http import JsonResponse


def search_mailbox(request):
    criteria = request.GET.get("q", "")

    mail = imaplib.IMAP4_SSL("imap.example.com")
    mail.login(request.user.email, request.session["imap_password"])
    mail.select("INBOX")

    typ, data = mail.search(None, f'SUBJECT "{criteria}"')

    return JsonResponse({"ids": data[0].decode().split()})
