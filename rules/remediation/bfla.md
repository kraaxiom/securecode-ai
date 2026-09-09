# Remédiation — Broken Function Level Authorization (BFLA)

## Principe
Vérifier explicitement le rôle/permission requis sur chaque endpoint exposant une fonctionnalité sensible, y compris ceux non visibles dans l'interface, sans jamais s'appuyer sur le masquage côté client.

## PHP (Laravel)
```php
// Avant — vulnérable : endpoint admin protégé uniquement par l'authentification
Route::delete('/users/{id}', [UserController::class, 'destroy'])->middleware('auth');

// Après — sécurisé : middleware/policy de rôle explicite
Route::delete('/users/{id}', [UserController::class, 'destroy'])
    ->middleware(['auth', 'can:delete,App\Models\User']);

// UserPolicy.php
public function delete(User $actor, User $target): bool
{
    return $actor->role === 'admin';
}
```

## JS / Node.js (Express)
```js
// Avant — vulnérable : seule l'authentification est vérifiée
router.delete('/users/:id', requireAuth, async (req, res) => {
  await User.deleteOne({ _id: req.params.id });
  res.sendStatus(204);
});

// Après — sécurisé : vérification de rôle explicite et centralisée
function requireRole(role) {
  return (req, res, next) => {
    if (req.user.role !== role) return res.status(403).json({ error: 'Accès refusé' });
    next();
  };
}

router.delete('/users/:id', requireAuth, requireRole('admin'), async (req, res) => {
  await User.deleteOne({ _id: req.params.id });
  res.sendStatus(204);
});
```

## Python (Django REST Framework)
```python
# Avant — vulnérable : IsAuthenticated seul, pas de vérification de rôle
class UserViewSet(viewsets.ModelViewSet):
    permission_classes = [IsAuthenticated]

# Après — sécurisé : permission de rôle explicite sur l'action sensible
class IsAdmin(BasePermission):
    def has_permission(self, request, view):
        return request.user.is_authenticated and request.user.role == "admin"

class UserViewSet(viewsets.ModelViewSet):
    permission_classes = [IsAuthenticated]

    def get_permissions(self):
        if self.action == "destroy":
            return [IsAdmin()]
        return super().get_permissions()
```

## Checklist de vérification post-patch
- [ ] Chaque endpoint sensible vérifie explicitement le rôle/permission requis, pas seulement l'authentification.
- [ ] Le contrôle de rôle est centralisé (middleware/policy/permission class), pas dupliqué par endpoint.
- [ ] Les endpoints non visibles dans l'UI (mobile, API interne, legacy) appliquent le même contrôle.
- [ ] Un test confirme qu'un utilisateur standard reçoit un 403 sur une fonction d'administration.
- [ ] Une cartographie des fonctionnalités par niveau de privilège requis existe et est testée automatiquement.
