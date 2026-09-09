---
id: service-worker-abuse
category: frontend
cwe: CWE-494
owasp: A08:2021-Software and Data Integrity Failures
severity_default: high
languages: [js, ts]
---

# Abus de Service Worker

## Description
Un service worker s'exécute avec un pouvoir important : il peut intercepter, modifier ou mettre en cache toutes les requêtes réseau de son scope, et persister même après la fermeture de l'onglet. Un abus survient quand un service worker est enregistré depuis une source non fiable, chargé sans vérification d'intégrité, ou lorsque son scope est trop large, permettant à un script injecté (ex: via XSS) d'installer un service worker malveillant qui intercepte durablement le trafic, vole des données ou sert du contenu falsifié même après correction de la faille initiale.

## Où ça apparaît typiquement
- Enregistrement de service worker (`navigator.serviceWorker.register(...)`) avec une URL de script construite dynamiquement à partir d'une entrée utilisateur.
- Scope de service worker défini trop largement (`/`) alors que seule une sous-section de l'application en a besoin.
- Absence de Subresource Integrity ou de vérification d'origine sur le script de service worker chargé.
- Application vulnérable à une XSS permettant l'injection d'un appel `register()` vers un script contrôlé par l'attaquant.
- Cache du service worker (`Cache API`) alimenté par des réponses non validées, pouvant persister du contenu falsifié.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à `navigator.serviceWorker.register(url)` où `url` provient d'une variable non constante ou dérivée d'une entrée utilisateur.
- Scope de service worker déclaré à la racine (`{ scope: '/' }`) sans justification métier.
- Absence de Content-Security-Policy restreignant `worker-src`/`script-src` pour limiter l'origine des service workers.
- Mise en cache dans la Cache API de réponses sans validation de leur intégrité ou de leur origine.

## Remédiation
- N'enregistrer des service workers qu'à partir de chemins statiques codés en dur, jamais dérivés d'une entrée utilisateur.
- Restreindre le scope du service worker au strict périmètre fonctionnel nécessaire.
- Ajouter une directive CSP `worker-src 'self'` pour empêcher l'enregistrement de workers depuis une autre origine.
- Traiter toute XSS comme critique dans une application utilisant des service workers, car elle peut mener à une persistance longue durée de code malveillant.
- Voir `rules/remediation/service-worker-abuse.md` pour les diffs par langage/framework.

## Exemple avant/après
Voir `examples/js/service-worker-abuse/`.

## Références
- OWASP Cheat Sheet: HTML5 Security Cheat Sheet
- CWE-494: Download of Code Without Integrity Check
