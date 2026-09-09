# Vue Django corrigée — CWE-97 (Server-Side Includes Injection)
# Correction : échappement HTML du commentaire et neutralisation de
# toute séquence de directive SSI résiduelle ("<!--#") avant écriture.
# Préférable encore : rendre le contenu via le moteur de templates
# Django (Jinja2/DTL) plutôt que d'écrire un fichier .shtml interprété.

import html
from pathlib import Path

from django.http import HttpResponse

SHTML_PATH = Path("/var/www/public/comments.shtml")


def strip_ssi_directives(value: str) -> str:
    return value.replace("<!--#", "&lt;!--#")


def add_comment(request):
    comment = request.POST.get("comment", "")

    # Sécurisé : échappement HTML puis neutralisation des directives SSI
    safe_comment = strip_ssi_directives(html.escape(comment))

    with open(SHTML_PATH, "a", encoding="utf-8") as f:
        f.write(f"<p>{safe_comment}</p>")

    return HttpResponse("Commentaire ajouté")
