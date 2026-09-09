# Remédiation — Price Manipulation (manipulation de prix)

## Principe
Toujours recalculer le prix final côté serveur à partir du catalogue et des règles de tarification, sans jamais faire confiance à une valeur envoyée par le client. Revérifier le montant au moment de la capture effective du paiement.

## PHP
```php
// Avant — vulnérable : le prix vient directement du client
public function checkout(Request $request) {
    $amount = $request->input('price'); // falsifiable
    PaymentGateway::charge($request->user(), $amount);
}

// Après — sécurisé
public function checkout(Request $request) {
    $cart = Cart::where('user_id', $request->user()->id)->firstOrFail();
    $amount = 0;
    foreach ($cart->items as $item) {
        $product = Product::findOrFail($item->product_id); // prix depuis le catalogue serveur
        $amount += $product->price * $item->quantity;
    }
    $amount = apply_discount_rules($amount, $request->user()); // recalculé serveur
    PaymentGateway::charge($request->user(), $amount);
}
```

## JS / Node
```js
// Avant — vulnérable
app.post('/checkout', async (req, res) => {
  await chargeCard(req.user, req.body.amount); // montant fourni par le client
  res.sendStatus(200);
});

// Après — sécurisé
app.post('/checkout', async (req, res) => {
  const cart = await getCart(req.user.id);
  const amount = cart.items.reduce((total, item) => {
    const product = catalog.getPrice(item.productId); // source serveur, pas le client
    return total + product.price * item.quantity;
  }, 0);
  const finalAmount = applyDiscountRules(amount, req.user);
  await chargeCard(req.user, finalAmount);
  res.sendStatus(200);
});
```

## Python
```python
# Avant — vulnérable
@app.route('/checkout', methods=['POST'])
def checkout():
    amount = request.json['amount']  # falsifiable
    charge_card(current_user, amount)
    return '', 200

# Après — sécurisé
@app.route('/checkout', methods=['POST'])
def checkout():
    cart = get_cart(current_user.id)
    amount = sum(
        catalog.get_price(item.product_id) * item.quantity for item in cart.items
    )  # recalculé depuis le catalogue serveur
    amount = apply_discount_rules(amount, current_user)
    charge_card(current_user, amount)
    return '', 200
```

## Checklist de vérification post-patch
- [ ] Aucun paramètre `price`/`amount`/`discount` en provenance du client n'est utilisé directement pour la facturation.
- [ ] Le prix final est recalculé côté serveur à partir du catalogue au moment du checkout ET revérifié à la capture du paiement.
- [ ] Le moteur de tarification est centralisé et partagé par tous les clients (web, mobile, partenaires).
- [ ] Les écarts entre prix catalogue et prix soumis sont journalisés pour détecter les tentatives de manipulation.
