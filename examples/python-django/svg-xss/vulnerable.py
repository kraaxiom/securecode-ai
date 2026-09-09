# Faille : XSS via SVG (CWE-79) — un fichier SVG uploadé par l'utilisateur
# est stocké et servi tel quel, avec un Content-Type image/svg+xml et en
# affichage inline. Un SVG peut contenir <script> ou des gestionnaires
# d'événements exécutés par le navigateur au rendu.

from django.http import FileResponse
from django.shortcuts import get_object_or_404
from .models import UserAvatar


def upload_avatar(request):
    """Réception d'un avatar SVG et stockage brut sur le disque."""
    uploaded_file = request.FILES["avatar"]
    avatar = UserAvatar.objects.create(
        user=request.user,
        file=uploaded_file,  # aucune sanitisation du contenu SVG
    )
    return avatar


def serve_avatar(request, avatar_id):
    """Sert le SVG stocké, affiché inline dans le même domaine que l'app."""
    avatar = get_object_or_404(UserAvatar, pk=avatar_id)
    # Servi depuis le même domaine, en affichage inline (pas de
    # Content-Disposition: attachment) : tout script embarqué s'exécute
    # dans le contexte de sécurité de l'application.
    return FileResponse(avatar.file.open("rb"), content_type="image/svg+xml")
