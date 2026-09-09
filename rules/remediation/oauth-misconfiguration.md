# Remédiation — Mauvaise configuration OAuth / OpenID Connect

## Principe
Valider `redirect_uri` par correspondance exacte avec une liste blanche, générer et vérifier un `state` (et `nonce` en OIDC) unique par requête, et valider intégralement le jeton d'identité reçu (signature, émetteur, audience, expiration).

## PHP (league/oauth2-client)
```php
// Avant — vulnérable : pas de state, redirect_uri validé par préfixe côté fournisseur
$authUrl = $provider->getAuthorizationUrl(); // pas de state généré/stocké

// Après — sécurisé : state unique généré et vérifié, redirect_uri en liste blanche exacte
$state = bin2hex(random_bytes(16));
session(['oauth2_state' => $state]);
$authUrl = $provider->getAuthorizationUrl(['state' => $state]);

// Sur le callback :
public function callback(Request $request)
{
    if (! $request->has('state') || $request->input('state') !== session('oauth2_state')) {
        abort(403, 'State invalide — tentative de CSRF sur le callback OAuth.');
    }
    session()->forget('oauth2_state');
    $token = $provider->getAccessToken('authorization_code', ['code' => $request->input('code')]);
    // ...
}
```

## JS / Node.js (openid-client)
```js
// Avant — vulnérable : pas de vérification de state/nonce, ID token non validé
app.get('/callback', async (req, res) => {
  const tokenSet = await client.callback(redirectUri, req.query);
  res.json(tokenSet.claims()); // claims utilisés sans validation stricte iss/aud
});

// Après — sécurisé : state et nonce générés, ID token validé intégralement par la lib
app.get('/login', (req, res) => {
  const state = generators.state();
  const nonce = generators.nonce();
  req.session.state = state;
  req.session.nonce = nonce;
  res.redirect(client.authorizationUrl({ scope: 'openid profile', state, nonce }));
});

app.get('/callback', async (req, res) => {
  const params = client.callbackParams(req);
  const tokenSet = await client.callback(redirectUri, params, {
    state: req.session.state,
    nonce: req.session.nonce, // la lib valide iss/aud/exp/signature/nonce
  });
  res.json(tokenSet.claims());
});
```

## Python (Authlib / Django)
```python
# Avant — vulnérable : redirect_uri non contrôlé strictement, state ignoré au retour
def callback_view(request):
    token = oauth.provider.authorize_access_token(request)  # state non vérifié explicitement
    return JsonResponse(token)

# Après — sécurisé : state généré/stocké, redirect_uri en liste blanche exacte, ID token validé
def login_view(request):
    state = secrets.token_urlsafe(24)
    request.session["oauth_state"] = state
    redirect_uri = "https://app.example.com/callback"  # correspondance exacte, pas de wildcard
    return oauth.provider.authorize_redirect(request, redirect_uri, state=state)

def callback_view(request):
    if request.GET.get("state") != request.session.pop("oauth_state", None):
        return JsonResponse({"error": "State invalide"}, status=403)
    token = oauth.provider.authorize_access_token(request)  # signature/iss/aud/exp validés par la lib
    return JsonResponse(token)
```

## Checklist de vérification post-patch
- [ ] Les `redirect_uri` enregistrés sont validés par correspondance exacte, sans wildcard ni préfixe large.
- [ ] Un `state` unique et imprévisible est généré par requête d'autorisation et strictement vérifié au retour.
- [ ] Un `nonce` est utilisé en OpenID Connect et vérifié dans l'ID token.
- [ ] L'ID token est intégralement validé (signature, `iss`, `aud`, expiration) avant utilisation.
- [ ] Un test confirme qu'un callback avec `state` absent ou incorrect est rejeté.
