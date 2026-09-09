---
id: regex-dos
category: injections
cwe: CWE-1333
owasp: A03:2021-Injection
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# ReDoS (Regular Expression Denial of Service)

## Description
Le déni de service par expression régulière (ReDoS) survient lorsqu'une expression régulière mal conçue présente un comportement de complexité exponentielle ou polynomiale sur certaines entrées (souvent via des groupes imbriqués avec quantificateurs, ou de l'alternance ambiguë). Un attaquant qui contrôle l'entrée testée par cette expression peut fabriquer une chaîne provoquant un temps de calcul démesuré, bloquant le thread ou le processus qui l'exécute.

## Où ça apparaît typiquement
- Validation de formats utilisateur (email, URL, mot de passe) avec des regex complexes appliquées à une entrée non bornée en taille.
- Parsers ou moteurs de templates utilisant des regex sur du contenu fourni par l'utilisateur.
- Règles de filtrage ou de nettoyage de log/texte appliquées à des flux externes.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Présence de motifs de quantificateurs imbriqués ou chevauchants dans une regex (ex: `(a+)+`, `(a|a)*`, groupes répétés contenant eux-mêmes des quantificateurs).
- Application d'une regex à une entrée utilisateur sans limite de longueur ni timeout d'exécution.
- Absence de test de performance de la regex sur des entrées pathologiques lors de la revue de code.

## Remédiation
- Réécrire les expressions régulières pour éliminer l'ambiguïté de correspondance (éviter les quantificateurs imbriqués).
- Imposer une limite de longueur sur l'entrée avant application de la regex, et un timeout d'exécution si le moteur le permet.
- Utiliser des moteurs de regex à complexité linéaire garantie (ex: RE2) pour les entrées non fiables.
- Voir `rules/remediation/regex-dos.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/nodejs-express/regex-dos/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Denial of Service Cheat Sheet (section ReDoS)
- CWE-1333: Inefficient Regular Expression Complexity
