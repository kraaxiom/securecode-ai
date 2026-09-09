# Remédiation — Mass Assignment

## Principe
Utiliser une liste blanche explicite des champs autorisés en entrée pour chaque opération de création/mise à jour, et séparer le modèle de données interne des objets d'entrée (DTO) exposés à l'utilisateur.

## PHP (Laravel)
```php
// Avant — vulnérable : le corps de la requête entier alimente le modèle
public function update(Request $request, $id)
{
    $user = User::findOrFail($id);
    $user->update($request->all()); // 'role', 'is_admin' etc. peuvent être injectés
    return response()->json($user);
}

// Après — sécurisé : liste blanche explicite des champs modifiables
public function update(Request $request, $id)
{
    $user = User::findOrFail($id);
    $validated = $request->validate([
        'name' => 'sometimes|string|max:255',
        'email' => 'sometimes|email',
    ]); // 'role' et 'is_admin' ne peuvent jamais transiter par cette validation
    $user->update($validated);
    return response()->json($user);
}
// Complément défense en profondeur : $fillable = ['name', 'email']; sur le modèle User
```

## JS / Node.js (Express + Mongoose)
```js
// Avant — vulnérable : liaison directe du body au modèle
app.put('/api/users/:id', requireAuth, async (req, res) => {
  const user = await User.findByIdAndUpdate(req.params.id, req.body, { new: true });
  res.json(user);
});

// Après — sécurisé : DTO explicite, liste blanche de champs
app.put('/api/users/:id', requireAuth, async (req, res) => {
  const { name, email } = req.body; // extraction explicite, 'role'/'isAdmin' ignorés
  const user = await User.findByIdAndUpdate(req.params.id, { name, email }, { new: true });
  res.json(user);
});
```

## Python (Django REST Framework)
```python
# Avant — vulnérable : serializer exposant tous les champs du modèle, y compris 'role'
class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = "__all__"

# Après — sécurisé : serializer d'entrée restreint, séparé du modèle interne complet
class UserUpdateSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ["name", "email"]  # 'role', 'is_staff' explicitement exclus
        read_only_fields = []
```

## Checklist de vérification post-patch
- [ ] Chaque opération de création/mise à jour utilise une liste blanche explicite de champs autorisés, jamais une liaison automatique complète du corps de requête.
- [ ] Les attributs sensibles (rôle, statut, solde, drapeaux internes) sont exclus des DTO/serializers exposés à l'utilisateur.
- [ ] Le modèle de données interne est distinct de l'objet d'entrée accepté depuis le client.
- [ ] Un test confirme qu'un champ non autorisé (`role`, `isAdmin`) envoyé dans la requête est ignoré, pas appliqué.
- [ ] Les attributs sensibles sont marqués non assignables en masse au niveau du framework/ORM (`$guarded`/`$fillable`, etc.).
