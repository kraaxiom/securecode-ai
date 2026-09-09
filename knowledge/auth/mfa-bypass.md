---
id: mfa-bypass
category: auth
cwe: CWE-287
owasp: A07:2021-Identification and Authentication Failures
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Contournement de l'authentification multi-facteurs (MFA Bypass)

## Description
Un contournement de MFA survient quand la logique applicative permet à un utilisateur d'atteindre un état authentifié sans avoir validé le second facteur requis. Cela peut résulter d'un flux d'authentification mal séquencé (état de session marqué "authentifié" avant la vérification du second facteur), d'un endpoint alternatif oublié, ou d'une vérification du second facteur reposant uniquement sur un paramètre modifiable côté client.

## Où ça apparaît typiquement
- Flux d'authentification en plusieurs étapes où la session ou le token est émis avant la validation du code MFA.
- Endpoints d'API legacy ou mobiles ne repassant pas par le même flux MFA que l'interface web principale.
- Vérification du second facteur basée sur un paramètre de requête (`mfa_verified=true`) plutôt que sur un état serveur.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Émission d'un token/session pleinement privilégié avant l'étape de vérification du second facteur, avec une étape MFA "additionnelle" facultative en apparence.
- Existence de plusieurs endpoints d'authentification (web, mobile, API) n'appliquant pas tous la même politique MFA.
- État "MFA validé" stocké côté client (cookie, paramètre) plutôt que dans une session serveur infalsifiable.

## Remédiation
- Ne délivrer un token/session pleinement privilégié qu'après validation complète et côté serveur du second facteur.
- Utiliser un état de session intermédiaire ("authentification partielle") distinct et à privilèges limités entre le premier et le second facteur.
- Appliquer la même politique MFA de façon cohérente sur tous les points d'entrée (web, mobile, API).
- Voir `rules/remediation/mfa-bypass.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/mfa-bypass/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Multifactor Authentication
- CWE-287: Improper Authentication
