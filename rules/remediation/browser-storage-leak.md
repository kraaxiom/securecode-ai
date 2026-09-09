# Remédiation — Fuite de données via le stockage navigateur

## Principe
Ne jamais placer de données sensibles (jetons de session, mots de passe, secrets d'API, données personnelles) dans un stockage accessible en JavaScript (`localStorage`, `sessionStorage`, `IndexedDB`) ou dans un cookie sans protections. Préférer des cookies `HttpOnly`/`Secure`/`SameSite` pour les sessions, et garder les secrets côté serveur.

## JavaScript (frontend vanilla)
```js
// Avant — vulnérable
async function login(credentials) {
  const res = await fetch('/api/login', { method: 'POST', body: JSON.stringify(credentials) });
  const { token, refreshToken } = await res.json();
  localStorage.setItem('token', token);
  localStorage.setItem('refreshToken', refreshToken);
}

// Après — sécurisé
// Le serveur pose directement un cookie de session HttpOnly/Secure/SameSite
// à la réponse de /api/login. Le frontend ne manipule plus le jeton.
async function login(credentials) {
  await fetch('/api/login', {
    method: 'POST',
    credentials: 'include', // envoie/reçoit le cookie de session
    body: JSON.stringify(credentials),
  });
  // Aucune donnée sensible stockée côté client.
}
```

## Node.js (Express — configuration serveur/headers)
```js
// Avant — vulnérable
app.post('/api/login', (req, res) => {
  const token = signToken(req.user);
  res.json({ token }); // le frontend devra le stocker lui-même (localStorage)
});

// Après — sécurisé
app.post('/api/login', (req, res) => {
  const token = signToken(req.user);
  res.cookie('session', token, {
    httpOnly: true,
    secure: true,
    sameSite: 'strict',
    maxAge: 15 * 60 * 1000,
  });
  res.json({ ok: true });
});
```

## PHP (headers/configuration serveur)
```php
// Avant — vulnérable
$token = generateToken($user);
echo json_encode(['token' => $token]); // stocké côté client par le JS

// Après — sécurisé
$token = generateToken($user);
setcookie('session', $token, [
    'expires'  => time() + 900,
    'path'     => '/',
    'secure'   => true,
    'httponly' => true,
    'samesite' => 'Strict',
]);
echo json_encode(['ok' => true]);
```

## Checklist de vérification post-patch
- [ ] Aucun jeton d'authentification ou secret n'est écrit via `localStorage.setItem`/`sessionStorage.setItem`.
- [ ] Les cookies de session portent bien les attributs `HttpOnly`, `Secure` et `SameSite`.
- [ ] Le stockage client est purgé à la déconnexion (`localStorage.clear()`/`sessionStorage.clear()` pour les données non sensibles restantes).
- [ ] Aucun secret backend (clé API privée, secret de signature) n'apparaît dans le bundle JS livré au navigateur.
