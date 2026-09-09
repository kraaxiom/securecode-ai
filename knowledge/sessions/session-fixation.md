---
id: session-fixation
category: sessions
cwe: CWE-384
owasp: A07:2021-Identification and Authentication Failures
severity_default: high
languages: [php, js, python, java, csharp, go]
---

# Session Fixation

## Description
La fixation de session survient quand une application accepte un identifiant de session fourni ou connu avant l'authentification, puis authentifie l'utilisateur sans régénérer cet identifiant. Un attaquant peut ainsi imposer à une victime un identifiant de session qu'il connaît déjà (via un lien, un cookie déposé au préalable, ou un paramètre d'URL), puis, une fois la victime authentifiée, réutiliser ce même identifiant pour accéder à la session authentifiée. Le problème central est l'absence de renouvellement de l'identifiant de session lors d'un changement de niveau de privilège.

## Où ça apparaît typiquement
- Logique d'authentification qui réutilise l'identifiant de session existant après un login réussi, au lieu d'en générer un nouveau.
- Frameworks où la régénération de session après authentification n'est pas activée par défaut.
- Acceptation d'un identifiant de session transmis via un paramètre d'URL ou un champ de formulaire, plutôt que généré uniquement côté serveur.
- Applications multi-domaines/sous-domaines qui partagent un cookie de session sans contrôle strict de son origine.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fonction de login qui ne fait pas appel à un mécanisme de régénération d'identifiant de session (`session_regenerate_id`, `req.session.regenerate`, équivalent du framework) après authentification réussie.
- Identifiant de session accepté depuis une source externe (paramètre GET/POST) sans validation qu'il provient bien du serveur.
- Absence de changement d'identifiant de session lors d'une élévation de privilège (login, changement de rôle).

## Remédiation
- Régénérer systématiquement l'identifiant de session immédiatement après une authentification réussie ou un changement de niveau de privilège.
- Invalider l'ancienne session côté serveur plutôt que de simplement en créer une nouvelle en parallèle.
- Ne jamais accepter un identifiant de session provenant d'un paramètre d'URL.
- Voir `rules/remediation/session-fixation.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/session-fixation/`.

## Références
- OWASP Cheat Sheet: Session Management Cheat Sheet
- CWE-384: Session Fixation
