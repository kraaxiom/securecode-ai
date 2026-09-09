# Remédiation — Excessive Data Exposure

## Principe
Ne jamais sérialiser un modèle interne/entité ORM tel quel dans une réponse API. Définir explicitement un schéma de sortie (DTO/serializer) en liste blanche des champs destinés au client, quel que soit le contexte d'appel.

## PHP (Laravel / Eloquent)
```php
// Avant — vulnérable
public function show($id)
{
    return response()->json(User::findOrFail($id));
}

// Après — sécurisé
public function show($id)
{
    $user = User::findOrFail($id);
    return response()->json([
        'id' => $user->id,
        'name' => $user->name,
        'email' => $user->email,
        // password, remember_token, api_token volontairement exclus
    ]);
}
```

## PHP (natif)
```php
// Avant — vulnérable
echo json_encode($user); // contient password_hash, reset_token...

// Après — sécurisé
echo json_encode([
    'id' => $user['id'],
    'name' => $user['name'],
    'email' => $user['email'],
]);
```

## Node.js (Express + Mongoose)
```js
// Avant — vulnérable
app.get('/api/users/:id', async (req, res) => {
  const user = await User.findById(req.params.id);
  res.json(user); // expose passwordHash, internalNotes...
});

// Après — sécurisé
app.get('/api/users/:id', async (req, res) => {
  const user = await User.findById(req.params.id).select('_id name email');
  res.json(user);
});
```

## Python (Django REST Framework)
```python
# Avant — vulnérable
class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = '__all__'  # inclut password, is_staff, tokens internes

# Après — sécurisé
class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ['id', 'username', 'email']  # liste blanche explicite
```

## Checklist de vérification post-patch
- [ ] Chaque endpoint utilise un DTO/serializer en liste blanche, jamais l'entité brute.
- [ ] Les champs sensibles (mots de passe hachés, tokens, notes internes) sont exclus par défaut, pas par exception.
- [ ] Un schéma de sortie distinct existe pour le contexte administrateur vs utilisateur standard.
- [ ] Une vérification de la réponse API brute (hors frontend) confirme l'absence de sur-exposition.
- [ ] Les messages d'erreur/debug ne fuient pas de structure interne (stack trace, requête SQL) en production.
