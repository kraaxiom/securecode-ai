# Correctif : Stored XSS (CWE-79) — suppression de mark_safe() ; les
# commentaires sont transmis bruts au template qui applique l'échappement
# automatique de Django à l'affichage.

from django.shortcuts import render
from .models import Comment


def post_detail(request, post_id):
    """Affiche un article et ses commentaires stockés en base."""
    comments = Comment.objects.filter(post_id=post_id)

    # Le template affiche {{ comment.body }} pour chaque commentaire ;
    # Django échappe automatiquement le contenu, aucun mark_safe/|safe
    # n'est appliqué sur cette donnée d'origine utilisateur.
    return render(request, "blog/post_detail.html", {
        "post_id": post_id,
        "comments": comments,
    })
