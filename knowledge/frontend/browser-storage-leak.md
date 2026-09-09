---
id: browser-storage-leak
category: frontend
cwe: CWE-312
owasp: A02:2021-Cryptographic Failures
severity_default: medium
languages: [js, ts]
---

# Fuite de données via le stockage navigateur

## Description
Une fuite de stockage navigateur survient quand une application place des données sensibles (jetons de session, mots de passe, informations personnelles, secrets d'API) dans `localStorage`, `sessionStorage`, `IndexedDB` ou les cookies non protégés, sans tenir compte du fait que ce stockage est accessible en clair à tout script JavaScript exécuté dans la page. En cas de faille XSS, ou simplement via les outils de développement du navigateur, ces données deviennent immédiatement exposées. Contrairement à un cookie `HttpOnly`, rien n'empêche un script tiers ou injecté de lire ce stockage.

## Où ça apparaît typiquement
- Stockage d'un token JWT ou d'un refresh token dans `localStorage` après authentification.
- Mise en cache de données personnelles (email, numéro de téléphone, adresse) côté client pour éviter des appels réseau répétés.
- Sauvegarde de clés d'API tierces ou de secrets de configuration dans `sessionStorage` pour les réutiliser côté frontend.
- Persistance de formulaires contenant des données sensibles (mot de passe, numéro de carte) via `IndexedDB` pour un mode hors-ligne.
- Cookies applicatifs créés sans les attributs `HttpOnly`, `Secure` ou `SameSite`.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appels `localStorage.setItem(...)` ou `sessionStorage.setItem(...)` avec des noms de clés évocateurs (`token`, `jwt`, `auth`, `password`, `apiKey`, `secret`).
- Lecture d'un token d'authentification stocké côté client puis réinjection dans les en-têtes de requêtes (`Authorization: Bearer <valeur lue depuis localStorage>`).
- Création de cookies côté client (`document.cookie = ...`) sans les attributs `HttpOnly`/`Secure`/`SameSite`.
- Absence de chiffrement ou d'obfuscation avant écriture dans `IndexedDB` pour des champs qualifiés de sensibles dans le modèle de données.
- Journaux ou objets de configuration exposant des secrets dans le bundle JS livré au navigateur.

## Remédiation
- Préférer les cookies `HttpOnly`, `Secure` et `SameSite=Strict/Lax` pour les jetons de session, plutôt que le stockage accessible en JavaScript.
- Ne jamais placer de secrets destinés au backend (clés API privées, secrets de signature) dans un bundle ou un stockage côté client.
- Limiter la durée de vie des données sensibles en mémoire (variable JS non persistée) plutôt qu'en stockage persistant.
- Mettre en place une politique de nettoyage systématique du stockage à la déconnexion.
- Voir `rules/remediation/browser-storage-leak.md` pour les diffs par langage/framework.

## Exemple avant/après
Voir `examples/js/browser-storage-leak/`.

## Références
- OWASP Cheat Sheet: HTML5 Security Cheat Sheet (Local Storage)
- OWASP Cheat Sheet: Session Management Cheat Sheet
- CWE-312: Cleartext Storage of Sensitive Information
