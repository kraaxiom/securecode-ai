# Remédiation — Prototype Pollution avancée

## Principe
Filtrer explicitement les clés dangereuses (`__proto__`, `constructor`, `prototype`) dans toute fonction de fusion/assignation récursive, utiliser `Object.create(null)` ou des `Map` pour les objets construits à partir d'entrées utilisateur, et maintenir à jour les bibliothèques de fusion/clonage.

## Node.js (fusion maison)
```js
// Avant — vulnérable : fusion récursive sans filtrage de clés
function merge(target, source) {
  for (const key in source) {
    if (typeof source[key] === 'object' && source[key] !== null) {
      target[key] = target[key] || {};
      merge(target[key], source[key]);
    } else {
      target[key] = source[key];
    }
  }
  return target;
}
merge(config, JSON.parse(userInput));

// Après — sécurisé : filtrage explicite des clés dangereuses + objet sans prototype
const DANGEROUS_KEYS = new Set(['__proto__', 'constructor', 'prototype']);

function safeMerge(target, source) {
  for (const key of Object.keys(source)) {
    if (DANGEROUS_KEYS.has(key)) continue;
    if (typeof source[key] === 'object' && source[key] !== null) {
      target[key] = target[key] && typeof target[key] === 'object'
        ? target[key] : Object.create(null);
      safeMerge(target[key], source[key]);
    } else {
      target[key] = source[key];
    }
  }
  return target;
}
safeMerge(config, JSON.parse(userInput));
```

## Node.js (bibliothèque tierce)
```js
// Avant — vulnérable : version non patchée
// package.json: "lodash": "4.17.15"
_.merge(config, userInput);

// Après — sécurisé : version patchée + validation de schéma en amont
// package.json: "lodash": "^4.17.21"
const Ajv = require('ajv');
const ajv = new Ajv();
const valid = ajv.validate(configSchema, userInput);
if (!valid) throw new Error('Invalid config payload');
_.merge(config, userInput);
```

## Node.js (désérialisation JSON directe dans un objet partagé)
```js
// Avant — vulnérable : JSON.parse fusionné dans un état global partagé entre requêtes
app.post('/settings', (req, res) => {
  Object.assign(globalSettings, JSON.parse(req.body.settings));
});

// Après — sécurisé : objet par requête, sans prototype, clés validées
app.post('/settings', (req, res) => {
  const payload = JSON.parse(req.body.settings);
  const safe = Object.create(null);
  for (const key of Object.keys(payload)) {
    if (DANGEROUS_KEYS.has(key)) continue;
    if (ALLOWED_SETTINGS_KEYS.has(key)) safe[key] = payload[key];
  }
  applyPerRequestSettings(req, safe);
});
```

## Checklist de vérification post-patch
- [ ] Toute fonction de fusion/assignation récursive exclut explicitement `__proto__`, `constructor` et `prototype`.
- [ ] Les objets construits à partir d'entrées utilisateur utilisent `Object.create(null)` ou une `Map`.
- [ ] Les bibliothèques de fusion/clonage (`lodash`, `deepmerge`, etc.) sont à jour vers des versions corrigées.
- [ ] Un test confirme qu'un payload contenant `{"__proto__": {"polluted": true}}` ne modifie pas `Object.prototype` après traitement.
