---
id: vulnerable-libraries
category: supply-chain
cwe: CWE-1035
owasp: A06:2021-Vulnerable-and-Outdated-Components
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Bibliothèques vulnérables ou obsolètes

## Description
Ce pattern couvre l'usage plus général de bibliothèques présentant des faiblesses de sécurité structurelles : composants non maintenus, versions obsolètes n'ayant pas reçu de correctifs depuis longtemps, ou bibliothèques historiquement connues pour des pratiques peu sûres (désérialisation non sécurisée, cryptographie faible par défaut). Il se distingue de `known-cves` en ce qu'il englobe aussi le risque diffus lié à l'ancienneté et au manque de maintenance, même sans CVE publiée précise.

## Où ça apparaît typiquement
- Dépendances majeures obsolètes (versions majeures dépassées de plusieurs générations).
- Bibliothèques abandonnées par leurs mainteneurs (dernier commit très ancien, dépôt archivé).
- Frameworks legacy maintenus en interne sans mise à niveau vers les versions supportées.
- Dépendances transitives profondes jamais auditées directement.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Version de dépendance très en retard par rapport à la dernière version stable disponible.
- Absence de mise à jour dans le fichier de verrouillage depuis une période prolongée.
- Dépôt source de la dépendance marqué comme archivé/non maintenu.
- Absence de politique de renouvellement des dépendances (pas de bot de mise à jour, pas de revue périodique).

## Remédiation
- Mettre en place un inventaire des composants (SBOM) et un suivi de leur état de maintenance.
- Planifier des mises à jour régulières et remplacer les bibliothèques abandonnées par des alternatives maintenues.
- Prioriser la mise à jour selon la criticité d'exposition (composant exposé côté internet vs interne).
- Automatiser la détection avec un scanner SCA intégré à la CI/CD.
- Voir `rules/remediation/vulnerable-libraries.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/vulnerable-libraries/`.

## Références
- OWASP Top 10: A06:2021 – Vulnerable and Outdated Components
- CWE-1035: 2017 Top 25 - Use of Known Vulnerable Component (catégorie générique de composants vulnérables/obsolètes)
