---
id: postmessage-abuse
category: frontend
cwe: CWE-346
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: [js, ts]
---

# Abus de l'API postMessage

## Description
L'API `postMessage` permet la communication entre fenêtres/iframes d'origines différentes. Un abus survient quand l'émetteur envoie des messages avec un `targetOrigin` trop permissif (`*`), ou quand le récepteur traite les messages entrants sans vérifier la propriété `event.origin`. Dans ce dernier cas, n'importe quelle page (y compris une page malveillante ouverte dans un onglet ou un iframe) peut envoyer des messages forgés que l'application traite comme légitimes, menant à de la falsification de données, du vol d'information ou de l'exécution d'actions non autorisées.

## Où ça apparaît typiquement
- Gestionnaires `window.addEventListener('message', handler)` qui traitent `event.data` sans vérifier `event.origin`.
- Widgets tiers intégrés en iframe (chat, paiement, authentification) échangeant des données sensibles via postMessage.
- Communication entre une fenêtre principale et une popup d'authentification OAuth/SSO.
- Envoi de messages avec `window.postMessage(data, '*')` au lieu de préciser l'origine cible exacte.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Gestionnaire d'événement `message` ne contenant aucune comparaison de `event.origin` contre une valeur attendue.
- Utilisation de `postMessage(data, '*')` pour envoyer des données sensibles.
- Vérification d'origine faible basée sur une inclusion de sous-chaîne (`event.origin.includes('monsite.com')`) plutôt qu'une égalité stricte.
- Traitement de `event.data` directement comme du code ou une commande (ex: passage à `eval` ou insertion dans le DOM) sans validation de structure.

## Remédiation
- Toujours vérifier `event.origin` contre une liste blanche explicite d'origines de confiance avant de traiter un message reçu.
- Toujours préciser un `targetOrigin` exact (jamais `*`) lors de l'envoi de données sensibles via `postMessage`.
- Valider strictement la structure et le type de `event.data` avant tout traitement.
- Voir `rules/remediation/postmessage-abuse.md` pour les diffs par langage/framework.

## Exemple avant/après
Voir `examples/js/postmessage-abuse/`.

## Références
- OWASP Cheat Sheet: HTML5 Security Cheat Sheet (postMessage)
- CWE-346: Origin Validation Error
