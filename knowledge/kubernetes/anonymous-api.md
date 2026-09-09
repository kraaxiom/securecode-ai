---
id: anonymous-api
category: kubernetes
cwe: CWE-306
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Accès anonyme au serveur API Kubernetes

## Description
L'API server de Kubernetes est le point d'entrée central du cluster : toute personne pouvant l'atteindre sans authentification peut potentiellement lister, créer ou modifier des ressources. Ce constat est de niveau infrastructure (manifests K8s, configuration du cluster) et non lié à un langage de programmation. Quand le flag `--anonymous-auth` reste activé et qu'aucune politique RBAC restrictive n'est associée au groupe `system:anonymous`, un attaquant réseau peut interagir avec l'API sans jamais présenter d'identifiant. C'est souvent combiné à une exposition réseau non maîtrisée (API server accessible depuis Internet) pour former une chaîne de compromission complète.

## Où ça apparaît typiquement
- Configuration du kube-apiserver (flags de démarrage, manifests statiques `/etc/kubernetes/manifests/kube-apiserver.yaml`).
- ClusterRoleBinding liant le groupe `system:anonymous` ou `system:unauthenticated` à un rôle (même `view`).
- Configurations managées (EKS/GKE/AKS) où l'endpoint public de l'API n'est pas restreint par IP allowlist.
- Fichiers `kubeadm-config` ou Helm charts de bootstrap de cluster.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Flag `--anonymous-auth=true` explicite dans la configuration du kube-apiserver (ou absence du flag alors que la valeur par défaut historique est `true` selon la version).
- `ClusterRoleBinding` ou `RoleBinding` référençant `system:anonymous` ou le groupe `system:unauthenticated` dans `subjects`.
- Endpoint API server configuré en accès public (`0.0.0.0/0`) sans authentification forte documentée à côté.
- Absence de `--authorization-mode=RBAC` (mode `AlwaysAllow` ou absent).

## Remédiation
- Désactiver l'authentification anonyme (`--anonymous-auth=false`) sauf besoin explicite et documenté (ex: endpoints de santé publics gérés séparément).
- S'assurer qu'aucun binding RBAC n'accorde de droits au groupe `system:anonymous`/`system:unauthenticated`.
- Restreindre l'accès réseau à l'API server (allowlist IP, réseau privé, bastion) en complément de l'authentification.
- Voir `rules/remediation/anonymous-api.md` pour les correctifs de configuration détaillés.

## Exemple avant/après
Voir `examples/k8s/anonymous-api/`.

## Références
- CIS Kubernetes Benchmark: section 1.2 "API Server" — contrôle sur `--anonymous-auth`
- CWE-306: Missing Authentication for Critical Function
