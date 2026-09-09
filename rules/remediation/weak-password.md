# Remédiation — Politique de mot de passe faible

## Principe
Exiger une longueur minimale d'au moins 12 caractères sans règles de composition artificielles, vérifier le mot de passe contre une liste de mots de passe compromis connus, et ne pas imposer de rotation périodique sans motif (conforme NIST 800-63B).

## PHP (Laravel)
```php
// Avant — vulnérable
$request->validate([
    'password' => 'required|min:6',
]);

// Après — sécurisé
use Illuminate\Validation\Rules\Password;

$request->validate([
    'password' => ['required', Password::min(12)->uncompromised()],
]);
```

## Node.js (Express + zxcvbn/HIBP)
```js
// Avant — vulnérable
if (password.length < 6) return res.status(400).send('Mot de passe trop court');

// Après — sécurisé
const { isPasswordPwned } = require('hibp');

if (password.length < 12) {
  return res.status(400).send('Le mot de passe doit contenir au moins 12 caractères');
}
if (await isPasswordPwned(password)) {
  return res.status(400).send('Ce mot de passe a été trouvé dans une fuite de données connue, choisissez-en un autre');
}
```

## Python (Django)
```python
# Avant — vulnérable
AUTH_PASSWORD_VALIDATORS = [
    {'NAME': 'django.contrib.auth.password_validation.MinimumLengthValidator', 'OPTIONS': {'min_length': 6}},
]

# Après — sécurisé
AUTH_PASSWORD_VALIDATORS = [
    {'NAME': 'django.contrib.auth.password_validation.MinimumLengthValidator', 'OPTIONS': {'min_length': 12}},
    {'NAME': 'django.contrib.auth.password_validation.CommonPasswordValidator'},
    {'NAME': 'django.contrib.auth.password_validation.UserAttributeSimilarityValidator'},
]
```

## Checklist de vérification post-patch
- [ ] La longueur minimale exigée est d'au moins 12 caractères.
- [ ] Aucune règle de composition artificielle (majuscule/chiffre/spécial obligatoire) ne remplace le contrôle de longueur.
- [ ] Le mot de passe choisi est vérifié contre une liste de mots de passe compromis connus (API k-anonymity type HaveIBeenPwned ou liste locale équivalente).
- [ ] Aucune politique de rotation périodique obligatoire n'est appliquée sans motif de compromission avérée.
- [ ] Un test de non-régression confirme qu'un mot de passe valide de 12+ caractères est accepté et qu'un mot de passe compromis connu est rejeté.
