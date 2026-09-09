# Remédiation — Web Cache Deception

## Principe
Ne mettre en cache que sur la base du `Content-Type` réel de la réponse, jamais sur l'apparence de l'URL. Ajouter `Cache-Control: private, no-store` sur toute réponse contenant des données utilisateur, et configurer le routeur pour retourner une 404 stricte sur les chemins non reconnus.

## PHP (routeur maison)
```php
// Avant — vulnérable : segments de chemin arbitraires acceptés après une route valide
if (str_starts_with($path, '/compte/profil')) {
    header('Cache-Control: public, max-age=3600'); // hérité d'une règle générique par erreur
    echo renderProfile($user);
}

// Après — sécurisé : correspondance exacte de route + cache explicite
if ($path === '/compte/profil') {
    header('Cache-Control: private, no-store');
    echo renderProfile($user);
} else {
    http_response_code(404);
    exit('Not Found');
}
```

## Node.js (Express)
```js
// Avant — vulnérable : Express ignore silencieusement les extensions suffixées sur certaines routes
app.get('/account/profile', (req, res) => {
  res.render('profile', { user: req.user });
});

// Après — sécurisé : route stricte + en-tête explicite, extensions non reconnues renvoient 404
app.get('/account/profile', (req, res) => {
  res.set('Cache-Control', 'private, no-store');
  res.render('profile', { user: req.user });
});
app.use((req, res) => res.status(404).send('Not Found')); // catch-all strict
```

## Configuration cache/CDN (Nginx)
```nginx
# Avant — vulnérable : mise en cache basée sur l'extension apparente de l'URL
location ~* \.(css|js|jpg|png)$ {
    proxy_cache my_cache;
    proxy_pass http://backend;
}

# Après — sécurisé : cache conditionné au Content-Type réel de la réponse
location / {
    proxy_pass http://backend;
    proxy_cache my_cache;
    proxy_cache_valid 200 1h;
    proxy_no_cache $upstream_http_cache_control ~* "no-store|private";
    proxy_ignore_headers Set-Cookie;
}
```

## Checklist de vérification post-patch
- [ ] Le cache/CDN se base sur le `Content-Type` réel de la réponse et non sur l'extension de l'URL.
- [ ] Toute réponse authentifiée/personnalisée porte `Cache-Control: private, no-store`.
- [ ] Le routeur applicatif retourne une 404 stricte pour tout segment de chemin non reconnu après une route valide.
- [ ] Test : une requête vers `/compte/profil/inexistant.css` (ou équivalent) renvoie 404 et n'est pas mise en cache.
