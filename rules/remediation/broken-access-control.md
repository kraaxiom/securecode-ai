# Remédiation — Broken Access Control (catégorie générale)

## Principe
Centraliser la logique d'autorisation dans une couche unique et systématiquement appliquée (middleware, policy, guard), appliquer le principe "deny by default", et tester explicitement les cas de refus d'accès.

## PHP (Laravel)
```php
// Avant — vulnérable : aucune vérification d'autorisation sur une ressource sensible
Route::get('/invoices/{id}/download', [InvoiceController::class, 'download'])->middleware('auth');

// Après — sécurisé : policy centralisée appliquée systématiquement, deny by default
Route::get('/invoices/{id}/download', [InvoiceController::class, 'download'])
    ->middleware(['auth', 'can:view,invoice']);

// InvoicePolicy.php — refus par défaut sauf autorisation explicite
public function view(User $user, Invoice $invoice): bool
{
    return $invoice->user_id === $user->id || $user->role === 'admin';
}
```

## JS / Node.js (Express)
```js
// Avant — vulnérable : contrôle d'accès uniquement côté interface (bouton masqué)
router.get('/invoices/:id/download', requireAuth, async (req, res) => {
  const invoice = await Invoice.findById(req.params.id);
  res.download(invoice.filePath);
});

// Après — sécurisé : middleware d'autorisation centralisé, deny by default
function authorize(checkFn) {
  return async (req, res, next) => {
    const resource = await req.loadResource?.() ?? await Invoice.findById(req.params.id);
    if (!resource || !checkFn(req.user, resource)) {
      return res.status(403).json({ error: 'Accès refusé' });
    }
    req.resource = resource;
    next();
  };
}

router.get(
  '/invoices/:id/download',
  requireAuth,
  authorize((user, invoice) => invoice.ownerId === user.id || user.role === 'admin'),
  (req, res) => res.download(req.resource.filePath)
);
```

## Python (Django)
```python
# Avant — vulnérable : logique d'autorisation dupliquée et incohérente entre vues
def download_invoice(request, invoice_id):
    invoice = Invoice.objects.get(id=invoice_id)  # aucune vérification d'autorisation
    return FileResponse(invoice.file)

# Après — sécurisé : vérification centralisée, deny by default
def download_invoice(request, invoice_id):
    invoice = get_object_or_404(Invoice, id=invoice_id)
    if not (invoice.owner_id == request.user.id or request.user.role == "admin"):
        return HttpResponseForbidden("Accès refusé")
    return FileResponse(invoice.file)
```

## Checklist de vérification post-patch
- [ ] Chaque ressource/action sensible passe par une couche d'autorisation centralisée, pas une vérification ad hoc dispersée.
- [ ] Le principe "deny by default" est appliqué : tout accès est refusé sauf autorisation explicite.
- [ ] Des tests automatisés couvrent explicitement les cas de refus d'accès (pas seulement les cas de succès).
- [ ] Aucun contrôle d'accès ne repose uniquement sur le masquage d'éléments côté client.
- [ ] Les règles d'autorisation sont cohérentes entre toutes les couches (contrôleur, service, base de données).
