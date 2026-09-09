---
id: missing-referrer-policy
category: headers
cwe: CWE-200
owasp: A01:2021-Broken Access Control
severity_default: low
languages: [php, js, python, java, csharp, go, rust]
---

# Referrer-Policy manquant

## Description
Sans en-tête `Referrer-Policy`, le navigateur applique un comportement par défaut qui peut transmettre l'URL complète de la page courante (y compris des paramètres sensibles présents en query string : tokens, identifiants de session, identifiants personnels) dans l'en-tête `Referer` lors de la navigation vers un site tiers ou du chargement de ressources externes. Ces données fuient alors vers des domaines externes (analytics, CDN, liens sortants) sans contrôle.

## Où ça apparaît typiquement
- Applications transmettant des données sensibles dans l'URL (tokens de réinitialisation de mot de passe, identifiants de session) sans en-tête `Referrer-Policy`.
- Pages contenant des liens sortants ou des ressources tierces (images, scripts) chargées depuis des domaines externes.
- Configuration serveur web ou framework n'ayant jamais défini de politique de référent.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- En-tête `Referrer-Policy` absent des réponses HTTP.
- Présence de données sensibles (tokens, identifiants) dans les paramètres d'URL de pages contenant des liens ou ressources externes.

## Remédiation
- Ajouter `Referrer-Policy: strict-origin-when-cross-origin` (ou `no-referrer` pour un contrôle maximal) sur toutes les réponses.
- Éviter de placer des données sensibles dans l'URL ; préférer les headers, le corps de requête ou des tokens à usage unique déjà consommés côté serveur.
- Voir `rules/remediation/missing-referrer-policy.md`.

## Exemple avant/après
Voir `examples/js/missing-referrer-policy/`.

## Références
- OWASP Secure Headers Project
- MDN Web Docs: Referrer-Policy
- CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
