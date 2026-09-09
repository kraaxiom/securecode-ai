---
id: circleci-secrets
category: cicd
cwe: CWE-798
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Fuite de secrets dans CircleCI

## Description
Une fuite de secrets CircleCI survient quand des identifiants sont codés en dur dans `.circleci/config.yml`, imprimés dans les logs de build, ou exposés via des contextes ("contexts") partagés avec des projets/branches non fiables. Les workflows déclenchés par des pull requests provenant de forks peuvent, si mal configurés, accéder aux variables d'environnement du projet et donc aux secrets.

## Où ça apparaît typiquement
- Identifiants en clair dans `.circleci/config.yml` au lieu de variables d'environnement du projet/contexte.
- Commandes de debug (`env`, `printenv`, `echo $TOKEN`) affichant des secrets dans les logs de build.
- Contextes CircleCI partagés entre plusieurs projets ou accessibles à des branches non protégées.
- Build déclenché automatiquement sur les PR de forks externes avec accès aux mêmes variables d'environnement que les builds internes.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Motif de secret en clair dans le YAML de configuration CircleCI.
- Étape de job affichant explicitement le contenu d'une variable d'environnement sensible.
- Contexte CircleCI sans restriction de sécurité (aucune approbation requise, pas de scoping par groupe).
- Absence de séparation entre builds "forked PR" et builds internes ayant accès aux secrets.

## Remédiation
- Stocker les secrets dans les "Environment Variables" du projet ou dans un "Context" restreint, jamais en dur dans le YAML.
- Ne jamais journaliser des secrets ; interdire les commandes de dump d'environnement dans les jobs.
- Exiger une approbation manuelle pour les builds provenant de forks avant d'exposer des contextes sensibles.
- Faire tourner (rotation) régulièrement les secrets et limiter leur portée au minimum nécessaire.
- Voir `rules/remediation/circleci-secrets.md`.

## Exemple avant/après
Voir `examples/yaml/circleci-secrets/`.

## Références
- OWASP Cheat Sheet: CI/CD Security
- CircleCI Docs: Security recommendations
- CWE-798: Use of Hard-coded Credentials
