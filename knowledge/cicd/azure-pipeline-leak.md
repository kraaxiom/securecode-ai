---
id: azure-pipeline-leak
category: cicd
cwe: CWE-798
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Fuite de secrets dans Azure Pipelines

## Description
Une fuite de secrets Azure Pipelines survient quand des identifiants sont codés en dur dans `azure-pipelines.yml`, affichés dans les logs de build, ou exposés via des groupes de variables/service connections partagés avec des pipelines non fiables. Un pipeline déclenché sur une pull request de fork ayant accès aux service connections de production représente un risque majeur d'exfiltration.

## Où ça apparaît typiquement
- Identifiants en clair dans `azure-pipelines.yml` au lieu de "Variable Groups" ou secrets liés à Azure Key Vault.
- Commande de script (`echo`, `Write-Host`) affichant une variable secrète dans le log.
- Service connections partagées entre plusieurs pipelines sans restriction d'approbation.
- Déclenchement automatique de pipeline sur PR externe avec accès aux mêmes variables que la branche principale.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Motif de secret en clair dans le YAML du pipeline.
- Variable non marquée "secret" (icône cadenas) dans un Variable Group Azure DevOps.
- Étape de script imprimant explicitement le contenu d'une variable sensible.
- Service connection avec accès non restreint ("Grant access permission to all pipelines").

## Remédiation
- Stocker les secrets dans des Variable Groups liés à Azure Key Vault, jamais en dur dans le YAML.
- Marquer chaque variable sensible comme "secret" pour activer le masquage automatique dans les logs.
- Restreindre l'accès aux service connections par pipeline explicite, désactiver l'accès global.
- Exiger une validation pour les builds déclenchés depuis des forks avant d'exposer des secrets.
- Voir `rules/remediation/azure-pipeline-leak.md`.

## Exemple avant/après
Voir `examples/yaml/azure-pipeline-leak/`.

## Références
- OWASP Cheat Sheet: CI/CD Security
- Microsoft Learn: Set secret variables in Azure Pipelines
- CWE-798: Use of Hard-coded Credentials
