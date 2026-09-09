# Remédiation — Workflow Bypass (contournement de séquence métier)

## Principe
Modéliser le processus comme une machine à états côté serveur, chaque transition étant validée contre l'état courant réel de la ressource. Ne jamais déterminer la progression d'un workflow à partir de données fournies par le client.

## PHP
```php
// Avant — vulnérable : l'étape d'activation est accessible sans vérifier les étapes précédentes
public function activateAccount(Request $request) {
    User::find($request->user()->id)->update(['status' => 'active']);
}

// Après — sécurisé
private const TRANSITIONS = [
    'registered' => ['kyc_pending'],
    'kyc_pending' => ['kyc_verified'],
    'kyc_verified' => ['active'],
];

public function activateAccount(Request $request) {
    $user = $request->user();
    if (!in_array('active', self::TRANSITIONS[$user->status] ?? [], true)) {
        abort(409, 'Étapes préalables non complétées.');
    }
    $user->update(['status' => 'active']);
    AuditLog::record($user->id, 'workflow_transition', ['to' => 'active']);
}
```

## JS / Node
```js
// Avant — vulnérable
app.post('/order/:id/confirm', async (req, res) => {
  await db.updateOrderStatus(req.params.id, 'confirmed'); // pas de vérification d'étape
  res.sendStatus(200);
});

// Après — sécurisé
const TRANSITIONS = { cart: ['shipping'], shipping: ['payment'], payment: ['confirmed'] };

app.post('/order/:id/confirm', async (req, res) => {
  const order = await db.getOrder(req.params.id);
  if (!TRANSITIONS[order.status]?.includes('confirmed')) {
    return res.status(409).json({ error: 'Étape précédente non complétée' });
  }
  await db.updateOrderStatus(order.id, 'confirmed');
  await db.logWorkflowTransition(order.id, 'confirmed');
  res.sendStatus(200);
});
```

## Python
```python
# Avant — vulnérable
@app.route('/kyc/activate', methods=['POST'])
def activate():
    db.execute("UPDATE users SET status = 'active' WHERE id = %s", [current_user.id])
    return '', 200

# Après — sécurisé
TRANSITIONS = {
    'registered': {'kyc_pending'},
    'kyc_pending': {'kyc_verified'},
    'kyc_verified': {'active'},
}

@app.route('/kyc/activate', methods=['POST'])
def activate():
    user = get_user(current_user.id)
    if 'active' not in TRANSITIONS.get(user.status, set()):
        return {'error': 'Étapes préalables non complétées'}, 409
    db.execute("UPDATE users SET status = 'active' WHERE id = %s", [user.id])
    log_workflow_transition(user.id, 'active')
    return '', 200
```

## Checklist de vérification post-patch
- [ ] Chaque étape sensible d'un workflow multi-étapes vérifie l'état d'avancement réel côté serveur, jamais un paramètre client.
- [ ] Une machine à états serveur définit explicitement les transitions valides et rejette les autres.
- [ ] L'endpoint d'étape finale (activation, confirmation, publication) vérifie que tous les prérequis sont satisfaits en base.
- [ ] Les transitions d'état sont journalisées pour détecter les tentatives de saut d'étape.
