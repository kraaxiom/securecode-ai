# Remédiation — Broken Object Level Authorization (BOLA)

## Principe
Ne jamais faire confiance à un identifiant d'objet fourni par le client. Vérifier systématiquement, à chaque accès, que l'utilisateur authentifié est propriétaire ou explicitement autorisé sur la ressource ciblée — pas seulement authentifié.

## PHP (Laravel / Eloquent)
```php
// Avant — vulnérable
Route::get('/api/orders/{id}', function ($id) {
    return Order::find($id);
});

// Après — sécurisé
Route::get('/api/orders/{id}', function ($id) {
    $order = Order::where('id', $id)
        ->where('user_id', auth()->id())
        ->firstOrFail();
    return $order;
});
```

## PHP (PDO natif)
```php
// Avant — vulnérable
$id = $_GET['id'];
$stmt = $pdo->prepare("SELECT * FROM invoices WHERE id = :id");
$stmt->execute(['id' => $id]);

// Après — sécurisé
$id = (int) $_GET['id'];
$stmt = $pdo->prepare("SELECT * FROM invoices WHERE id = :id AND user_id = :uid");
$stmt->execute(['id' => $id, 'uid' => $_SESSION['user_id']]);
```

## Node.js (Express)
```js
// Avant — vulnérable
app.get('/api/orders/:id', async (req, res) => {
  const order = await Order.findById(req.params.id);
  res.json(order);
});

// Après — sécurisé
app.get('/api/orders/:id', requireAuth, async (req, res) => {
  const order = await Order.findOne({ _id: req.params.id, userId: req.user.id });
  if (!order) return res.status(404).json({ error: 'Not found' });
  res.json(order);
});
```

## Python (Django REST Framework)
```python
# Avant — vulnérable
class OrderDetail(APIView):
    def get(self, request, pk):
        order = Order.objects.get(pk=pk)
        return Response(OrderSerializer(order).data)

# Après — sécurisé
class OrderDetail(APIView):
    permission_classes = [IsAuthenticated]

    def get(self, request, pk):
        order = get_object_or_404(Order, pk=pk, user=request.user)
        return Response(OrderSerializer(order).data)
```

## Checklist de vérification post-patch
- [ ] Chaque endpoint accédant à un objet par identifiant filtre explicitement sur l'utilisateur/tenant courant.
- [ ] La logique d'autorisation est centralisée (policy/guard) plutôt que dupliquée dans chaque contrôleur.
- [ ] Un test confirme qu'un utilisateur A ne peut pas accéder à un objet appartenant à l'utilisateur B (test croisé).
- [ ] Les identifiants séquentiels prévisibles sont, si possible, remplacés par des UUID en complément (pas en substitut) du contrôle d'autorisation.
- [ ] Les opérations d'écriture (`PUT`/`PATCH`/`DELETE`) appliquent le même contrôle que la lecture.
