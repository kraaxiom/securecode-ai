---
id: xml-injection
category: injections
cwe: CWE-91
owasp: A03:2021-Injection
severity_default: medium
languages: [php, js, python, java, csharp, go]
---

# XML Injection

## Description
L'injection XML survient lorsqu'une entrée utilisateur non neutralisée est insérée directement dans un document XML sans échappement des caractères spéciaux (`<`, `>`, `&`, guillemets), permettant à un attaquant d'altérer la structure du document (ajout de nœuds, modification d'attributs). Selon le contexte applicatif, cela peut conduire à une falsification de données, un contournement de logique métier basée sur le contenu XML, ou servir de point d'entrée vers d'autres vulnérabilités XML plus graves comme XXE.

## Où ça apparaît typiquement
- Génération manuelle de documents XML par concaténation de chaînes incluant des données utilisateur.
- API SOAP ou export/import de données au format XML construits sans bibliothèque de sérialisation dédiée.
- Fichiers de configuration ou de flux (RSS, SAML) générés dynamiquement avec des valeurs utilisateur non échappées.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Concaténation de chaînes pour construire un document XML incluant une variable utilisateur non échappée.
- Absence d'utilisation d'une bibliothèque de sérialisation XML (DOM, API de génération) au profit de templates de chaînes.
- Insertion de valeurs utilisateur dans des attributs XML sans encodage des guillemets.

## Remédiation
- Utiliser une bibliothèque de sérialisation XML qui échappe automatiquement le contenu des nœuds et attributs, plutôt que de construire le XML par concaténation.
- Valider la structure attendue du document avec un schéma (XSD) après génération.
- Encoder systématiquement les caractères spéciaux XML dans toute valeur insérée dynamiquement.
- Voir `rules/remediation/xml-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/java-spring/xml-injection/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: XML Security Cheat Sheet
- CWE-91: XML Injection (aka Blind XPath Injection)
