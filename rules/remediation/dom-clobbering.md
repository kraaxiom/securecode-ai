# Remédiation — DOM Clobbering

## Principe
Toujours valider le type des objets globaux avant utilisation plutôt que de faire confiance à leur simple présence. Sanitiser strictement le HTML utilisateur, notamment les attributs `id`/`name`.

## JS (validation de type avant usage)
```js
// Avant — vulnérable
if (window.appConfig) {
  fetch(window.appConfig.apiUrl); // appConfig peut être un <a id="appConfig"> injecté
}

// Après — sécurisé
if (window.appConfig && typeof window.appConfig === 'object' && !(window.appConfig instanceof HTMLElement)) {
  fetch(window.appConfig.apiUrl);
} else {
  throw new Error('Configuration invalide ou absente');
}
```

## JS (éviter la dépendance à des IDs DOM globaux)
```js
// Avant — vulnérable
const apiUrl = document.getElementById('apiUrl').value; // écrasable par du HTML injecté

// Après — sécurisé
// Configuration chargée depuis un module JS encapsulé, jamais depuis le DOM
import { API_URL } from './config.js';
const apiUrl = API_URL;
```

## Sanitisation HTML (DOMPurify) restreignant id/name
```js
// Avant — vulnérable
element.innerHTML = DOMPurify.sanitize(userHtml); // config par défaut autorise id/name

// Après — sécurisé
element.innerHTML = DOMPurify.sanitize(userHtml, {
  FORBID_ATTR: ['id', 'name'],
});
```

## PHP (sanitisation côté serveur avant rendu)
```php
// Avant — vulnérable
echo $userSubmittedHtml; // id/name arbitraires autorisés

// Après — sécurisé
$config = HTMLPurifier_Config::createDefault();
$config->set('Attr.ForbiddenClasses', []);
$config->set('HTML.ForbiddenAttributes', ['id', 'name']);
$purifier = new HTMLPurifier($config);
echo $purifier->purify($userSubmittedHtml);
```

## Checklist de vérification post-patch
- [ ] Toute variable globale implicite (`window.X`) sensible est validée par type (`instanceof`, vérification de forme) avant usage.
- [ ] Le code applicatif ne dépend plus de `document.getElementById`/`document.forms` pour de la configuration sensible.
- [ ] La sanitisation HTML restreint ou neutralise les attributs `id`/`name` réservés à l'usage applicatif.
- [ ] Un test confirme qu'un élément HTML injecté avec `id="appConfig"` ne peut pas écraser la configuration attendue.
