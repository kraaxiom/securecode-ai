# Remédiation — Privilege Escalation

## Principe
Protéger explicitement tout champ de rôle/permission contre la modification directe, vérifier que l'appelant d'une fonction d'attribution de rôle est lui-même autorisé à accorder ce niveau de privilège, et invalider les sessions après tout changement de rôle.

## PHP (Laravel)
```php
// Avant — vulnérable : attribution de rôle sans vérifier le niveau de l'appelant
public function assignRole(Request $request, $userId)
{
    $user = User::findOrFail($userId);
    $user->role = $request->input('role'); // n'importe quel rôle, y compris 'admin'
    $user->save();
    return response()->json($user);
}

// Après — sécurisé : vérification que l'appelant peut accorder ce rôle précis + invalidation de session
public function assignRole(Request $request, $userId)
{
    $requestedRole = $request->input('role');
    if (! auth()->user()->canGrantRole($requestedRole)) {
        abort(403, 'Vous ne pouvez pas attribuer ce niveau de privilège.');
    }
    $user = User::findOrFail($userId);
    $user->role = $requestedRole;
    $user->save();
    $user->tokens()->delete(); // invalide les sessions/tokens existants
    return response()->json($user);
}
```

## JS / Node.js (Express)
```js
// Avant — vulnérable : le rôle est mis à jour comme n'importe quel autre champ de profil
app.put('/api/users/:id/role', requireAuth, async (req, res) => {
  const user = await User.findByIdAndUpdate(req.params.id, { role: req.body.role }, { new: true });
  res.json(user);
});

// Après — sécurisé : vérification du droit d'attribution + révocation des sessions actives
app.put('/api/users/:id/role', requireAuth, async (req, res) => {
  if (!canGrantRole(req.user, req.body.role)) {
    return res.status(403).json({ error: 'Attribution de rôle non autorisée' });
  }
  const user = await User.findByIdAndUpdate(req.params.id, { role: req.body.role }, { new: true });
  await revokeActiveSessions(user.id); // force une réauthentification avec le nouveau rôle
  res.json(user);
});
```

## Python (Django REST Framework)
```python
# Avant — vulnérable : rôle modifiable via le serializer standard sans contrôle du niveau accordé
class UserRoleUpdateView(APIView):
    def put(self, request, user_id):
        user = User.objects.get(id=user_id)
        user.role = request.data["role"]
        user.save()
        return Response({"ok": True})

# Après — sécurisé : vérification explicite du droit d'attribution + invalidation de session
class UserRoleUpdateView(APIView):
    def put(self, request, user_id):
        requested_role = request.data["role"]
        if not request.user.can_grant_role(requested_role):
            return Response({"error": "Attribution non autorisée"}, status=403)
        user = User.objects.get(id=user_id)
        user.role = requested_role
        user.save()
        Session.objects.filter(user=user).delete()  # invalide les sessions actives
        return Response({"ok": True})
```

## Checklist de vérification post-patch
- [ ] Tout champ de rôle/permission est protégé contre la modification directe par l'utilisateur (voir aussi remédiation mass assignment).
- [ ] Toute fonction d'attribution de rôle vérifie que l'appelant est lui-même autorisé à accorder ce niveau de privilège précis.
- [ ] Les sessions/tokens sont invalidés ou rafraîchis après tout changement de rôle.
- [ ] Un test confirme qu'un utilisateur standard ne peut pas s'auto-attribuer ni attribuer à un tiers un rôle supérieur au sien.
- [ ] Les claims de rôle dans les tokens sont revérifiés après une opération sensible de changement de permissions.
