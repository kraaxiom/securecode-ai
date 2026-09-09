---
id: known-cves
category: supply-chain
cwe: CWE-1104
owasp: A06:2021-Vulnerable-and-Outdated-Components
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Utilisation de composants avec CVE connues

## Description
Une application peut intégrer une bibliothèque, un framework ou un composant tiers pour lequel une vulnérabilité publique (CVE) a déjà été identifiée et documentée. Contrairement à une vulnérabilité inconnue (0-day), le risque ici est connu, souvent accompagné d'un exploit public, ce qui rend l'exploitation triviale pour un attaquant qui identifie la version du composant utilisé. Le maintien de dépendances non patchées est l'une des causes les plus fréquentes de compromissions.

## Où ça apparaît typiquement
- Fichiers de gestion de dépendances (`composer.lock`, `package-lock.json`, `requirements.txt`, `pom.xml`, `go.sum`, `Cargo.lock`) figeant une version ancienne d'un composant.
- Images de conteneurs construites sur une base système ou des paquets non mis à jour.
- Frameworks ou librairies tierces embarquées manuellement (vendoring) sans suivi de version.
- Absence de pipeline de mise à jour automatisée des dépendances.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Version de dépendance figée correspondant à une plage connue comme vulnérable dans les bases publiques (NVD, GitHub Advisory Database, OSV).
- Absence de fichier de verrouillage à jour ou dépendances sans contrainte de version minimale de sécurité.
- Absence d'outil de scan de dépendances (SCA) dans la CI/CD.
- Composants tiers non maintenus depuis longtemps (absence de release récente) toujours en production.

## Remédiation
- Intégrer un outil d'analyse de composition logicielle (SCA) dans la CI/CD (ex: audit natif du gestionnaire de paquets, scanners dédiés).
- Mettre à jour régulièrement les dépendances et suivre les avis de sécurité (GitHub Advisories, NVD, OSV.dev).
- Définir une politique de gestion du cycle de vie des dépendances (versions maximales autorisées d'ancienneté, SLA de patch).
- Isoler les composants non patchables via des contrôles compensatoires (WAF, sandboxing) en attendant la mise à jour.
- Voir `rules/remediation/known-cves.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/known-cves/`.

## Références
- OWASP Top 10: A06:2021 – Vulnerable and Outdated Components
- CWE-1104: Use of Unmaintained Third Party Components
