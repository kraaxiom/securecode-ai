# Faille : injection LDAP (CWE-90).
# L'identifiant utilisateur est concaténé directement dans le filtre de
# recherche LDAP, permettant de modifier la logique du filtre (ex:
# contournement d'authentification via wildcards).
import ldap
from django.http import JsonResponse


def find_user(request):
    uid = request.GET.get("uid", "")

    conn = ldap.initialize("ldap://directory.example.com")
    filter_str = f"(uid={uid})"

    results = conn.search_s(
        "dc=example,dc=com", ldap.SCOPE_SUBTREE, filter_str
    )

    return JsonResponse({"found": len(results) > 0})
