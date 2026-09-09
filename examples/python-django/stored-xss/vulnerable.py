# Faille : Stored XSS (CWE-79) — le corps d'un commentaire, persisté en
# base de données, est réaffiché via mark_safe() sans échappement ni
# sanitisation. Le payload injecté une fois s'exécute pour chaque visiteur
# consultant la page.

from django.utils.safestring import mark_safe
from django.shortcuts import render
from .models import Comment


def post_detail(request, post_id):
    """Affiche un article et ses commentaires stockés en base."""
    comments = Comment.objects.filter(post_id=post_id)

    rendered_comments = [
        mark_safe(f"<div class='comment'>{c.body}</div>") for c in comments
    ]

    return render(request, "blog/post_detail.html", {
        "post_id": post_id,
        "rendered_comments": rendered_comments,
    })
