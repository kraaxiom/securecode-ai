# Remédiation — Business Rule Bypass (contournement de règle métier)

## Principe
Ne jamais considérer une règle métier comme appliquée tant qu'elle n'a pas été revalidée côté serveur, indépendamment de ce que fait l'interface. Ne jamais faire confiance à un paramètre sensible (prix, rôle, statut) transmis par le client ; le recalculer ou le vérifier contre une source serveur faisant autorité.

## PHP
```php
// Avant — vulnérable : l'âge minimum n'est vérifié qu'en JS, le serveur fait confiance au client
public function subscribe(Request $request) {
    // Aucune revérification de l'âge côté serveur
    Subscription::create(['user_id' => $request->user()->id, 'plan' => $request->plan]);
}

// Après — sécurisé
public function subscribe(Request $request) {
    $user = $request->user();
    if ($user->birthdate === null || $user->age() < 18) {
        abort(403, 'Âge minimum non satisfait.');
    }
    $plan = Plan::findOrFail($request->integer('plan_id')); // dérivé serveur, pas du prix client
    Subscription::create(['user_id' => $user->id, 'plan_id' => $plan->id, 'price' => $plan->price]);
}
```

## JS / Node
```js
// Avant — vulnérable : le rôle est accepté tel quel depuis le corps de la requête
app.post('/api/users/:id/role', (req, res) => {
  db.updateUserRole(req.params.id, req.body.role); // rôle non revalidé
  res.sendStatus(200);
});

// Après — sécurisé
const ALLOWED_ROLES = ['member', 'viewer']; // un admin ne peut jamais s'auto-attribuer via cet endpoint
app.post('/api/users/:id/role', requireAdmin, (req, res) => {
  const { role } = req.body;
  if (!ALLOWED_ROLES.includes(role)) {
    return res.status(400).json({ error: 'Rôle non autorisé' });
  }
  db.updateUserRole(req.params.id, role);
  res.sendStatus(200);
});
```

## Python
```python
# Avant — vulnérable : statut de commande accepté du client sans machine à états serveur
@app.route('/orders/<int:order_id>/status', methods=['POST'])
def set_status(order_id):
    db.execute("UPDATE orders SET status = %s WHERE id = %s", [request.json['status'], order_id])
    return '', 200

# Après — sécurisé
VALID_TRANSITIONS = {
    'pending': {'paid', 'cancelled'},
    'paid': {'shipped'},
    'shipped': {'delivered'},
}

@app.route('/orders/<int:order_id>/status', methods=['POST'])
def set_status(order_id):
    order = get_order(order_id)
    new_status = request.json['status']
    if new_status not in VALID_TRANSITIONS.get(order.status, set()):
        return {'error': 'Transition invalide'}, 400
    db.execute("UPDATE orders SET status = %s WHERE id = %s", [new_status, order_id])
    return '', 200
```

## Checklist de vérification post-patch
- [ ] Chaque règle métier critique validée en JS possède un équivalent serveur strictement appliqué.
- [ ] Les valeurs sensibles (prix, rôle, statut) sont recalculées ou vérifiées côté serveur, jamais acceptées telles quelles depuis le client.
- [ ] Les processus multi-étapes reposent sur une machine à états serveur qui rejette les transitions invalides.
- [ ] La règle métier est appliquée de façon identique sur toutes les APIs consommant le même flux (web, mobile, partenaire).
