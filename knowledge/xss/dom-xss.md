---
id: dom-xss
category: xss
cwe: CWE-79
owasp: A03:2021-Injection
severity_default: high
languages: [js]
---

# DOM-based Cross-Site Scripting (DOM XSS)

## Description
Le XSS basé sur le DOM est une variante où la vulnérabilité réside entièrement dans le code JavaScript côté client, sans que la charge malveillante transite nécessairement par le serveur. Une donnée provenant d'une source contrôlable par l'attaquant (URL, `document.referrer`, `localStorage`, message `postMessage`) est écrite dans un puits (sink) du DOM qui interprète du HTML ou exécute du code, sans passer par un mécanisme de sécurité du navigateur adapté.

## Où ça apparaît typiquement
- Lecture de `location.hash`, `location.search` ou `document.URL` réinjectée via `innerHTML`, `document.write` ou `eval`.
- Gestionnaires d'événements `postMessage` qui traitent le contenu du message sans validation de l'origine ni du format.
- Frameworks SPA manipulant directement le DOM (hors du système de binding sécurisé) avec des données issues de l'URL ou du stockage local.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Utilisation de sinks dangereux (`innerHTML`, `outerHTML`, `document.write`, `insertAdjacentHTML`, `eval`, `Function()`) avec une valeur dérivée d'une source contrôlable par l'utilisateur.
- Écouteur `message` sur `window` sans vérification de `event.origin`.
- Utilisation de `location.hash`/`location.search` directement dans une opération de rendu DOM sans passage par les API sûres du framework (binding textuel).

## Remédiation
- Préférer les API texte (`textContent`, binding de framework) aux sinks HTML lorsque le contenu n'a pas besoin d'être interprété comme du markup.
- Si du HTML dynamique est nécessaire, le passer par une bibliothèque de sanitisation DOM reconnue avant insertion.
- Valider systématiquement `event.origin` dans tout gestionnaire `postMessage`.
- Voir `rules/remediation/dom-xss.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/nodejs-express/dom-xss/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: DOM based XSS Prevention
- CWE-79: Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')
