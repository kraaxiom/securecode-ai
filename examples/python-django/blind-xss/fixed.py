# Correctif : Blind XSS (CWE-79) — on retire mark_safe() et on laisse le
# moteur de templates Django échapper automatiquement la donnée non fiable.
# Toute interface interne/admin doit traiter les données externes exactement
# comme les pages publiques : aucune confiance implicite.

from django.shortcuts import render
from .models import SupportTicket


def admin_ticket_detail(request, ticket_id):
    """Vue interne (back-office) affichant un ticket support."""
    ticket = SupportTicket.objects.get(pk=ticket_id)

    # Le message brut est transmis tel quel au template ; Django l'échappe
    # automatiquement à l'affichage ({{ ticket.message }}) — pas de |safe,
    # pas de mark_safe() sur une donnée d'origine externe.
    return render(request, "admin/ticket_detail.html", {
        "ticket": ticket,
    })
