# Remédiation — Mass Assignment sur API

## Principe
Ne jamais lier directement le corps de requête à un modèle de persistance. Définir un DTO/schéma d'entrée avec une liste blanche explicite des champs autorisés, et protéger les champs sensibles (rôle, statut, solde) via les mécanismes natifs du framework.

## PHP (Laravel)
```php
// Avant — vulnérable
public function update(Request $request, $id)
{
    $user = User::findOrFail($id);
    $user->update($request->all());
    return $user;
}

// Après — sécurisé
public function update(Request $request, $id)
{
    $user = User::findOrFail($id);
    $validated = $request->validate([
        'name'  => 'sometimes|string|max:255',
        'email' => 'sometimes|email',
    ]);
    $user->update($validated); // $fillable sur le modèle exclut 'role', 'is_admin', 'balance'
    return $user;
}
```

## Node.js (Express)
```js
// Avant — vulnérable
app.put('/api/users/:id', async (req, res) => {
  const user = await User.findByIdAndUpdate(req.params.id, req.body, { new: true });
  res.json(user);
});

// Après — sécurisé
const { name, email } = req.body; // liste blanche explicite
app.put('/api/users/:id', async (req, res) => {
  const { name, email } = req.body;
  const user = await User.findByIdAndUpdate(
    req.params.id,
    { name, email }, // 'role', 'isAdmin' jamais acceptés depuis le client
    { new: true }
  );
  res.json(user);
});
```

## Python (Flask/Django/FastAPI)
```python
# Avant — vulnérable
@app.route("/api/users/<int:user_id>", methods=["PUT"])
def update_user(user_id):
    user = User.query.get_or_404(user_id)
    for key, value in request.json.items():
        setattr(user, key, value)
    db.session.commit()
    return jsonify(user.to_dict())

# Après — sécurisé (Pydantic / FastAPI)
class UserUpdateDTO(BaseModel):
    name: str | None = None
    email: EmailStr | None = None
    # 'role' et 'is_admin' volontairement absents du schéma d'entrée

@app.put("/api/users/{user_id}")
def update_user(user_id: int, payload: UserUpdateDTO):
    user = get_user_or_404(user_id)
    for field, value in payload.dict(exclude_unset=True).items():
        setattr(user, field, value)
    db.session.commit()
    return user
```

## Checklist de vérification post-patch
- [ ] Aucun endpoint de création/mise à jour ne passe le body brut (`request.all()`, `req.body`, `request.json`) directement à une opération de persistance.
- [ ] Un DTO ou schéma de validation d'entrée distinct du modèle existe pour chaque endpoint sensible.
- [ ] Les champs privilégiés (`role`, `is_admin`, `balance`, `verified`) sont protégés par `$fillable`/`$guarded`, une exclusion de schéma, ou une liste blanche explicite.
- [ ] Un test confirme qu'un payload contenant un champ non autorisé (ex: `"role":"admin"`) est ignoré sans erreur silencieuse côté serveur.
