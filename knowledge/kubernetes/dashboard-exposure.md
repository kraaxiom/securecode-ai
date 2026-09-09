---
id: dashboard-exposure
category: kubernetes
cwe: CWE-306
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Exposition du Kubernetes Dashboard

## Description
Le Kubernetes Dashboard est une interface web d'administration du cluster : mal configuré, il devient une porte d'entrée directe vers l'ensemble des ressources (pods, secrets, workloads). Ce risque est de niveau infrastructure (manifests de déploiement du Dashboard, configuration de service/ingress) et non lié à un langage applicatif. Le problème typique combine une exposition réseau non maîtrisée (Service de type `LoadBalancer`/`NodePort`, Ingress public) avec un compte de service disposant de privilèges élevés et un login désactivé ou contournable (skip login, token admin par défaut).

## Où ça apparaît typiquement
- Manifests de déploiement du Dashboard (`kubernetes-dashboard.yaml`, Helm chart officiel ou fork).
- Définition du `Service` exposant le Dashboard (`type: NodePort` ou `type: LoadBalancer`).
- Ressources `Ingress` routant vers le Dashboard sans authentification en amont (pas d'OAuth2 proxy, pas de mTLS).
- `ServiceAccount`/`ClusterRoleBinding` associant le Dashboard à `cluster-admin`.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Service Dashboard avec `type: NodePort` ou `type: LoadBalancer` au lieu de `ClusterIP`.
- Flag de déploiement `--enable-skip-login=true` ou `--enable-insecure-login=true`.
- `ClusterRoleBinding` liant le ServiceAccount du Dashboard au rôle `cluster-admin`.
- Absence de ressource `NetworkPolicy` restreignant l'accès au namespace du Dashboard.
- Ingress sans annotation d'authentification (basic-auth, oauth2-proxy) devant le Dashboard.

## Remédiation
- Exposer le Dashboard uniquement en interne (`ClusterIP` + accès via `kubectl proxy` ou VPN), jamais en accès public direct.
- Désactiver le skip-login et exiger une authentification forte (OIDC, token à courte durée de vie).
- Appliquer le principe du moindre privilège au ServiceAccount du Dashboard (rôle scoped, jamais `cluster-admin`).
- Voir `rules/remediation/dashboard-exposure.md` pour les manifests corrigés.

## Exemple avant/après
Voir `examples/k8s/dashboard-exposure/`.

## Références
- CIS Kubernetes Benchmark: recommandation sur la restriction d'accès au Kubernetes Dashboard
- CWE-306: Missing Authentication for Critical Function
