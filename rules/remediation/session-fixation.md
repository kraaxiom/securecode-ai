# Remédiation — Session Fixation

## Principe
Régénérer systématiquement l'identifiant de session immédiatement après une authentification réussie ou tout changement de niveau de privilège, et invalider l'ancienne session côté serveur. Ne jamais accepter un identifiant de session provenant d'un paramètre d'URL ou de formulaire.

## PHP
```php
// Avant — vulnérable
session_start();
if (verifierIdentifiants($_POST['login'], $_POST['password'])) {
    $_SESSION['user_id'] = $userId;
    $_SESSION['authenticated'] = true;
}

// Après — sécurisé
session_start();
if (verifierIdentifiants($_POST['login'], $_POST['password'])) {
    session_regenerate_id(true); // détruit l'ancienne session, en crée une nouvelle
    $_SESSION['user_id'] = $userId;
    $_SESSION['authenticated'] = true;
}
```

## Node.js (Express)
```js
// Avant — vulnérable
app.post('/login', (req, res) => {
  if (checkCredentials(req.body)) {
    req.session.userId = user.id;
    res.redirect('/dashboard');
  }
});

// Après — sécurisé
app.post('/login', (req, res) => {
  if (checkCredentials(req.body)) {
    req.session.regenerate((err) => {
      if (err) return res.status(500).end();
      req.session.userId = user.id;
      res.redirect('/dashboard');
    });
  }
});
```

## Python (Flask/Django)
```python
# Avant — vulnérable (Flask)
@app.route('/login', methods=['POST'])
def login():
    if check_credentials(request.form):
        session['user_id'] = user.id
        return redirect('/dashboard')

# Après — sécurisé (Flask)
@app.route('/login', methods=['POST'])
def login():
    if check_credentials(request.form):
        session.clear()  # invalide toute donnée de session préexistante
        session['user_id'] = user.id  # Flask régénère l'ID de session à la prochaine réponse
        return redirect('/dashboard')

# Django : authentifier puis appeler login() qui régénère la session automatiquement
from django.contrib.auth import login as django_login

def login_view(request):
    user = authenticate(request, username=..., password=...)
    if user is not None:
        django_login(request, user)  # cycle_key() régénère l'ID de session
        return redirect('/dashboard')
```

## Checklist de vérification post-patch
- [ ] L'identifiant de session est régénéré immédiatement après tout login réussi.
- [ ] L'ancienne session est invalidée côté serveur, pas seulement remplacée côté client.
- [ ] Aucun identifiant de session n'est accepté depuis un paramètre GET/POST externe.
- [ ] Un test confirme que l'ID de session avant et après authentification diffère.
