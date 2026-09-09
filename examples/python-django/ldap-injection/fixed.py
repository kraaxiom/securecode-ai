# Correction : l'identifiant utilisateur est échappé via la fonction
# dédiée du module python-ldap, qui neutralise les caractères spéciaux
# de filtre (*, (, ), \, NUL) avant construction de la requête.
import ldap
from ldap.filter import escape_filter_chars
from django.http import JsonResponse


def find_user(request):
    uid = escape_filter_chars(request.GET.get("uid", ""))

    conn = ldap.initialize("ldap://directory.example.com")
    filter_str = f"(uid={uid})"

    results = conn.search_s(
        "dc=example,dc=com", ldap.SCOPE_SUBTREE, filter_str
    )

    return JsonResponse({"found": len(results) > 0})
