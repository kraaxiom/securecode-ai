---
id: iam-misconfiguration
category: cloud
cwe: CWE-269
owasp: A01:2021-Broken Access Control
severity_default: high
languages: []
---

# Mauvaise configuration IAM (permissions excessives)

## Description
Une mauvaise configuration IAM survient quand des utilisateurs, rôles ou service accounts cloud (AWS, Azure, GCP) se voient accorder des permissions plus larges que nécessaire pour leur fonction, violant le principe du moindre privilège. Les politiques utilisant des wildcards (`"Action": "*"`, `"Resource": "*"`) ou des rôles administrateur attribués par défaut augmentent considérablement l'impact d'une compromission de compte.

## Où ça apparaît typiquement
- Policies IAM AWS avec `"Effect": "Allow", "Action": "*", "Resource": "*"`.
- Rôles GCP/Azure de type "Owner"/"Contributor" attribués à des service accounts applicatifs au lieu de rôles granulaires.
- Service accounts ou rôles CI/CD disposant de permissions de production alors qu'ils n'exécutent que des tâches de build.
- Trust policies (`AssumeRole`) trop permissives permettant à n'importe quel compte d'assumer un rôle privilégié.
- Absence de rotation/expiration des clés d'accès long-lived (access keys statiques au lieu de rôles temporaires).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Wildcard `*` sur `Action` et/ou `Resource` dans une policy IAM (JSON, Terraform, ARM, Bicep).
- Rôle prédéfini "Owner"/"Admin"/"*Contributor" attribué à un compte de service applicatif.
- Trust relationship (`AssumeRolePolicyDocument`) avec `"Principal": "*"` ou sans condition (`ExternalId`, MFA).
- Clés d'accès statiques (access key/secret key) utilisées à la place de rôles/identités fédérées.

## Remédiation
- Appliquer le principe du moindre privilège : définir des policies granulaires listant les actions et ressources précises nécessaires.
- Préférer les rôles temporaires (STS, Workload Identity Federation) aux clés d'accès statiques.
- Restreindre les trust policies avec des conditions (`ExternalId`, source IP, MFA requis).
- Auditer régulièrement les permissions inutilisées (AWS IAM Access Analyzer, GCP Policy Analyzer) et les révoquer.
- Voir `rules/remediation/iam-misconfiguration.md`.

## Exemple avant/après
Voir `examples/terraform/iam-misconfiguration/`.

## Références
- OWASP Cloud Security Cheat Sheet
- CIS Benchmarks (AWS/Azure/GCP) — section IAM
- CWE-269: Improper Privilege Management
