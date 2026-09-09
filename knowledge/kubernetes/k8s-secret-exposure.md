---
id: k8s-secret-exposure
category: kubernetes
cwe: CWE-312
owasp: A02:2021-Cryptographic Failures
severity_default: high
languages: []
---

# Exposition de Secrets Kubernetes

## Description
Les objets `Secret` de Kubernetes encodent leurs valeurs en base64, ce qui n'est pas un chiffrement — quiconque a accès au manifest ou à l'API peut trivialement décoder la valeur en clair. Ce risque est de niveau infrastructure (manifests YAML, configuration etcd, RBAC) et ne dépend d'aucun langage de programmation applicatif. Le problème s'aggrave quand des secrets sont commités en clair dans des dépôts Git, stockés sans chiffrement au repos dans etcd, ou accessibles à des rôles RBAC trop permissifs.

## Où ça apparaît typiquement
- Manifests `Secret` versionnés dans un dépôt Git (y compris via Helm values ou Kustomize).
- Configuration etcd du cluster sans chiffrement au repos (`EncryptionConfiguration` absente).
- ConfigMaps utilisés à tort pour stocker des données sensibles (mots de passe, clés API) au lieu de Secrets.
- Variables d'environnement de pods référençant directement des valeurs sensibles en clair (`env.value` au lieu de `envFrom.secretKeyRef`).
- Logs ou sorties `kubectl describe`/`kubectl get -o yaml` incluant des secrets en clair dans des pipelines CI/CD.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Champ `data` d'un objet `Secret` contenant des valeurs base64 commitées dans un dépôt versionné.
- Absence de ressource `EncryptionConfiguration` référencée dans la configuration du kube-apiserver (`--encryption-provider-config`).
- Utilisation de `ConfigMap` pour des clés nommées `password`, `token`, `api_key`, `secret`.
- `RoleBinding`/`ClusterRoleBinding` accordant `get`/`list` sur la ressource `secrets` à un large ensemble de comptes de service.
- Valeurs de secrets injectées en clair dans `env:` d'un manifest de Pod/Deployment plutôt que via `secretKeyRef`.

## Remédiation
- Ne jamais committer de manifests `Secret` avec des valeurs réelles ; utiliser un gestionnaire externe (Vault, sealed-secrets, SOPS, cloud KMS).
- Activer le chiffrement au repos d'etcd (`EncryptionConfiguration` avec un provider KMS ou AES).
- Restreindre l'accès RBAC à la ressource `secrets` au strict nécessaire (moindre privilège).
- Voir `rules/remediation/k8s-secret-exposure.md` pour les patterns de référence externe de secrets.

## Exemple avant/après
Voir `examples/k8s/k8s-secret-exposure/`.

## Références
- CIS Kubernetes Benchmark: section 1.2 "API Server" — recommandation sur `--encryption-provider-config` et section 5 sur la gestion des Secrets
- CWE-312: Cleartext Storage of Sensitive Information
