# Remédiation — Prototype Pollution

## Principe
Filtrer explicitement les clés dangereuses (`__proto__`, `constructor`, `prototype`) lors de toute fusion/assignation dynamique d'objets issus d'entrée utilisateur, et préférer `Object.create(null)`/`Map` pour les structures indexées par des clés utilisateur.

## Node.js (fusion d'objets maison)
```js
// Avant — vulnérable
function merge(target, source) {
  for (const key in source) {
    if (typeof source[key] === 'object') {
      target[key] = merge(target[key] || {}, source[key]);
    } else {
      target[key] = source[key];
    }
  }
  return target;
}
merge(config, JSON.parse(req.body.settings)); // clé __proto__ exploitable

// Après — sécurisé
const DANGEROUS_KEYS = new Set(['__proto__', 'constructor', 'prototype']);
function merge(target, source) {
  for (const key of Object.keys(source)) {
    if (DANGEROUS_KEYS.has(key)) continue; // clé dangereuse ignorée
    if (typeof source[key] === 'object' && source[key] !== null) {
      target[key] = merge(target[key] || {}, source[key]);
    } else {
      target[key] = source[key];
    }
  }
  return target;
}
```

## Node.js (lodash.set avec chemin utilisateur)
```js
// Avant — vulnérable
_.set(config, req.body.path, req.body.value); // path = '__proto__.isAdmin'

// Après — sécurisé
const ALLOWED_PATHS = ['theme.color', 'notifications.enabled'];
if (!ALLOWED_PATHS.includes(req.body.path)) {
  return res.status(400).json({ error: 'Chemin non autorisé' });
}
_.set(config, req.body.path, req.body.value);
```

## Store sûr par défaut
```js
// Avant — vulnérable
const store = {}; // objet standard, hérite d'Object.prototype

// Après — sécurisé
const store = Object.create(null); // pas de prototype, __proto__ neutre
// ou : const store = new Map();
```

## Checklist de vérification post-patch
- [ ] Toute fonction de fusion/assignation récursive exclut explicitement `__proto__`, `constructor` et `prototype`.
- [ ] Les bibliothèques de manipulation d'objets (lodash, etc.) sont à jour et ne sont jamais appelées avec un chemin dérivé directement de l'entrée utilisateur sans liste blanche.
- [ ] Les structures indexées par des clés utilisateur utilisent `Object.create(null)` ou `Map`.
- [ ] Un test confirme qu'une entrée contenant `__proto__` ou `constructor.prototype` ne modifie pas `Object.prototype`.
