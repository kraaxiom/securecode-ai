---
id: ldap-injection
category: injections
cwe: CWE-90
owasp: A03:2021-Injection
severity_default: high
languages: [php, js, python, java, csharp, go]
---

# LDAP Injection

## Description
L'injection LDAP survient lorsqu'une entrée utilisateur non neutralisée est intégrée dans un filtre de recherche LDAP ou un DN (Distinguished Name), permettant à un attaquant de modifier la logique du filtre. Cela peut conduire à un contournement d'authentification, à une extraction non autorisée d'informations de l'annuaire, ou à une élévation de privilèges selon les droits du compte de service utilisé pour la connexion LDAP.

## Où ça apparaît typiquement
- Formulaires d'authentification s'appuyant sur un annuaire LDAP/Active Directory avec construction dynamique du filtre de recherche.
- Fonctionnalités de recherche d'utilisateurs ou de groupes dans un annuaire, où le terme recherché est inséré tel quel dans le filtre.
- Construction dynamique d'un DN à partir d'une entrée utilisateur pour des opérations de bind ou de modification.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Concaténation de chaînes pour construire un filtre LDAP (`(uid=" + input + ")`).
- Absence d'utilisation des fonctions d'échappement dédiées (échappement des caractères spéciaux LDAP: `* ( ) \ NUL`).
- Construction d'un DN par concaténation directe d'une valeur utilisateur.

## Remédiation
- Utiliser les fonctions d'échappement fournies par la bibliothèque LDAP du langage pour les valeurs de filtre et les composants de DN.
- Valider le format attendu des entrées (ex: identifiant alphanumérique) avant construction du filtre.
- Appliquer le principe du moindre privilège sur le compte de service utilisé pour interroger l'annuaire.
- Voir `rules/remediation/ldap-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/java-spring/ldap-injection/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: LDAP Injection Prevention
- CWE-90: Improper Neutralization of Special Elements used in an LDAP Query ('LDAP Injection')
