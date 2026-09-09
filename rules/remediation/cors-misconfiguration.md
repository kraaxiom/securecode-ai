# Remédiation — Mauvaise configuration CORS

## Principe
Définir une liste blanche stricte et statique des origines autorisées, comparée par égalité exacte. Ne jamais combiner reflet dynamique d'origine avec `Access-Control-Allow-Credentials: true`.

## PHP (natif)
```php
// Avant — vulnérable
header('Access-Control-Allow-Origin: ' . $_SERVER['HTTP_ORIGIN']);
header('Access-Control-Allow-Credentials: true');

// Après — sécurisé
$allowedOrigins = ['https://app.example.com'];
$origin = $_SERVER['HTTP_ORIGIN'] ?? '';
if (in_array($origin, $allowedOrigins, true)) {
    header("Access-Control-Allow-Origin: $origin");
    header('Access-Control-Allow-Credentials: true');
}
```

## PHP (Laravel — config/cors.php)
```php
// Avant — vulnérable
'allowed_origins' => ['*'],
'supports_credentials' => true,

// Après — sécurisé
'allowed_origins' => ['https://app.example.com'],
'supports_credentials' => true,
```

## Node.js (Express + cors)
```js
// Avant — vulnérable
app.use(cors({ origin: true, credentials: true })); // reflète toute origine

// Après — sécurisé
const allowedOrigins = ['https://app.example.com'];
app.use(cors({
  origin: (origin, callback) => {
    if (!origin || allowedOrigins.includes(origin)) {
      callback(null, true);
    } else {
      callback(new Error('Origine non autorisée'));
    }
  },
  credentials: true,
}));
```

## Python (Django + django-cors-headers)
```python
# Avant — vulnérable
CORS_ORIGIN_ALLOW_ALL = True
CORS_ALLOW_CREDENTIALS = True

# Après — sécurisé
CORS_ALLOWED_ORIGINS = ["https://app.example.com"]
CORS_ALLOW_CREDENTIALS = True
```

## Checklist de vérification post-patch
- [ ] La liste des origines autorisées est statique et comparée par égalité stricte (pas de regex/sous-chaîne permissive).
- [ ] Aucune combinaison de `Access-Control-Allow-Origin: *` (ou reflet dynamique) avec `Access-Control-Allow-Credentials: true`.
- [ ] Les méthodes et en-têtes autorisés sont restreints au strict nécessaire.
- [ ] Les API publiques sans credentials sont séparées des API authentifiées à CORS restreint.
- [ ] Un test confirme qu'une origine non listée reçoit une réponse CORS refusée.
