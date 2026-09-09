---
id: session-replay
category: sessions
cwe: CWE-294
owasp: A07:2021-Identification and Authentication Failures
severity_default: high
languages: [php, js, python, java, csharp, go]
---

# Session Replay

## Description
Une attaque par rejeu (replay) consiste à capturer un identifiant de session ou un jeton d'authentification valide, puis à le réutiliser tel quel plus tard pour obtenir un accès non autorisé, sans avoir besoin de le déchiffrer ni de connaître les identifiants de connexion. Le problème sous-jacent est l'absence de mécanisme rendant un jeton capturé inutilisable après usage ou après un certain temps : pas de nonce, pas d'horodatage vérifié, pas de liaison à une utilisation unique. C'est une menace distincte du hijacking au sens large : elle cible spécifiquement l'absence de protection contre la réutilisation d'une preuve d'authentification déjà observée.

## Où ça apparaît typiquement
- Jetons d'authentification ou de session sans horodatage ni nonce vérifiable côté serveur.
- API acceptant un jeton signé mais ne vérifiant jamais sa date d'expiration (`exp`) ni son unicité.
- Liens de connexion "magiques" (magic links) ou jetons de réinitialisation de mot de passe réutilisables plusieurs fois.
- Mécanismes d'authentification par requête signée sans protection anti-rejeu (pas de nonce à usage unique côté serveur).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Vérification de jeton (JWT ou autre) qui ne contrôle pas les champs `exp`/`iat`/`nbf` ou équivalent.
- Jeton à usage unique (reset password, magic link) qui n'est pas marqué comme consommé après sa première utilisation.
- Absence de stockage côté serveur d'un nonce ou d'un identifiant d'utilisation pour invalider un jeton déjà présenté.
- Logique d'authentification qui ne compare jamais l'horodatage de la requête à une fenêtre de validité.

## Remédiation
- Inclure et vérifier systématiquement une expiration courte (`exp`) sur tout jeton d'authentification.
- Marquer les jetons à usage unique comme consommés dès leur première utilisation valide, côté serveur.
- Utiliser un nonce ou un identifiant unique par requête pour les mécanismes sensibles, invalidé après consommation.
- Privilégier des canaux chiffrés (HTTPS) pour réduire le risque de capture initiale.
- Voir `rules/remediation/session-replay.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js/session-replay/`.

## Références
- OWASP Cheat Sheet: Session Management Cheat Sheet
- CWE-294: Authentication Bypass by Capture-replay
