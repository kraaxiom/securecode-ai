---
id: formula-injection
category: injections
cwe: CWE-1236
owasp: A03:2021-Injection
severity_default: medium
languages: [php, js, python, java, csharp, go]
---

# Formula Injection

## Description
L'injection de formule est le terme général désignant l'insertion d'une formule de tableur active dans un document exporté (CSV, XLSX, ODS) à partir d'une donnée utilisateur non neutralisée. À l'ouverture du fichier, le tableur exécute la formule, ce qui peut conduire à une exfiltration de données via des requêtes réseau (ex: `WEBSERVICE`, `IMPORTXML`) ou, dans certaines configurations, à une exécution de commande via DDE.

## Où ça apparaît typiquement
- Génération de fichiers Excel/ODS avec des bibliothèques de génération de feuilles de calcul incluant des champs libres utilisateur.
- Export de données CRM, tickets, ou formulaires vers un format tableur.
- Import/réexport de données tierces sans re-neutralisation des cellules.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Écriture directe d'une chaîne utilisateur dans une cellule sans vérifier son premier caractère.
- Bibliothèque de génération de tableur utilisée sans option d'échappement de formule activée.
- Absence de tests de non-régression sur l'export de contenu commençant par `=`, `+`, `-`, `@`.

## Remédiation
- Préfixer toute valeur de cellule commençant par un caractère déclencheur de formule (`=`, `+`, `-`, `@`) d'une apostrophe ou d'un caractère neutre.
- Privilégier les bibliothèques d'export maintenues qui appliquent cet échappement par défaut (ex: options de sécurité des générateurs XLSX modernes).
- Documenter et tester spécifiquement ce cas dans la suite de tests d'export.
- Voir `rules/remediation/formula-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/formula-injection/`.

## Références
- OWASP Cheat Sheet: CSV Injection
- CWE-1236: Improper Neutralization of Formula Elements in a CSV File
