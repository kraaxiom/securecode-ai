# Faille : Mutation XSS / mXSS (CWE-79) — le contenu riche est sanitisé
# "à la main" via un aller-retour chaîne -> parsing -> ré-sérialisation
# maison, sans bibliothèque de sanitisation maintenue tenant compte des
# quirks de reparsing du navigateur. Le HTML final peut "muter" une fois
# réinterprété par le client.

from lxml import html as lxml_html
from django.utils.safestring import mark_safe
from django.shortcuts import render
from .models import RichNote


def note_detail(request, note_id):
    """Affiche une note en HTML riche stockée par l'utilisateur (éditeur WYSIWYG)."""
    note = RichNote.objects.get(pk=note_id)

    # Sanitisation "naïve" maison : suppression de quelques balises connues
    # seulement, puis ré-sérialisation — ne couvre pas les vecteurs de
    # mutation liés au reparsing navigateur (ex. balises imbriquées mal
    # formées qui se reconstituent différemment côté client).
    tree = lxml_html.fromstring(note.content)
    for bad_tag in tree.xpath("//script | //style"):
        bad_tag.drop_tree()
    cleaned = lxml_html.tostring(tree, encoding="unicode")

    return render(request, "notes/detail.html", {
        "note": note,
        "content_html": mark_safe(cleaned),
    })
