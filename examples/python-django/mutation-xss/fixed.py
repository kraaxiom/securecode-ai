# Correctif : Mutation XSS / mXSS (CWE-79) — remplacement de la sanitisation
# maison par une bibliothèque activement maintenue (bleach, basé sur
# html5lib) qui traite correctement les cas de mutation liés au parsing
# HTML5, plutôt qu'un filtrage ad hoc de balises.

import bleach
from django.utils.safestring import mark_safe
from django.shortcuts import render
from .models import RichNote

ALLOWED_TAGS = ["p", "strong", "em", "u", "ul", "ol", "li", "br", "a"]
ALLOWED_ATTRS = {"a": ["href", "title", "rel"]}


def note_detail(request, note_id):
    """Affiche une note en HTML riche stockée par l'utilisateur (éditeur WYSIWYG)."""
    note = RichNote.objects.get(pk=note_id)

    # Sanitisation via une bibliothèque maintenue, tenant compte des cas de
    # mutation connus lors du reparsing navigateur.
    cleaned = bleach.clean(
        note.content,
        tags=ALLOWED_TAGS,
        attributes=ALLOWED_ATTRS,
        protocols=["http", "https"],
        strip=True,
    )

    return render(request, "notes/detail.html", {
        "note": note,
        "content_html": mark_safe(cleaned),
    })
