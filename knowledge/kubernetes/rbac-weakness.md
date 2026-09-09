---
id: rbac-weakness
category: kubernetes
cwe: CWE-269
owasp: A01:2021-Broken Access Control
severity_default: high
languages: []
---

# Faiblesse RBAC (Role/ClusterRole trop permissifs)

## Description
Le contrôle d'accès basé sur les rôles (RBAC) de Kubernetes définit qui peut faire quoi sur quelles ressources. Une politique RBAC trop permissive (verbes ou ressources en wildcard, liaison de `cluster-admin` à des comptes non nécessaires) viole le principe du moindre privilège et permet à un compte compromis ou à un utilisateur mal intentionné d'agir bien au-delà de son besoin réel. Ce risque est de niveau infrastructure (manifests `Role`, `ClusterRole`, `RoleBinding`, `ClusterRoleBinding`) et indépendant de tout langage de programmation.

## Où ça apparaît typiquement
- Manifests `ClusterRole`/`Role` définissant les règles d'accès (`rules`).
- `ClusterRoleBinding`/`RoleBinding` associant des utilisateurs, groupes ou ServiceAccounts à un rôle.
- Charts Helm d'applications tierces créant leurs propres rôles avec des permissions larges "pour simplifier".
- ServiceAccounts par défaut de namespace auxquels des permissions supplémentaires ont été attachées.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- `rules` contenant `verbs: ["*"]` ou `resources: ["*"]` ou `apiGroups: ["*"]`.
- `ClusterRoleBinding` liant un ServiceAccount applicatif au rôle `cluster-admin`.
- Permissions sur des ressources sensibles (`secrets`, `pods/exec`, `clusterroles`, `clusterrolebindings`) accordées à un rôle à portée large.
- `RoleBinding` référençant le ServiceAccount `default` d'un namespace avec des droits élevés.
- Verbe `impersonate` ou `escalate` accordé sans justification claire.

## Remédiation
- Appliquer le principe du moindre privilège : définir des `rules` scoped à des ressources et verbes précis, jamais de wildcard par défaut.
- Éviter de lier `cluster-admin` à des ServiceAccounts applicatifs ; créer des rôles dédiés et namespacés.
- Auditer régulièrement les bindings RBAC (`kubectl auth can-i --list`, outils comme `rbac-lookup` ou `kubectl-who-can`).
- Voir `rules/remediation/rbac-weakness.md` pour des exemples de rôles scoped.

## Exemple avant/après
Voir `examples/k8s/rbac-weakness/`.

## Références
- CIS Kubernetes Benchmark: section 5.1 "RBAC and Service Accounts"
- CWE-269: Improper Privilege Management
