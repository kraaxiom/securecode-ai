---
id: malicious-packages
category: supply-chain
cwe: CWE-1357
owasp: A08:2021-Software-and-Data-Integrity-Failures
severity_default: critical
languages: [js, python, php, java, csharp, go, rust]
---

# Paquets malveillants (supply chain compromise)

## Description
Un paquet malveillant est un composant publié sur un registre public (npm, PyPI, RubyGems, crates.io, Packagist, etc.) contenant intentionnellement du code nuisible : vol d'identifiants, backdoor, mineur de cryptomonnaie ou exfiltration de données. Il peut s'agir d'un paquet créé de toutes pièces, d'un paquet légitime dont le mainteneur a été compromis, ou d'une mise à jour malveillante injectée dans une dépendance de confiance. L'installation se fait généralement de façon totalement transparente lors d'un `install` classique.

## Où ça apparaît typiquement
- Dépendances directes ou transitives ajoutées récemment sans revue de sécurité.
- Scripts de cycle de vie du gestionnaire de paquets (`postinstall`, `preinstall`) exécutés automatiquement à l'installation.
- Dépendances mises à jour automatiquement sans contrôle (auto-merge de bots de mise à jour sans vérification).
- Paquets provenant de mainteneurs individuels sans authentification forte ni signature de publication.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Scripts d'installation (`postinstall`/`preinstall`) exécutant du code réseau ou obscurci (encodage base64, `eval` dynamique).
- Changement brutal de mainteneur ou de dépôt source d'un paquet entre deux versions.
- Paquet sans code source public correspondant, ou dépôt source récemment créé/vide.
- Absence de vérification de checksum/signature lors de l'installation en CI/CD.
- Pic anormal de permissions demandées par une dépendance (accès réseau, système de fichiers) sans justification fonctionnelle.

## Remédiation
- Auditer les scripts de cycle de vie avant d'autoriser leur exécution (désactiver `postinstall` par défaut si possible).
- Utiliser des outils de scan de paquets malveillants dans la CI/CD en complément du SCA classique.
- Épingler les versions et vérifier les hachages/signatures des paquets publiés.
- Limiter le nombre de mainteneurs de confiance et exiger une revue manuelle avant mise à jour majeure des dépendances critiques.
- Voir `rules/remediation/malicious-packages.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/malicious-packages/`.

## Références
- OWASP Top 10: A08:2021 – Software and Data Integrity Failures
- CWE-1357: Reliance on Insufficiently Trustworthy Component
