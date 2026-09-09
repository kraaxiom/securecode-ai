# Remédiation — Contournement de l'authentification multi-facteurs (MFA Bypass)

## Principe
Ne jamais délivrer un token/session pleinement privilégié avant validation complète et côté serveur du second facteur. Utiliser un état de session intermédiaire à privilèges limités entre les deux étapes.

## PHP (Laravel)
```php
// Avant — vulnérable : session pleinement authentifiée avant vérification MFA
public function login(Request $request)
{
    if (Auth::attempt($request->only('email', 'password'))) {
        return redirect()->route('dashboard'); // accès complet avant MFA !
    }
    return back()->withErrors(['email' => 'Identifiants invalides']);
}

// Après — sécurisé : état intermédiaire tant que le second facteur n'est pas validé
public function login(Request $request)
{
    if (Auth::attempt($request->only('email', 'password'))) {
        $user = Auth::user();
        if ($user->mfa_enabled) {
            Auth::logout(); // pas de session pleine tant que MFA non validé
            session(['mfa_pending_user_id' => $user->id]);
            return redirect()->route('mfa.challenge');
        }
        return redirect()->route('dashboard');
    }
    return back()->withErrors(['email' => 'Identifiants invalides']);
}

public function verifyMfa(Request $request)
{
    $userId = session('mfa_pending_user_id');
    abort_unless($userId, 403);
    $user = User::findOrFail($userId);

    if (! Mfa::verifyCode($user, $request->input('code'))) {
        return back()->withErrors(['code' => 'Code invalide']);
    }
    session()->forget('mfa_pending_user_id');
    Auth::login($user); // session pleinement privilégiée uniquement ici
    return redirect()->route('dashboard');
}
```

## JS / Node.js (Express)
```js
// Avant — vulnérable : token complet émis avant validation MFA
app.post('/login', async (req, res) => {
  const user = await authenticate(req.body.email, req.body.password);
  if (!user) return res.status(401).json({ error: 'Identifiants invalides' });
  res.json({ token: issueFullToken(user) }); // privilège complet avant MFA !
});

// Après — sécurisé : token intermédiaire limité, token complet après MFA
app.post('/login', async (req, res) => {
  const user = await authenticate(req.body.email, req.body.password);
  if (!user) return res.status(401).json({ error: 'Identifiants invalides' });

  if (user.mfaEnabled) {
    const partialToken = issuePartialToken(user, { scope: 'mfa_pending' });
    return res.json({ mfaRequired: true, partialToken });
  }
  res.json({ token: issueFullToken(user) });
});

app.post('/mfa/verify', requirePartialToken, async (req, res) => {
  const valid = await verifyMfaCode(req.user, req.body.code);
  if (!valid) return res.status(401).json({ error: 'Code invalide' });
  res.json({ token: issueFullToken(req.user) }); // uniquement après validation serveur
});
```

## Python (Django)
```python
# Avant — vulnérable : login() complet appelé avant vérification MFA
def login_view(request):
    user = authenticate(request, username=request.POST["email"], password=request.POST["password"])
    if user is None:
        return JsonResponse({"error": "Identifiants invalides"}, status=401)
    login(request, user)  # session complète avant MFA !
    return JsonResponse({"ok": True})

# Après — sécurisé : état de session intermédiaire jusqu'à validation MFA
def login_view(request):
    user = authenticate(request, username=request.POST["email"], password=request.POST["password"])
    if user is None:
        return JsonResponse({"error": "Identifiants invalides"}, status=401)

    if user.profile.mfa_enabled:
        request.session["mfa_pending_user_id"] = user.id  # pas de login() complet
        return JsonResponse({"mfa_required": True})

    login(request, user)
    return JsonResponse({"ok": True})

def verify_mfa_view(request):
    user_id = request.session.get("mfa_pending_user_id")
    if not user_id:
        return JsonResponse({"error": "Aucune session MFA en attente"}, status=403)
    user = User.objects.get(id=user_id)
    if not verify_mfa_code(user, request.POST["code"]):
        return JsonResponse({"error": "Code invalide"}, status=401)
    del request.session["mfa_pending_user_id"]
    login(request, user)  # session complète uniquement ici
    return JsonResponse({"ok": True})
```

## Checklist de vérification post-patch
- [ ] Aucun token/session pleinement privilégié n'est émis avant validation serveur complète du second facteur.
- [ ] L'état "MFA en attente" est stocké côté serveur (session), jamais dans un paramètre client modifiable.
- [ ] Tous les points d'entrée (web, mobile, API) appliquent la même politique MFA.
- [ ] Un test confirme qu'un token intermédiaire ne permet pas d'accéder aux endpoints protégés sans validation MFA.
- [ ] Le code MFA est à usage unique et expire après un délai court.
