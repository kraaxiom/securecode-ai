# Remédiation — DOM-based Cross-Site Scripting (DOM XSS)

## Principe
Ne jamais écrire une donnée provenant d'une source contrôlable par l'attaquant (`location.hash`, `location.search`, `document.referrer`, `postMessage`, `localStorage`) dans un sink DOM qui interprète du HTML ou exécute du code (`innerHTML`, `document.write`, `eval`, `Function()`). Préférer les API texte et valider systématiquement l'origine des messages entrants.

## JavaScript (client-side / DOM)
```js
// Avant — vulnérable
const query = new URLSearchParams(location.search).get('q');
document.getElementById('result').innerHTML = query;

// Après — sécurisé
const query = new URLSearchParams(location.search).get('q');
document.getElementById('result').textContent = query;
```

```js
// Avant — vulnérable
window.addEventListener('message', (event) => {
  document.body.innerHTML = event.data.html;
});

// Après — sécurisé
window.addEventListener('message', (event) => {
  if (event.origin !== 'https://trusted.example.com') return;
  document.getElementById('content').textContent = event.data.text;
});
```

```js
// Avant — vulnérable — HTML dynamique nécessaire
element.innerHTML = untrustedHtml;

// Après — sécurisé — sanitisation avant insertion
import DOMPurify from 'dompurify';
element.innerHTML = DOMPurify.sanitize(untrustedHtml, { USE_PROFILES: { html: true } });
```

## Node.js (rendu serveur associé)
```js
// Avant — vulnérable — injection de données non fiables dans un script inline
res.send(`<script>const data = "${req.query.q}";</script>`);

// Après — sécurisé — passage par un attribut data-* échappé et lu côté client
const escaped = encodeURIComponent(req.query.q);
res.send(`<div id="app" data-query="${escaped}"></div>`);
```

## PHP (angle serveur — éviter d'injecter des données dans du JS inline)
```php
// Avant — vulnérable
echo "<script>var q = '" . $_GET['q'] . "';</script>";

// Après — sécurisé — donnée passée en JSON échappé, pas de concaténation directe
echo "<script>var q = " . json_encode($_GET['q'], JSON_HEX_TAG | JSON_HEX_APOS | JSON_HEX_QUOT) . ";</script>";
```

## Python (Flask — angle serveur)
```python
# Avant — vulnérable
return f"<script>var q = '{request.args.get('q')}';</script>"

# Après — sécurisé
import json
from markupsafe import Markup

safe_json = json.dumps(request.args.get('q', ''))
return Markup(f"<script>var q = {safe_json};</script>")
```

## Checklist de vérification post-patch
- [ ] Aucun sink dangereux (`innerHTML`, `outerHTML`, `document.write`, `insertAdjacentHTML`, `eval`, `Function()`) n'reçoit directement une source contrôlable par l'utilisateur (`location.*`, `document.referrer`, `postMessage`, `localStorage`).
- [ ] Tout gestionnaire `message` sur `window` vérifie explicitement `event.origin` avant de traiter la charge utile.
- [ ] Le HTML dynamique strictement nécessaire passe par une bibliothèque de sanitisation DOM maintenue avant insertion.
- [ ] Un test de non-régression confirme que la donnée légitime (ex: `?q=texte`) s'affiche toujours correctement.
- [ ] Une Content Security Policy est en place en défense en profondeur (restriction de `script-src`).
