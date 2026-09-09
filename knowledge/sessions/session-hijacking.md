---
id: session-hijacking
category: sessions
cwe: CWE-613
owasp: A07:2021-Identification and Authentication Failures
severity_default: high
languages: [php, js, python, java, csharp, go]
---

# Session Hijacking

## Description
Le détournement de session (session hijacking) désigne l'utilisation par un attaquant d'un identifiant de session valide appartenant à un autre utilisateur pour usurper son identité sans avoir besoin de ses identifiants de connexion. Il résulte généralement d'une combinaison de faiblesses : cookies mal protégés, sessions qui n'expirent jamais, absence de liaison de la session au contexte de la requête (IP, user-agent), ou fuite de l'identifiant via des canaux non sécurisés. Contrairement à la fixation de session, l'attaquant n'impose pas l'identifiant à l'avance : il l'obtient ou le devine après coup.

## Où ça apparaît typiquement
- Sessions configurées avec une durée de vie illimitée ou excessivement longue, sans expiration ni renouvellement périodique.
- Absence de invalidation de session côté serveur lors de la déconnexion (logout côté client uniquement).
- Identifiants de session exposés dans les logs applicatifs, les URLs, ou des messages d'erreur.
- Absence de contrôle de cohérence minimal entre la session et le contexte de la requête (changement brutal d'IP/user-agent non détecté).
- Transmission de l'identifiant de session sur un canal non chiffré (voir aussi `missing-secure`).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Configuration de durée de vie de session absente ou définie sur une valeur très longue (`session.gc_maxlifetime`, `maxAge`, équivalent) sans mécanisme de renouvellement.
- Fonction de déconnexion qui supprime uniquement le cookie côté client sans appeler la destruction de session côté serveur (`session_destroy`, `req.session.destroy`).
- Journalisation (logs) contenant l'identifiant de session en clair.
- Absence de tout mécanisme de détection d'anomalie associé à la session (changement d'agent utilisateur, d'IP) dans le code de validation de session.

## Remédiation
- Fixer une durée de vie de session raisonnable et renouveler l'identifiant périodiquement ainsi qu'après toute action sensible.
- Invalider la session côté serveur (pas seulement le cookie) à la déconnexion.
- Ne jamais journaliser l'identifiant de session en clair.
- Protéger le transport (`HttpOnly`, `Secure`, `SameSite`) pour réduire les canaux de fuite.
- Envisager une liaison légère de la session au contexte de la requête, avec gestion prudente des faux positifs.
- Voir `rules/remediation/session-hijacking.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js/session-hijacking/`.

## Références
- OWASP Cheat Sheet: Session Management Cheat Sheet
- CWE-613: Insufficient Session Expiration
