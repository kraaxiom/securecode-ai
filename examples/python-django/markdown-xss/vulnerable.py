# Faille : Markdown XSS (CWE-79) — le contenu Markdown soumis par
# l'utilisateur est converti en HTML avec l'extension autorisant le HTML
# brut, puis marqué comme sûr sans aucune sanitisation avant affichage.

import markdown
from django.utils.safestring import mark_safe
from django.shortcuts import render
from .models import Article


def article_detail(request, slug):
    """Affiche un article rédigé en Markdown par un utilisateur."""
    article = Article.objects.get(slug=slug)

    # extensions=['md_in_html'] + safe_mode implicite désactivé : tout HTML
    # brut présent dans le Markdown (ex. <script>, <img onerror=...>) est
    # conservé tel quel puis marqué "safe" pour le template.
    rendered_html = markdown.markdown(article.body, extensions=["md_in_html", "extra"])
    rendered_html = mark_safe(rendered_html)

    return render(request, "articles/detail.html", {
        "article": article,
        "rendered_html": rendered_html,
    })
