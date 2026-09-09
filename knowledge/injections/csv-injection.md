---
id: csv-injection
category: injections
cwe: CWE-1236
owasp: A03:2021-Injection
severity_default: medium
languages: [php, js, python, java, csharp, go]
---

# CSV Injection (Formula Injection)

## Description
L'injection CSV survient lorsqu'une application exporte des données utilisateur non neutralisées dans un fichier CSV qui sera ensuite ouvert dans un tableur (Excel, LibreOffice, Google Sheets). Si une cellule commence par un caractère interprété comme déclencheur de formule (`=`, `+`, `-`, `@`), le tableur peut exécuter cette formule à l'ouverture, permettant exfiltration de données locales ou exécution de commandes via des fonctions comme `HYPERLINK` ou des appels DDE.

## Où ça apparaît typiquement
- Fonctionnalités d'export CSV/Excel de données utilisateur (noms, commentaires, adresses).
- Génération de rapports automatisés incluant des champs de formulaire libres.
- Export de logs ou de tickets support contenant du texte non filtré.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Écriture de valeurs utilisateur dans un fichier CSV sans vérifier le premier caractère de chaque cellule.
- Absence de préfixe d'échappement (`'` ou espace protégé) pour les cellules commençant par `=`, `+`, `-`, `@`, tabulation ou retour chariot.
- Utilisation d'une bibliothèque d'export CSV "maison" plutôt qu'une bibliothèque maintenue gérant l'échappement.

## Remédiation
- Préfixer toute cellule dont le premier caractère est `=`, `+`, `-`, `@`, tab ou CR par une apostrophe ou un caractère neutralisant avant écriture.
- Utiliser une bibliothèque CSV reconnue qui applique cet échappement par défaut.
- Ajouter un en-tête `Content-Disposition` clair et sensibiliser les utilisateurs à ne pas activer les macros/liens à l'ouverture de fichiers externes.
- Voir `rules/remediation/csv-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/csv-injection/`.

## Références
- OWASP Cheat Sheet: CSV Injection
- CWE-1236: Improper Neutralization of Formula Elements in a CSV File
