# Vue Django vulnérable — CWE-97 (Server-Side Includes Injection)
# Le commentaire utilisateur est écrit tel quel dans un fichier .shtml
# servi par un serveur web avec SSI activé (ex: Apache mod_include).
# Un attaquant peut injecter une directive "<!--#exec cmd=...-->".

from pathlib import Path

from django.http import HttpResponse

SHTML_PATH = Path("/var/www/public/comments.shtml")


def add_comment(request):
    comment = request.POST.get("comment", "")

    # Vulnérable : contenu utilisateur écrit sans échappement dans un
    # fichier interprété par SSI côté serveur web
    with open(SHTML_PATH, "a", encoding="utf-8") as f:
        f.write(f"<p>{comment}</p>")

    return HttpResponse("Commentaire ajouté")
