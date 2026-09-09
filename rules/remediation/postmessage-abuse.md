# Remédiation — Abus de l'API postMessage

## Principe
Toujours vérifier `event.origin` contre une liste blanche explicite avant de traiter un message reçu, et toujours préciser un `targetOrigin` exact (jamais `*`) lors de l'envoi de données sensibles.

## JS (réception — récepteur)
```js
// Avant — vulnérable
window.addEventListener('message', (event) => {
  updateUserData(event.data); // aucune vérification d'origine
});

// Après — sécurisé
const TRUSTED_ORIGINS = new Set(['https://partner.example.com']);
window.addEventListener('message', (event) => {
  if (!TRUSTED_ORIGINS.has(event.origin)) {
    return; // message ignoré, origine non fiable
  }
  if (typeof event.data !== 'object' || event.data === null || !('type' in event.data)) {
    return; // structure inattendue, rejetée
  }
  updateUserData(event.data);
});
```

## JS (émission)
```js
// Avant — vulnérable
iframe.contentWindow.postMessage(sensitiveData, '*');

// Après — sécurisé
iframe.contentWindow.postMessage(sensitiveData, 'https://partner.example.com');
```

## Popup OAuth/SSO
```js
// Avant — vulnérable
popup.postMessage({ token }, '*');
window.addEventListener('message', (e) => handleAuthCallback(e.data));

// Après — sécurisé
const AUTH_ORIGIN = 'https://auth.example.com';
popup.postMessage({ token }, AUTH_ORIGIN);
window.addEventListener('message', (e) => {
  if (e.origin !== AUTH_ORIGIN) return;
  handleAuthCallback(e.data);
});
```

## Checklist de vérification post-patch
- [ ] Chaque gestionnaire `message` vérifie `event.origin` contre une liste blanche explicite avant tout traitement.
- [ ] La comparaison d'origine se fait par égalité stricte, jamais par sous-chaîne (`includes`).
- [ ] Chaque appel `postMessage` envoyant des données sensibles précise un `targetOrigin` exact, jamais `*`.
- [ ] La structure et le type de `event.data` sont validés avant tout traitement (pas de `eval`/insertion DOM directe).
- [ ] Un test confirme qu'un message provenant d'une origine non listée est bien ignoré.
