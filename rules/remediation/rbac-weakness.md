# Remédiation — Faiblesse RBAC (Role/ClusterRole trop permissifs)

## Principe
Remplacer les règles RBAC en wildcard par des règles scopées à des ressources, verbes et `apiGroups` précis. Ne jamais lier `cluster-admin` à un ServiceAccount applicatif. Auditer régulièrement les bindings.

## ClusterRole avec wildcard
```yaml
# Avant — vulnérable
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: app-role
rules:
  - apiGroups: ["*"]
    resources: ["*"]
    verbs: ["*"]

# Après — sécurisé : scoped à ce dont l'application a réellement besoin
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: app-role
  namespace: production
rules:
  - apiGroups: [""]
    resources: ["pods", "services"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["apps"]
    resources: ["deployments"]
    verbs: ["get", "list"]
```

## Binding cluster-admin à un ServiceAccount applicatif
```yaml
# Avant — vulnérable
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: app-binding
subjects:
  - kind: ServiceAccount
    name: app-sa
    namespace: production
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io

# Après — sécurisé : RoleBinding namespacé vers un rôle dédié
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: app-binding
  namespace: production
subjects:
  - kind: ServiceAccount
    name: app-sa
    namespace: production
roleRef:
  kind: Role
  name: app-role
  apiGroup: rbac.authorization.k8s.io
```

## Accès large sur les ressources sensibles
```yaml
# Avant — vulnérable
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: ops-role
rules:
  - apiGroups: [""]
    resources: ["secrets", "pods/exec"]
    verbs: ["get", "list", "create"]
  - apiGroups: ["rbac.authorization.k8s.io"]
    resources: ["clusterroles", "clusterrolebindings"]
    verbs: ["get", "list", "create", "update", "delete"]

# Après — sécurisé : séparer les responsabilités, restreindre par namespace
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: ops-role-scoped
  namespace: production
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    resourceNames: ["app-config-secret"]
    verbs: ["get"]
```

## Binding sur le ServiceAccount `default`
```yaml
# Avant — vulnérable
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: default-elevated
  namespace: production
subjects:
  - kind: ServiceAccount
    name: default
    namespace: production
roleRef:
  kind: ClusterRole
  name: edit
  apiGroup: rbac.authorization.k8s.io

# Après — sécurisé : créer un ServiceAccount dédié par application, ne jamais
# attacher de droits supplémentaires au ServiceAccount "default" du namespace.
apiVersion: v1
kind: ServiceAccount
metadata:
  name: app-sa
  namespace: production
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: app-binding
  namespace: production
subjects:
  - kind: ServiceAccount
    name: app-sa
    namespace: production
roleRef:
  kind: Role
  name: app-role
  apiGroup: rbac.authorization.k8s.io
```

## Checklist de vérification post-patch
- [ ] Aucune règle RBAC du manifest corrigé ne contient `verbs: ["*"]`, `resources: ["*"]` ou `apiGroups: ["*"]` sans justification documentée.
- [ ] Aucun `ClusterRoleBinding` ne lie un ServiceAccount applicatif à `cluster-admin` : `kubectl get clusterrolebindings -o json | jq '.items[] | select(.roleRef.name=="cluster-admin")'` ne liste que des identités d'administration légitimes.
- [ ] `kubectl auth can-i --list --as=system:serviceaccount:<ns>:<sa>` confirme un périmètre de droits limité au strict nécessaire.
- [ ] Le ServiceAccount `default` de chaque namespace ne porte aucun binding supplémentaire.
- [ ] Aucun rôle n'accorde le verbe `impersonate` ou `escalate` sans justification explicite.
- [ ] Un outil d'audit (`kubectl-who-can`, `rbac-lookup`) est exécuté après correctif et ne remonte plus l'accès excessif initial.
