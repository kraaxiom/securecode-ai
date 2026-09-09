# Remédiation — Mauvaise configuration de sécurité REST

## Principe
Restreindre explicitement chaque route à son verbe HTTP prévu, appliquer un contrôle d'autorisation homogène sur toutes les méthodes d'une ressource, désactiver les erreurs détaillées en production, et configurer CORS de façon restrictive.

## PHP (Laravel)
```php
// Avant — vulnérable
Route::any('/api/orders/{id}', [OrderController::class, 'handle']);

// Après — sécurisé
Route::get('/api/orders/{id}', [OrderController::class, 'show'])->middleware('auth:sanctum');
Route::delete('/api/orders/{id}', [OrderController::class, 'destroy'])->middleware(['auth:sanctum', 'can:delete,order']);
```

```php
// Avant — vulnérable (debug actif en prod)
// .env
APP_DEBUG=true

// Après — sécurisé
// .env (production)
APP_DEBUG=false
```

## Node.js (Express)
```js
// Avant — vulnérable
app.all('/api/orders/:id', ordersHandler); // pas de distinction de verbe/autorisation

// Après — sécurisé
app.get('/api/orders/:id', requireAuth, getOrder);
app.delete('/api/orders/:id', requireAuth, requireOwner, deleteOrder);

// gestion d'erreur générique en production
app.use((err, req, res, next) => {
  console.error(err); // log interne uniquement
  res.status(500).json({ error: 'Erreur interne du serveur' });
});
```

## CORS restrictif (Node.js)
```js
// Avant — vulnérable
app.use(cors({ origin: '*', credentials: true }));

// Après — sécurisé
const allowedOrigins = ['https://app.example.com'];
app.use(cors({
  origin: (origin, cb) => cb(null, allowedOrigins.includes(origin)),
  credentials: true,
}));
```

## Python (Django)
```python
# Avant — vulnérable
DEBUG = True  # en production

# Après — sécurisé
DEBUG = False
ALLOWED_HOSTS = ['app.example.com']
```

## Checklist de vérification post-patch
- [ ] Chaque route est restreinte à son verbe HTTP prévu, sans routage générique (`ANY`/`ALL`).
- [ ] Le contrôle d'autorisation est appliqué de façon homogène sur toutes les méthodes d'une même ressource.
- [ ] `debug`/`display_errors` est désactivé en production ; les erreurs renvoyées au client sont génériques.
- [ ] CORS utilise une liste blanche stricte d'origines, jamais de wildcard combiné à `credentials: true`.
- [ ] La documentation OpenAPI/Swagger n'est pas exposée publiquement sans authentification en production.
