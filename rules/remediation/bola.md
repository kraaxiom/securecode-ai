# Remédiation — Broken Object Level Authorization (BOLA)

## Principe
Vérifier systématiquement, au niveau de chaque endpoint/resolver manipulant un objet par ID, que l'appelant est autorisé sur cet objet précis, en filtrant l'appartenance directement dans la requête de données.

## PHP (Laravel)
```php
// Avant — vulnérable : l'objet est récupéré par ID sans vérifier son propriétaire
public function show($id)
{
    $order = Order::findOrFail($id);
    return response()->json($order);
}

// Après — sécurisé : filtrage d'appartenance directement dans la requête
public function show($id)
{
    $order = Order::where('id', $id)
        ->where('user_id', auth()->id())
        ->firstOrFail();
    return response()->json($order);
}
```

## JS / Node.js (Express + GraphQL resolver)
```js
// Avant — vulnérable : REST et resolver GraphQL sans vérification d'appartenance
app.get('/api/orders/:id', requireAuth, async (req, res) => {
  const order = await Order.findById(req.params.id);
  res.json(order);
});

// Après — sécurisé : filtrage d'appartenance appliqué à la requête
app.get('/api/orders/:id', requireAuth, async (req, res) => {
  const order = await Order.findOne({ _id: req.params.id, ownerId: req.user.id });
  if (!order) return res.status(404).json({ error: 'Introuvable' });
  res.json(order);
});

// Resolver GraphQL — vérification explicite dans le resolver
const resolvers = {
  Query: {
    order: async (_parent, { id }, ctx) => {
      const order = await Order.findOne({ _id: id, ownerId: ctx.user.id });
      if (!order) throw new ForbiddenError('Accès refusé');
      return order;
    },
  },
};
```

## Python (Django REST Framework)
```python
# Avant — vulnérable : queryset non filtré par utilisateur
class OrderViewSet(viewsets.ModelViewSet):
    queryset = Order.objects.all()
    permission_classes = [IsAuthenticated]

# Après — sécurisé : queryset filtré par propriétaire au niveau de la requête
class OrderViewSet(viewsets.ModelViewSet):
    permission_classes = [IsAuthenticated]

    def get_queryset(self):
        return Order.objects.filter(user=self.request.user)
```

## Checklist de vérification post-patch
- [ ] Toute récupération/modification d'objet par ID inclut une clause de filtrage sur le propriétaire/tenant dans la requête elle-même.
- [ ] Les resolvers GraphQL vérifient l'autorisation dans le resolver, pas seulement via l'authentification globale.
- [ ] Un test d'accès croisé entre deux comptes distincts confirme qu'un objet d'autrui renvoie 403/404, jamais les données.
- [ ] L'isolation multi-tenant ne repose pas uniquement sur un filtre côté client.
- [ ] Le filtrage d'appartenance est appliqué au niveau de la requête de données, pas en post-traitement applicatif.
