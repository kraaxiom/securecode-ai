---
id: dependency-confusion
category: supply-chain
cwe: CWE-1357
owasp: A08:2021-Software-and-Data-Integrity-Failures
severity_default: high
languages: [js, python, java, csharp, go, rust, php]
---

# Dependency Confusion

## Description
La confusion de dépendances exploite les gestionnaires de paquets qui résolvent un nom de package en interrogeant plusieurs registres (privé interne + public). Un attaquant publie sur le registre public un paquet portant le même nom qu'un paquet interne privé, souvent avec un numéro de version plus élevé, pour que le gestionnaire de paquets installe la version malveillante à la place de la version interne légitime. Cette technique a permis des exécutions de code dans des chaînes de build de grandes entreprises.

## Où ça apparaît typiquement
- Paquets internes/privés (npm scope non réservé, packages Python internes, artefacts Maven/NuGet internes) publiés sans namespace protégé.
- Fichiers de configuration de gestionnaire de paquets (`.npmrc`, `pip.conf`, `settings.xml`, `NuGet.Config`) sans registre privé explicitement prioritaire ou sans verrouillage de source.
- Pipelines CI/CD qui installent des dépendances depuis un registre public par défaut alors que certains noms de paquets correspondent à des projets internes.
- Absence de scoping (`@entreprise/paquet`) pour les paquets internes npm.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Nom de paquet interne sans préfixe de scope/namespace réservé sur le registre public correspondant.
- Fichier de configuration du gestionnaire de paquets ne définissant pas explicitement l'ordre de priorité des registres ou n'utilisant pas d'allowlist de sources.
- Absence de verrouillage de version exact (lockfile) ou de vérification d'intégrité (`integrity`/hash) pour les dépendances internes.
- CI/CD sans registre privé configuré en priorité exclusive pour les paquets internes connus.

## Remédiation
- Réserver les noms de paquets internes sur les registres publics correspondants, ou utiliser un scope/namespace dédié non réutilisable.
- Configurer explicitement la priorité et la portée des registres (ex: `.npmrc` avec scope pointant vers le registre privé uniquement).
- Utiliser des lockfiles avec vérification d'intégrité (hash) et épingler les versions.
- Mettre en place un registre privé qui fait office de proxy/miroir vérifié pour toutes les dépendances (internes et publiques).
- Voir `rules/remediation/dependency-confusion.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/dependency-confusion/`.

## Références
- OWASP Cheat Sheet: Software Supply Chain Security
- CWE-1357: Reliance on Insufficiently Trustworthy Component
