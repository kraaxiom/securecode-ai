# Remédiation — Absence de limitation de débit (rate limiting)

## Principe
Limiter le débit par IP, par utilisateur et/ou par clé API sur tous les endpoints sensibles ou coûteux, avec verrouillage progressif après échecs d'authentification répétés.

## PHP (Laravel)
```php
// Avant — vulnérable
Route::post('/login', [AuthController::class, 'login']);

// Après — sécurisé
Route::post('/login', [AuthController::class, 'login'])
    ->middleware('throttle:5,1'); // 5 tentatives par minute
```

## PHP (natif, compteur Redis)
```php
// Avant — vulnérable
function login($email, $password) { /* pas de compteur */ }

// Après — sécurisé
function login($email, $password) {
    $key = "login_attempts:$email";
    $attempts = (int) $redis->get($key);
    if ($attempts >= 5) {
        http_response_code(429);
        exit(json_encode(['error' => 'Trop de tentatives, réessayez plus tard.']));
    }
    $redis->incr($key);
    $redis->expire($key, 60);
    // ... suite de la logique de connexion
}
```

## Node.js (Express + express-rate-limit)
```js
// Avant — vulnérable
app.post('/api/login', loginHandler);

// Après — sécurisé
const rateLimit = require('express-rate-limit');
const loginLimiter = rateLimit({
  windowMs: 15 * 60 * 1000,
  max: 5,
  standardHeaders: true,
  message: { error: 'Trop de tentatives, réessayez plus tard.' },
});
app.post('/api/login', loginLimiter, loginHandler);
```

## Python (Flask + Flask-Limiter)
```python
# Avant — vulnérable
@app.route('/login', methods=['POST'])
def login():
    ...

# Après — sécurisé
from flask_limiter import Limiter
limiter = Limiter(app, key_func=get_remote_address)

@app.route('/login', methods=['POST'])
@limiter.limit("5 per minute")
def login():
    ...
```

## Checklist de vérification post-patch
- [ ] Tous les endpoints d'authentification, d'inscription et de réinitialisation de mot de passe sont limités en débit.
- [ ] Les API publiques/partenaires disposent d'un quota par clé API ou par IP.
- [ ] Un verrouillage progressif ou un délai croissant s'applique après échecs répétés.
- [ ] Le serveur renvoie `429 Too Many Requests` avec en-têtes indiquant les quotas restants (`Retry-After`).
- [ ] Un test confirme que le seuil de limitation déclenche correctement le blocage.
