---
id: firebase-misconfiguration
category: cloud
cwe: CWE-284
owasp: A01:2021-Broken Access Control
severity_default: high
languages: [js]
---

# Mauvaise configuration Firebase (Firestore/Realtime Database/Storage)

## Description
Une mauvaise configuration Firebase survient quand les règles de sécurité (Security Rules) de Firestore, Realtime Database ou Cloud Storage autorisent la lecture et/ou l'écriture sans authentification ni contrôle d'autorisation adéquat. Comme la logique d'accès s'exécute côté client dans de nombreuses applications Firebase, des règles laxistes (souvent héritées du mode "test" par défaut) exposent directement toute la base de données à Internet.

## Où ça apparaît typiquement
- Règles Firestore/RTDB en mode test : `allow read, write: if true;` laissées en production.
- Règles basées uniquement sur `request.auth != null` sans vérifier la propriété de la ressource (tout utilisateur connecté peut lire/modifier les données de tous les autres).
- Cloud Storage avec des règles permissives sur des chemins contenant des fichiers sensibles (uploads utilisateurs, documents).
- Clés d'API Firebase exposées côté client sans restriction (normal pour Firebase, mais dangereux si combiné à des règles laxistes).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fichier `firestore.rules` / `storage.rules` contenant `allow read, write: if true;`.
- Règles sans vérification de propriété (`resource.data.uid == request.auth.uid`) sur des collections utilisateur.
- Absence de règles de validation de schéma/type sur les écritures (permet l'injection de champs arbitraires).
- Configuration Firebase encore en "mode test" au-delà de sa période d'expiration prévue.

## Remédiation
- Écrire des règles de sécurité explicites basées sur l'authentification ET la propriété des données (`request.auth.uid == resource.data.ownerId`).
- Ne jamais déployer en production avec des règles en mode test (`if true`).
- Valider le schéma des données à l'écriture dans les règles pour empêcher l'injection de champs non prévus.
- Auditer régulièrement les règles avec l'émulateur Firebase et des tests automatisés.
- Voir `rules/remediation/firebase-misconfiguration.md`.

## Exemple avant/après
Voir `examples/js/firebase-misconfiguration/`.

## Références
- OWASP Cloud Security Cheat Sheet
- Firebase Docs: Security Rules
- CWE-284: Improper Access Control
