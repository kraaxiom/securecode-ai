# Remédiation — Universal Cross-Site Scripting (UXSS)

## Principe
Restreindre strictement les sources de script autorisées via une Content Security Policy, utiliser Subresource Integrity pour toute dépendance tierce chargée depuis un CDN, et isoler les composants tiers non indispensables dans des iframes avec attribut `sandbox` minimal. La cause première est souvent hors du code applicatif, mais l'application reste responsable de limiter sa surface d'exposition.

## JavaScript (client-side — chargement de script tiers)
```js
// Avant — vulnérable — script tiers chargé sans intégrité vérifiée ni CSP
const s = document.createElement('script');
s.src = 'https://cdn.example.com/widget.js';
document.head.appendChild(s);

// Après — sécurisé — Subresource Integrity + attribut crossorigin
const s = document.createElement('script');
s.src = 'https://cdn.example.com/widget.js';
s.integrity = 'sha384-<hash-du-fichier-verifie>';
s.crossOrigin = 'anonymous';
document.head.appendChild(s);
```

```html
<!-- Avant — vulnérable — iframe tierce sans sandbox -->
<iframe src="https://widget-tiers.example.com/chat"></iframe>

<!-- Après — sécurisé — sandbox restrictif limité au strict nécessaire -->
<iframe
  src="https://widget-tiers.example.com/chat"
  sandbox="allow-scripts allow-same-origin"
  referrerpolicy="no-referrer">
</iframe>
```

## Node.js (Express — en-têtes de sécurité applicatifs)
```js
// Avant — vulnérable — aucune CSP, aucune Permissions-Policy définie
app.use((req, res, next) => next());

// Après — sécurisé — CSP stricte + Permissions-Policy en défense en profondeur
const helmet = require('helmet');
app.use(helmet.contentSecurityPolicy({
  directives: {
    defaultSrc: ["'self'"],
    scriptSrc: ["'self'", 'https://cdn.example.com'],
    frameSrc: ["'self'", 'https://widget-tiers.example.com'],
  },
}));
app.use((req, res, next) => {
  res.setHeader('Permissions-Policy', 'camera=(), microphone=(), geolocation=()');
  next();
});
```

## PHP (en-têtes de sécurité côté serveur)
```php
// Avant — vulnérable
// Aucun en-tête de sécurité émis

// Après — sécurisé
header("Content-Security-Policy: default-src 'self'; script-src 'self' https://cdn.example.com; frame-src 'self' https://widget-tiers.example.com");
header("Permissions-Policy: camera=(), microphone=(), geolocation=()");
```

## Python (Flask — en-têtes de sécurité via after_request)
```python
# Avant — vulnérable
# Aucun en-tête de sécurité émis

# Après — sécurisé
@app.after_request
def set_security_headers(response):
    response.headers['Content-Security-Policy'] = (
        "default-src 'self'; script-src 'self' https://cdn.example.com; "
        "frame-src 'self' https://widget-tiers.example.com"
    )
    response.headers['Permissions-Policy'] = 'camera=(), microphone=(), geolocation=()'
    return response
```

## Checklist de vérification post-patch
- [ ] Toute dépendance JavaScript tierce chargée depuis un CDN utilise l'attribut `integrity` (Subresource Integrity) et `crossorigin`.
- [ ] Une Content Security Policy stricte restreint `script-src` et `frame-src` aux origines strictement nécessaires.
- [ ] Toute iframe tierce non indispensable dispose d'un attribut `sandbox` limité au strict nécessaire.
- [ ] Une politique `Permissions-Policy` limite les capacités des composants tiers embarqués.
- [ ] Les navigateurs ciblés et les composants tiers intégrés sont suivis et maintenus à jour (veille des avis de sécurité).
