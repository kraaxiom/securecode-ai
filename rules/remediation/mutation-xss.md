# Remédiation — Mutation-based Cross-Site Scripting (mXSS)

## Principe
Utiliser une bibliothèque de sanitisation HTML activement maintenue et connue pour couvrir les vecteurs de mutation liés au parsing du navigateur. Limiter au strict minimum les allers-retours entre représentation chaîne et DOM pour du contenu non fiable, et appliquer une Content Security Policy stricte en défense en profondeur.

## JavaScript (client-side / DOM)
```js
// Avant — vulnérable — sanitisation faite sur une chaîne puis réinsérée telle quelle,
// exposée aux quirks de reparsing du navigateur (aller-retour string -> DOM -> string)
function sanitizeNaive(html) {
  const div = document.createElement('div');
  div.innerHTML = html;
  return div.innerHTML; // ré-sérialisation qui peut "muter" le markup
}
target.innerHTML = sanitizeNaive(untrustedHtml);

// Après — sécurisé — bibliothèque de sanitisation maintenue, tenant compte des mutations connues
import DOMPurify from 'dompurify';
target.innerHTML = DOMPurify.sanitize(untrustedHtml);
```

```js
// Avant — vulnérable — éditeur riche avec allers-retours multiples DOM <-> chaîne
editor.innerHTML = sanitizeOnce(storedContent);
const saved = editor.innerHTML; // re-sérialisé sans re-sanitisation avant stockage

// Après — sécurisé — re-sanitisation à chaque étape de sérialisation
import DOMPurify from 'dompurify';
editor.innerHTML = DOMPurify.sanitize(storedContent);
const saved = DOMPurify.sanitize(editor.innerHTML);
```

## PHP (angle serveur — sanitisation avant stockage)
```php
// Avant — vulnérable — sanitisation via une bibliothèque obsolète non tenue à jour
$clean = strip_tags($userHtml, '<b><i><a>');
saveToDatabase($clean);

// Après — sécurisé — bibliothèque de sanitisation HTML maintenue (HTMLPurifier)
$config = HTMLPurifier_Config::createDefault();
$purifier = new HTMLPurifier($config);
$clean = $purifier->purify($userHtml);
saveToDatabase($clean);
```

## Python (Django/Flask — bleach à jour)
```python
# Avant — vulnérable — bibliothèque de sanitisation ancienne et mal configurée
import bleach

clean = bleach.clean(user_html, tags=[], strip=False)  # tags vides mais strip=False -> fuite possible

# Après — sécurisé — configuration stricte, bibliothèque tenue à jour
import bleach

clean = bleach.clean(user_html, tags=['p', 'b', 'i', 'a'], attributes={'a': ['href']}, strip=True)
```

## Checklist de vérification post-patch
- [ ] La bibliothèque de sanitisation HTML utilisée est activement maintenue et à jour (vérifier le CVE tracker du projet).
- [ ] Les allers-retours entre chaîne HTML et DOM pour du contenu non fiable sont réduits au strict minimum.
- [ ] Toute nouvelle sérialisation du contenu (avant sauvegarde ou réaffichage) déclenche une nouvelle passe de sanitisation.
- [ ] Une Content Security Policy stricte (`script-src 'self'`, pas de `unsafe-inline`) est active en défense en profondeur.
- [ ] Un test de non-régression avec des cas connus de mutation (balises imbriquées, entités mal formées) confirme l'absence de contournement.
