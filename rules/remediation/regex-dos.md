# Remédiation — ReDoS (Regular Expression Denial of Service)

## Principe
Réécrire les expressions régulières pour éliminer l'ambiguïté de correspondance (éviter les quantificateurs imbriqués), imposer une limite de longueur sur l'entrée avant application de la regex, et utiliser un timeout d'exécution ou un moteur à complexité linéaire garantie (ex: RE2) pour les entrées non fiables.

## PHP (PCRE)
```php
// Avant — vulnérable : quantificateurs imbriqués sur entrée non bornée
$email = $_POST['email'];
if (preg_match('/^([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)+$/', $email)) {
    // ...
}

// Après — sécurisé : regex simplifiée + limite de longueur + timeout PCRE
$email = substr($_POST['email'], 0, 254); // limite de longueur
ini_set('pcre.backtrack_limit', '100000');
if (preg_match('/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/', $email)) {
    // ...
}
```

## JavaScript / Node.js
```js
// Avant — vulnérable : groupes répétés imbriqués, aucune limite de taille
function isValidEmail(input) {
  return /^([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)+$/.test(input);
}
isValidEmail(req.body.email);

// Après — sécurisé : regex non ambiguë + limite de taille avant test
function isValidEmail(input) {
  if (typeof input !== 'string' || input.length > 254) return false;
  return /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(input);
}
isValidEmail(req.body.email);
```

## Python (re)
```python
# Avant — vulnérable : quantificateurs imbriqués
import re

def is_valid_email(value):
    return re.match(r'^([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)+$', value) is not None

is_valid_email(request.form['email'])

# Après — sécurisé : regex non ambiguë + limite de longueur en amont
import re

EMAIL_RE = re.compile(r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$')

def is_valid_email(value: str) -> bool:
    if len(value) > 254:
        return False
    return EMAIL_RE.match(value) is not None
```

## Checklist de vérification post-patch
- [ ] Aucune expression régulière du fichier corrigé ne contient de quantificateurs imbriqués ou d'alternance ambiguë sur une entrée utilisateur.
- [ ] Une limite de longueur est imposée sur l'entrée avant l'application de la regex.
- [ ] Un timeout ou un moteur à complexité linéaire garantie protège l'exécution de la regex sur une entrée non fiable.
- [ ] Un test de performance confirme un temps d'exécution constant sur une entrée pathologique (chaîne longue répétitive).
