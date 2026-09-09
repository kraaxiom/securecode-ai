# Remédiation — Credential Stuffing

## Principe
Compléter le rate limiting classique par une détection comportementale (vélocité, diversité de comptes testés depuis une même source) et imposer/proposer la MFA, seule protection réellement efficace contre la réutilisation de mots de passe volés.

## PHP (Laravel)
```php
// Avant — vulnérable : pas de détection de trafic distribué
public function login(Request $request)
{
    if (Auth::attempt($request->only('email', 'password'))) {
        return redirect()->intended();
    }
    return back()->withErrors(['email' => 'Identifiants invalides']);
}

// Après — sécurisé : détection de vélocité globale + MFA obligatoire
public function login(Request $request)
{
    $ipKey = 'login-attempts:' . $request->ip();
    if (RateLimiter::tooManyAttempts($ipKey, 30)) { // volume anormal tous comptes confondus
        SecurityAlert::dispatch('credential_stuffing_suspected', $request->ip());
        return back()->withErrors(['email' => 'Trop de tentatives depuis cette origine.']);
    }
    RateLimiter::hit($ipKey, 600);

    if (Auth::attempt($request->only('email', 'password'))) {
        if ($request->user()->mfa_enabled) {
            return redirect()->route('mfa.challenge');
        }
        return redirect()->intended();
    }
    return back()->withErrors(['email' => 'Identifiants invalides']);
}
```

## JS / Node.js (Express)
```js
// Avant — vulnérable : aucune corrélation entre tentatives
app.post('/login', async (req, res) => {
  const user = await authenticate(req.body.email, req.body.password);
  if (!user) return res.status(401).json({ error: 'Identifiants invalides' });
  res.json({ token: issueToken(user) });
});

// Après — sécurisé : détection de vélocité + vérification contre fuites connues + MFA
app.post('/login', async (req, res) => {
  const velocity = await getFailedAttemptVelocity(req.ip); // tous comptes confondus
  if (velocity > THRESHOLD) {
    await flagSuspiciousSource(req.ip);
    return res.status(429).json({ error: 'Trafic anormal détecté, réessayez plus tard.' });
  }

  const user = await authenticate(req.body.email, req.body.password);
  if (!user) {
    await recordFailedAttempt(req.body.email, req.ip);
    return res.status(401).json({ error: 'Identifiants invalides' });
  }
  if (user.mfaEnabled) {
    return res.json({ mfaRequired: true, challengeId: await createMfaChallenge(user) });
  }
  res.json({ token: issueToken(user) });
});
```

## Python (Django)
```python
# Avant — vulnérable : pas de détection comportementale
def login_view(request):
    user = authenticate(request, username=request.POST["email"], password=request.POST["password"])
    if user is None:
        return JsonResponse({"error": "Identifiants invalides"}, status=401)
    login(request, user)
    return JsonResponse({"ok": True})

# Après — sécurisé : vélocité par IP + MFA
def login_view(request):
    ip = get_client_ip(request)
    if get_failed_attempt_velocity(ip) > SUSPICIOUS_THRESHOLD:
        flag_suspicious_source(ip)
        return JsonResponse({"error": "Trafic anormal détecté."}, status=429)

    user = authenticate(request, username=request.POST["email"], password=request.POST["password"])
    if user is None:
        record_failed_attempt(request.POST["email"], ip)
        return JsonResponse({"error": "Identifiants invalides"}, status=401)

    if user.profile.mfa_enabled:
        return JsonResponse({"mfa_required": True, "challenge_id": create_mfa_challenge(user)})
    login(request, user)
    return JsonResponse({"ok": True})
```

## Checklist de vérification post-patch
- [ ] Une détection de vélocité globale (par IP/source, tous comptes confondus) complète le rate limiting par compte.
- [ ] La MFA est proposée ou imposée sur les comptes sensibles.
- [ ] Les mots de passe sont vérifiés contre une liste de fuites connues à l'inscription/au changement.
- [ ] Un test confirme qu'un volume élevé de tentatives sur des comptes différents depuis une même source déclenche une alerte/blocage.
- [ ] Les alertes de sécurité sont routées vers une équipe/outil de supervision (pas seulement journalisées).
