---
id: weak-session-id
category: sessions
cwe: CWE-330
owasp: A07:2021-Identification and Authentication Failures
severity_default: high
languages: [php, js, python, java, csharp, go]
---

# Identifiant de session faible

## Description
Un identifiant de session doit être imprévisible pour un attaquant, faute de quoi il peut être deviné ou reconstitué par force brute. Ce problème apparaît quand l'identifiant est généré à partir d'une source de hasard non cryptographique, d'un espace de valeurs trop restreint, ou d'informations prévisibles comme un compteur, un horodatage ou des données utilisateur (nom, email, ID incrémental). Une fois l'algorithme ou le motif de génération compris, un attaquant peut prédire ou énumérer des identifiants de session valides sans jamais intercepter de trafic.

## Où ça apparaît typiquement
- Génération manuelle d'identifiant de session avec des fonctions aléatoires non cryptographiques (`Math.random`, `rand()`, `random.random`).
- Identifiant de session construit par concaténation de données prévisibles (timestamp, ID utilisateur, compteur incrémental) éventuellement hashées.
- Espace d'identifiants trop court (peu de caractères ou d'entropie) rendant une énumération réaliste.
- Réimplémentation "maison" de la gestion de session au lieu d'utiliser le générateur intégré du framework.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à un générateur pseudo-aléatoire non cryptographique (`Math.random`, `rand`, `mt_rand`, `random` du module standard non-crypto) pour produire un identifiant de session ou un jeton.
- Construction de l'identifiant à partir de valeurs prévisibles concaténées (heure système, ID auto-incrémenté, adresse IP).
- Longueur ou charset de l'identifiant manifestement insuffisant pour offrir une entropie raisonnable.
- Génération de session "custom" en dehors du mécanisme fourni par le framework/langage.

## Remédiation
- Utiliser exclusivement un générateur aléatoire cryptographiquement sûr (`crypto.randomBytes`, `secrets.token_urlsafe`, `random_bytes`, `SecureRandom`, équivalent) pour tout identifiant de session ou jeton.
- S'appuyer sur le mécanisme de gestion de session natif du framework plutôt que de le réimplémenter.
- Garantir une entropie suffisante (longueur et charset) conforme aux recommandations OWASP (au moins 128 bits d'entropie).
- Voir `rules/remediation/weak-session-id.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/weak-session-id/`.

## Références
- OWASP Cheat Sheet: Session Management Cheat Sheet
- CWE-330: Use of Insufficiently Random Values
