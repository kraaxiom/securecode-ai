# Remédiation — Abus de Service Worker

## Principe
N'enregistrer des service workers qu'à partir de chemins statiques codés en dur, jamais dérivés d'une entrée utilisateur. Restreindre le scope au strict périmètre nécessaire et ajouter une CSP `worker-src 'self'`.

## JS (enregistrement)
```js
// Avant — vulnérable
const swUrl = new URLSearchParams(location.search).get('sw');
navigator.serviceWorker.register(swUrl); // URL contrôlée par l'attaquant

// Après — sécurisé
navigator.serviceWorker.register('/static/sw.js', { scope: '/app/' });
```

## Scope restreint
```js
// Avant — vulnérable
navigator.serviceWorker.register('/sw.js', { scope: '/' }); // trop large

// Après — sécurisé
navigator.serviceWorker.register('/app/sw.js', { scope: '/app/' }); // limité au périmètre nécessaire
```

## CSP restreignant worker-src (backend PHP servant l'en-tête)
```php
// Avant — vulnérable
// aucun worker-src défini, CSP absente ou trop large

// Après — sécurisé
header("Content-Security-Policy: worker-src 'self'; script-src 'self'");
```

## Cache API — ne pas mettre en cache des réponses non validées
```js
// Avant — vulnérable
self.addEventListener('fetch', (event) => {
  event.respondWith(
    fetch(event.request).then((res) => {
      caches.open('v1').then((cache) => cache.put(event.request, res.clone()));
      return res;
    })
  );
});

// Après — sécurisé
self.addEventListener('fetch', (event) => {
  event.respondWith(
    fetch(event.request).then((res) => {
      if (res.ok && res.type === 'basic') { // même origine, réponse valide
        caches.open('v1').then((cache) => cache.put(event.request, res.clone()));
      }
      return res;
    })
  );
});
```

## Checklist de vérification post-patch
- [ ] Aucun appel `serviceWorker.register()` n'utilise une URL dérivée d'une entrée utilisateur.
- [ ] Le scope du service worker est restreint au périmètre fonctionnel strictement nécessaire (pas `/` par défaut).
- [ ] Une CSP `worker-src 'self'` (et `script-src` cohérent) est en place.
- [ ] La Cache API ne met en cache que des réponses de même origine validées (`res.ok`, `res.type === 'basic'`).
- [ ] Toute XSS existante est traitée en priorité critique dans une application utilisant des service workers.
