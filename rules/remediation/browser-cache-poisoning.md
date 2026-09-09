# Remédiation — Browser Cache Poisoning

## Principe
Définir explicitement `Cache-Control: no-store` (ou `private, no-store`) sur toute réponse contenant des données sensibles ou spécifiques à l'utilisateur, en particulier les pages d'authentification, de déconnexion et les vues de profil/panier/solde. Ne jamais s'appuyer sur les valeurs par défaut du framework.

## PHP
```php
// Avant — vulnérable : aucune directive de cache sur une page personnalisée
header('Content-Type: text/html');
echo renderProfile($user);

// Après — sécurisé
header('Cache-Control: no-store, private');
header('Pragma: no-cache');
header('Content-Type: text/html');
echo renderProfile($user);
```

## PHP (Laravel)
```php
// Avant — vulnérable
return response()->view('profile', compact('user'));

// Après — sécurisé
return response()->view('profile', compact('user'))
    ->header('Cache-Control', 'no-store, private')
    ->header('Pragma', 'no-cache');
```

## Node.js (Express)
```js
// Avant — vulnérable
app.get('/account', (req, res) => {
  res.render('account', { user: req.user });
});

// Après — sécurisé
app.get('/account', (req, res) => {
  res.set('Cache-Control', 'no-store, private');
  res.set('Pragma', 'no-cache');
  res.render('account', { user: req.user });
});
```

## Python (Flask)
```python
# Avant — vulnérable
@app.route('/account')
def account():
    return render_template('account.html', user=current_user)

# Après — sécurisé
@app.route('/account')
def account():
    resp = make_response(render_template('account.html', user=current_user))
    resp.headers['Cache-Control'] = 'no-store, private'
    resp.headers['Pragma'] = 'no-cache'
    return resp
```

## Java (Spring)
```java
// Avant — vulnérable
@GetMapping("/account")
public String account(Model model) { return "account"; }

// Après — sécurisé
@GetMapping("/account")
public String account(Model model, HttpServletResponse response) {
    response.setHeader("Cache-Control", "no-store, private");
    response.setHeader("Pragma", "no-cache");
    return "account";
}
```

## Checklist de vérification post-patch
- [ ] Toute route affichant des données personnelles/de session renvoie `Cache-Control: no-store` (ou `private, no-store`).
- [ ] Les pages de login/logout/changement de mot de passe portent un en-tête empêchant leur mise en cache navigateur.
- [ ] Aucun en-tête `Cache-Control: public, max-age=...` n'est appliqué globalement à des routes dynamiques personnalisées.
- [ ] Test manuel : après déconnexion, le bouton "précédent" du navigateur ne réaffiche pas la page authentifiée depuis le cache.
