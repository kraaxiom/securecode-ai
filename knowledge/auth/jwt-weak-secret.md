---
id: jwt-weak-secret
category: auth
cwe: CWE-326
owasp: A02:2021-Cryptographic Failures
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# JWT signé avec un secret faible

## Description
Quand un token JWT est signé avec un algorithme symétrique (HMAC, ex. `HS256`) et que le secret utilisé est court, prévisible ou issu d'une liste de mots courants, un attaquant peut retrouver ce secret hors ligne par attaque par dictionnaire ou force brute, puis forger des tokens valides avec n'importe quel contenu, y compris des rôles privilégiés.

## Où ça apparaît typiquement
- Secret JWT défini en dur dans le code source ou dans un fichier de configuration versionné, avec une valeur courte ou par défaut.
- Variable d'environnement de secret JWT non générée aléatoirement (valeur d'exemple de la documentation laissée telle quelle).
- Applications ne faisant pas tourner (rotation) leur secret JWT après un incident ou un changement d'équipe.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Secret de signature JWT court, en dur dans le code, ou correspondant à une valeur d'exemple documentée publiquement.
- Absence de génération cryptographiquement aléatoire du secret (longueur insuffisante par rapport à l'algorithme utilisé).
- Absence de procédure de rotation du secret JWT.

## Remédiation
- Générer le secret JWT de façon cryptographiquement aléatoire, avec une entropie suffisante par rapport à l'algorithme (au minimum 256 bits pour HS256).
- Stocker le secret dans un gestionnaire de secrets dédié, jamais en dur dans le code ou en clair dans le dépôt.
- Mettre en place une rotation régulière du secret et une procédure de révocation en cas de compromission suspectée.
- Envisager un algorithme asymétrique (RS256/ES256) pour séparer clé de signature et clé de vérification.
- Voir `rules/remediation/jwt-weak-secret.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/java-spring/jwt-weak-secret/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: JSON Web Token for Java
- CWE-326: Inadequate Encryption Strength
