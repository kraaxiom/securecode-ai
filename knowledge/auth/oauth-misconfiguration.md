---
id: oauth-misconfiguration
category: auth
cwe: CWE-287
owasp: A07:2021-Identification and Authentication Failures
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Mauvaise configuration OAuth / OpenID Connect

## Description
Les intégrations OAuth 2.0 et OpenID Connect impliquent plusieurs points de configuration critiques (redirect_uri, state, validation du jeton d'identité) dont une mauvaise implémentation ouvre la voie à des attaques de vol de code d'autorisation, de fixation d'état ou d'usurpation d'identité. Une mauvaise configuration côté client (application relying party) ou fournisseur peut permettre à un attaquant de récupérer un jeton d'accès destiné à une victime ou de se connecter à sa place.

## Où ça apparaît typiquement
- Enregistrement d'application OAuth avec une liste de `redirect_uri` trop permissive (wildcard, sous-domaines non maîtrisés).
- Flux d'autorisation sans paramètre `state` (ou non vérifié) exposant à une CSRF sur le callback d'authentification.
- Vérification incomplète du jeton d'identité (ID token) : émetteur, audience ou expiration non contrôlés.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- `redirect_uri` validé par simple préfixe ou motif large plutôt que par correspondance exacte avec une liste blanche.
- Absence de génération et de vérification d'un paramètre `state` (ou `nonce` pour OIDC) unique par requête d'autorisation.
- Vérification du jeton d'identité reçu du fournisseur sans contrôle de l'émetteur (`iss`), de l'audience (`aud`) ou de la signature.

## Remédiation
- Enregistrer et valider les `redirect_uri` par correspondance exacte, sans wildcard.
- Générer un `state` (et un `nonce` en OIDC) unique et imprévisible par requête, et le vérifier strictement au retour.
- Valider intégralement le jeton d'identité reçu (signature, émetteur, audience, expiration) avant de faire confiance à son contenu.
- Voir `rules/remediation/oauth-misconfiguration.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/csharp-dotnet/oauth-misconfiguration/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: OAuth2
- CWE-287: Improper Authentication
