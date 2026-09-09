# Remédiation — Brute Force (absence de limitation des tentatives d'authentification)

## Principe
Limiter le nombre de tentatives d'authentification par compte ET par IP, avec verrouillage progressif, et journaliser les échecs. Ne jamais se reposer uniquement sur la complexité du mot de passe.

## PHP (Laravel)
```php
// Avant — vulnérable : aucune limitation
public function login(Request $request)
{
    if (Auth::attempt($request->only('email', 'password'))) {
        return redirect()->intended();
    }
    return back()->withErrors(['email' => 'Identifiants invalides']);
}

// Après — sécurisé : throttle par IP + par compte
public function login(Request $request)
{
    $key = 'login:' . $request->ip() . ':' . strtolower($request->input('email'));

    if (RateLimiter::tooManyAttempts($key, 5)) {
        $seconds = RateLimiter::availableIn($key);
        return back()->withErrors(['email' => "Trop de tentatives. Réessayez dans {$seconds}s."]);
    }

    if (Auth::attempt($request->only('email', 'password'))) {
        RateLimiter::clear($key);
        return redirect()->intended();
    }

    RateLimiter::hit($key, 900); // fenêtre de 15 min
    return back()->withErrors(['email' => 'Identifiants invalides']);
}
```

## JS / Node.js (Express)
```js
// Avant — vulnérable : aucune limitation
app.post('/login', async (req, res) => {
  const user = await authenticate(req.body.email, req.body.password);
  if (!user) return res.status(401).json({ error: 'Identifiants invalides' });
  res.json({ token: issueToken(user) });
});

// Après — sécurisé : rate limiting + verrouillage progressif par compte
const rateLimit = require('express-rate-limit');

const loginLimiter = rateLimit({
  windowMs: 15 * 60 * 1000,
  max: 5,
  keyGenerator: (req) => `${req.ip}:${req.body.email}`,
  message: { error: 'Trop de tentatives. Réessayez plus tard.' },
});

app.post('/login', loginLimiter, async (req, res) => {
  const user = await authenticate(req.body.email, req.body.password);
  if (!user) {
    await recordFailedAttempt(req.body.email, req.ip);
    return res.status(401).json({ error: 'Identifiants invalides' });
  }
  await clearFailedAttempts(req.body.email);
  res.json({ token: issueToken(user) });
});
```

## Python (Flask)
```python
# Avant — vulnérable : aucune limitation
@app.route("/login", methods=["POST"])
def login():
    user = authenticate(request.form["email"], request.form["password"])
    if not user:
        return jsonify(error="Identifiants invalides"), 401
    return jsonify(token=issue_token(user))

# Après — sécurisé : flask-limiter par compte + par IP
from flask_limiter import Limiter
from flask_limiter.util import get_remote_address

limiter = Limiter(app, key_func=get_remote_address)

@app.route("/login", methods=["POST"])
@limiter.limit("5 per 15 minutes", key_func=lambda: f"{get_remote_address()}:{request.form.get('email','')}")
def login():
    user = authenticate(request.form["email"], request.form["password"])
    if not user:
        record_failed_attempt(request.form["email"], get_remote_address())
        return jsonify(error="Identifiants invalides"), 401
    clear_failed_attempts(request.form["email"])
    return jsonify(token=issue_token(user))
```

## Checklist de vérification post-patch
- [ ] La limitation s'applique à la fois par compte et par IP/source (pas l'un sans l'autre).
- [ ] Le verrouillage est progressif (délai croissant) et non permanent sans procédure de déblocage.
- [ ] Les échecs d'authentification sont journalisés avec horodatage, compte ciblé et IP source.
- [ ] Un test confirme que la 6e tentative consécutive échouée est bloquée dans la fenêtre de temps définie.
- [ ] Le message d'erreur ne permet pas de distinguer un compte existant d'un compte inexistant.
