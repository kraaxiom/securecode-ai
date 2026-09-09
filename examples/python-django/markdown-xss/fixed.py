# Correctif : Markdown XSS (CWE-79) — le HTML brut n'est plus autorisé dans
# le Markdown, et le HTML généré passe par bleach avant d'être marqué sûr,
# avec une liste blanche stricte de balises/attributs.

import bleach
import markdown
from django.utils.safestring import mark_safe
from django.shortcuts import render
from .models import Article

ALLOWED_TAGS = ["p", "strong", "em", "ul", "ol", "li", "a", "code", "pre", "blockquote"]
ALLOWED_ATTRS = {"a": ["href", "title", "rel"]}


def article_detail(request, slug):
    """Affiche un article rédigé en Markdown par un utilisateur."""
    article = Article.objects.get(slug=slug)

    # extension "extra" seulement, sans md_in_html : pas de passthrough HTML.
    rendered_html = markdown.markdown(article.body, extensions=["extra"])

    # Sanitisation stricte du HTML généré avant affichage : liste blanche
    # de balises/attributs, schémas d'URL restreints.
    clean_html = bleach.clean(
        rendered_html,
        tags=ALLOWED_TAGS,
        attributes=ALLOWED_ATTRS,
        protocols=["http", "https", "mailto"],
        strip=True,
    )

    return render(request, "articles/detail.html", {
        "article": article,
        "rendered_html": mark_safe(clean_html),
    })
