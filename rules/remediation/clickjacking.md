# Remédiation — Clickjacking

## Principe
Définir `X-Frame-Options` et/ou la directive CSP `frame-ancestors` sur toutes les réponses HTML sensibles, pour empêcher l'intégration de la page dans un cadre tiers non autorisé.

## PHP (natif)
```php
// Avant — vulnérable
// aucun en-tête anti-framing envoyé

// Après — sécurisé
header('X-Frame-Options: DENY');
header("Content-Security-Policy: frame-ancestors 'none'");
```

## PHP (Laravel middleware)
```php
// Avant — vulnérable
// pas de middleware de sécurité global

// Après — sécurisé
class SecurityHeaders
{
    public function handle($request, Closure $next)
    {
        $response = $next($request);
        $response->headers->set('X-Frame-Options', 'SAMEORIGIN');
        $response->headers->set('Content-Security-Policy', "frame-ancestors 'self'");
        return $response;
    }
}
```

## Node.js (Express + helmet)
```js
// Avant — vulnérable
app.use(express.static('public')); // pas d'en-tête anti-framing

// Après — sécurisé
const helmet = require('helmet');
app.use(helmet.frameguard({ action: 'deny' }));
app.use(helmet.contentSecurityPolicy({
  directives: { frameAncestors: ["'none'"] },
}));
```

## Python (Django)
```python
# Avant — vulnérable
# settings.py sans middleware de protection clickjacking

# Après — sécurisé
MIDDLEWARE = [
    'django.middleware.clickjacking.XFrameOptionsMiddleware',
    ...
]
X_FRAME_OPTIONS = 'DENY'
```

## Checklist de vérification post-patch
- [ ] `X-Frame-Options: DENY` ou `SAMEORIGIN` est présent sur toutes les pages sensibles.
- [ ] La directive CSP `frame-ancestors` est définie et cohérente avec `X-Frame-Options`.
- [ ] Si l'intégration en iframe est un besoin métier réel, `frame-ancestors` restreint explicitement aux origines de confiance.
- [ ] Un test confirme qu'une tentative d'intégration de la page dans un iframe cross-origin est bloquée par le navigateur.
