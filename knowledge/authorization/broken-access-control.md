---
id: broken-access-control
category: authorization
cwe: CWE-284
owasp: A01:2021-Broken Access Control
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Broken Access Control (contrôle d'accès défaillant, catégorie générale)

## Description
Le contrôle d'accès défaillant regroupe l'ensemble des défauts qui permettent à un utilisateur d'accéder à des ressources ou d'effectuer des actions au-delà de ses permissions prévues. Il s'agit de la catégorie la plus large de l'OWASP Top 10, englobant des variantes spécifiques comme l'IDOR, la BOLA, la BFLA ou l'escalade de privilèges, mais aussi des cas plus génériques où aucune vérification d'autorisation n'est effectuée du tout sur une ressource sensible.

## Où ça apparaît typiquement
- Endpoints d'API ou de pages n'appliquant aucune vérification d'autorisation, en s'appuyant uniquement sur l'authentification.
- Contrôles d'accès effectués côté client (masquage d'éléments d'interface) sans vérification équivalente côté serveur.
- Règles d'autorisation dupliquées et incohérentes entre plusieurs couches (contrôleur, service, base de données).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Endpoint sensible (lecture/écriture/suppression) protégé uniquement par un middleware d'authentification, sans vérification d'autorisation propre à la ressource ou à l'action.
- Logique d'autorisation dupliquée dans plusieurs couches du code avec des règles divergentes.
- Absence de tests automatisés couvrant les cas d'accès refusé (utilisateur non autorisé sur une ressource d'autrui).

## Remédiation
- Centraliser la logique d'autorisation dans une couche unique et systématiquement appliquée (middleware, policy, guard), plutôt que de la disperser.
- Appliquer le principe du "deny by default" : tout accès est refusé sauf autorisation explicite.
- Écrire des tests automatisés dédiés aux cas de refus d'accès, pas seulement aux cas de succès.
- Voir `rules/remediation/broken-access-control.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/broken-access-control/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Authorization
- CWE-284: Improper Access Control
