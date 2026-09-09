# Faille : injection SQL booléenne (CWE-89).
# Le nom de produit est inséré par f-string dans une requête ORM brute
# (raw()), permettant de modifier la valeur de vérité de la clause WHERE.
from django.http import JsonResponse
from .models import Product


def search_products(request):
    name = request.GET.get("name", "")

    products = Product.objects.raw(
        f"SELECT * FROM shop_product WHERE name = '{name}'"
    )

    return JsonResponse({"results": [p.name for p in products]})
