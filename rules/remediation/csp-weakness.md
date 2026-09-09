# Remédiation — Faiblesse de Content Security Policy

## Principe
Définir une CSP restrictive basée sur une liste blanche explicite, sans `unsafe-inline`/`unsafe-eval`, avec `object-src 'none'` et `base-uri 'self'`, en mode bloquant (pas uniquement `report-only`).

## PHP (natif)
```php
// Avant — vulnérable
// aucun en-tête CSP envoyé

// Après — sécurisé
header("Content-Security-Policy: default-src 'self'; script-src 'self' 'nonce-{$nonce}'; object-src 'none'; base-uri 'self'; frame-ancestors 'none'");
```

## PHP (Laravel middleware)
```php
// Avant — vulnérable
$response->headers->set('Content-Security-Policy', "default-src *; script-src 'unsafe-inline' 'unsafe-eval'");

// Après — sécurisé
$nonce = base64_encode(random_bytes(16));
$response->headers->set('Content-Security-Policy',
    "default-src 'self'; script-src 'self' 'nonce-$nonce'; object-src 'none'; base-uri 'self'"
);
```

## Node.js (Express + helmet)
```js
// Avant — vulnérable
// pas de CSP, ou CSP en report-only uniquement
app.use(helmet.contentSecurityPolicy({ reportOnly: true, directives: { defaultSrc: ["*"] } }));

// Après — sécurisé
app.use(helmet.contentSecurityPolicy({
  directives: {
    defaultSrc: ["'self'"],
    scriptSrc: ["'self'", (req, res) => `'nonce-${res.locals.nonce}'`],
    objectSrc: ["'none'"],
    baseUri: ["'self'"],
    frameAncestors: ["'none'"],
  },
}));
```

## Python (Django + django-csp)
```python
# Avant — vulnérable
CSP_DEFAULT_SRC = ("*",)
CSP_SCRIPT_SRC = ("'unsafe-inline'", "'unsafe-eval'")

# Après — sécurisé
CSP_DEFAULT_SRC = ("'self'",)
CSP_SCRIPT_SRC = ("'self'",)
CSP_OBJECT_SRC = ("'none'",)
CSP_BASE_URI = ("'self'",)
```

## Checklist de vérification post-patch
- [ ] Un en-tête `Content-Security-Policy` en mode bloquant (pas seulement `report-only`) est présent sur toutes les pages HTML.
- [ ] `unsafe-inline` et `unsafe-eval` sont absents de `script-src` en production (nonces/hashes utilisés à la place).
- [ ] `object-src 'none'` et `base-uri 'self'` sont définis.
- [ ] Aucune directive `default-src`/`script-src`/`style-src` n'utilise un joker `*` trop large.
- [ ] La politique a été testée en `report-only` avant bascule en mode bloquant, sans régression fonctionnelle.
