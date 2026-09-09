# Correction : utilisation de l'ORM Django (QuerySet.filter) qui génère
# une requête paramétrée en interne, supprimant toute concaténation SQL
# et donc toute possibilité de modifier la logique booléenne du filtre.
from django.http import JsonResponse
from .models import Product


def search_products(request):
    name = request.GET.get("name", "")

    products = Product.objects.filter(name=name)

    return JsonResponse({"results": [p.name for p in products]})
