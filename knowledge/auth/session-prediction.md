---
id: session-prediction
category: auth
cwe: CWE-330
owasp: A07:2021-Identification and Authentication Failures
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Session Prediction

## Description
La prédiction de session survient quand les identifiants de session (session ID, token) sont générés selon un algorithme insuffisamment aléatoire ou prévisible (compteur incrémental, timestamp, hash faible d'informations connues). Un attaquant capable de deviner ou de reconstruire un identifiant de session valide peut usurper l'identité d'un autre utilisateur sans connaître son mot de passe.

## Où ça apparaît typiquement
- Génération d'identifiants de session maison, sans passer par le générateur de session sécurisé du framework.
- Identifiants dérivés de valeurs prévisibles (horodatage, ID utilisateur, compteur séquentiel).
- Utilisation d'un générateur de nombres pseudo-aléatoires non cryptographique pour produire des tokens de session ou de réinitialisation.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Génération d'identifiant de session par concaténation/hash de valeurs prévisibles plutôt que par un générateur aléatoire cryptographique dédié.
- Utilisation d'une fonction de génération aléatoire non cryptographique (ex. générateur standard non sécurisé) pour un identifiant sensible.
- Espace de valeurs des identifiants de session trop restreint (longueur insuffisante).

## Remédiation
- Générer les identifiants de session exclusivement via un générateur de nombres aléatoires cryptographiquement sûr, avec une entropie suffisante (128 bits minimum).
- Utiliser systématiquement les mécanismes de gestion de session fournis par le framework plutôt qu'une implémentation maison.
- Renouveler l'identifiant de session à chaque changement de niveau de privilège (connexion, élévation).
- Voir `rules/remediation/session-prediction.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/session-prediction/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Session Management
- CWE-330: Use of Insufficiently Random Values
