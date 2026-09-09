# Faille : Blind XSS (CWE-79) — la donnée soumise par un utilisateur externe
# (ex. formulaire de contact / ticket support) est réaffichée sans échappement
# dans l'interface d'administration interne via mark_safe(). L'attaquant ne
# voit jamais l'exécution : le payload se déclenche dans le navigateur de
# l'agent support qui consulte le back-office.

from django.utils.safestring import mark_safe
from django.shortcuts import render
from .models import SupportTicket


def admin_ticket_detail(request, ticket_id):
    """Vue interne (back-office) affichant un ticket support."""
    ticket = SupportTicket.objects.get(pk=ticket_id)

    # Le champ 'message' provient d'un utilisateur non authentifié via le
    # formulaire public de contact — il est traité comme fiable ici, à tort.
    ticket_html = mark_safe(f"<div class='ticket-message'>{ticket.message}</div>")

    return render(request, "admin/ticket_detail.html", {
        "ticket": ticket,
        "ticket_html": ticket_html,
    })
