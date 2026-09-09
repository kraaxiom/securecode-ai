# Remédiation — Insecure Direct Object Reference (IDOR)

## Principe
Toujours inclure une clause de vérification d'appartenance dans toute requête de récupération d'objet par identifiant, en complément (pas en remplacement) d'identifiants non séquentiels.

## PHP (Laravel)
```php
// Avant — vulnérable : l'ID est utilisé tel quel, sans vérifier le propriétaire
public function show($id)
{
    $invoice = Invoice::findOrFail($id);
    return view('invoices.show', compact('invoice'));
}

// Après — sécurisé : clause d'appartenance dans la requête elle-même
public function show($id)
{
    $invoice = Invoice::where('id', $id)
        ->where('user_id', auth()->id())
        ->firstOrFail();
    return view('invoices.show', compact('invoice'));
}
```

## JS / Node.js (Express)
```js
// Avant — vulnérable : document récupéré uniquement par ID transmis par le client
app.get('/api/documents/:id', requireAuth, async (req, res) => {
  const doc = await Document.findById(req.params.id);
  res.json(doc);
});

// Après — sécurisé : filtrage d'appartenance dans la requête de base de données
app.get('/api/documents/:id', requireAuth, async (req, res) => {
  const doc = await Document.findOne({ _id: req.params.id, ownerId: req.user.id });
  if (!doc) return res.status(404).json({ error: 'Introuvable' });
  res.json(doc);
});
```

## Python (Flask)
```python
# Avant — vulnérable : accès direct par ID sans vérification d'appartenance
@app.route("/documents/<int:doc_id>")
@login_required
def get_document(doc_id):
    doc = Document.query.get_or_404(doc_id)
    return jsonify(doc.to_dict())

# Après — sécurisé : filtrage d'appartenance appliqué à la requête
@app.route("/documents/<int:doc_id>")
@login_required
def get_document(doc_id):
    doc = Document.query.filter_by(id=doc_id, owner_id=current_user.id).first_or_404()
    return jsonify(doc.to_dict())
```

## Checklist de vérification post-patch
- [ ] Toute requête de récupération d'objet par identifiant inclut une clause de vérification d'appartenance (`WHERE user_id = ...`).
- [ ] La vérification d'autorisation ne se limite pas à "l'utilisateur est connecté" mais confirme qu'il possède/peut accéder à cet objet précis.
- [ ] Un test d'accès croisé entre deux comptes distincts confirme qu'un ID d'autrui renvoie 403/404.
- [ ] Le filtrage d'appartenance est appliqué au niveau de la requête de données, pas en post-traitement (ne pas charger puis vérifier après coup si évitable).
- [ ] Des identifiants non séquentiels (UUID) sont envisagés en complément du contrôle d'autorisation.
