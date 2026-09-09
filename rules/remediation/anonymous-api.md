# Remédiation — Accès anonyme au serveur API Kubernetes

## Principe
Désactiver l'authentification anonyme sur le kube-apiserver, s'assurer qu'aucun binding RBAC n'accorde de droits au groupe `system:anonymous`/`system:unauthenticated`, et restreindre l'accès réseau à l'API server.

## Configuration du kube-apiserver
```yaml
# Avant — vulnérable (/etc/kubernetes/manifests/kube-apiserver.yaml)
apiVersion: v1
kind: Pod
metadata:
  name: kube-apiserver
  namespace: kube-system
spec:
  containers:
    - name: kube-apiserver
      command:
        - kube-apiserver
        - --anonymous-auth=true
        - --insecure-bind-address=0.0.0.0
        - --authorization-mode=AlwaysAllow

# Après — sécurisé
apiVersion: v1
kind: Pod
metadata:
  name: kube-apiserver
  namespace: kube-system
spec:
  containers:
    - name: kube-apiserver
      command:
        - kube-apiserver
        - --anonymous-auth=false
        - --authorization-mode=Node,RBAC
        - --bind-address=127.0.0.1
```

## ClusterRoleBinding exposant le groupe anonyme
```yaml
# Avant — vulnérable
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: anonymous-view
subjects:
  - kind: Group
    name: system:unauthenticated
    apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: view
  apiGroup: rbac.authorization.k8s.io

# Après — sécurisé : supprimer purement et simplement ce binding.
# Aucun binding RBAC ne doit référencer system:anonymous ou system:unauthenticated.
```

## Restriction réseau (managed clusters EKS/GKE/AKS)
```yaml
# Avant — vulnérable : endpoint public de l'API server non restreint
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig
vpc:
  clusterEndpoints:
    publicAccess: true
    privateAccess: false

# Après — sécurisé : accès public restreint par allowlist ou désactivé
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig
vpc:
  clusterEndpoints:
    publicAccess: true
    privateAccess: true
  publicAccessCIDRs:
    - "203.0.113.10/32"
```

## Checklist de vérification post-patch
- [ ] `--anonymous-auth=false` confirmé dans la configuration effective du kube-apiserver.
- [ ] `--authorization-mode` inclut `RBAC` (jamais `AlwaysAllow`).
- [ ] Aucun `ClusterRoleBinding`/`RoleBinding` ne référence `system:anonymous` ou `system:unauthenticated` : `kubectl get clusterrolebindings -o json | jq '.items[] | select(.subjects[]?.name=="system:anonymous" or .subjects[]?.name=="system:unauthenticated")'` retourne vide.
- [ ] `kubectl auth can-i --list --as=system:anonymous` ne montre aucun droit.
- [ ] L'endpoint API server n'est pas accessible depuis `0.0.0.0/0` sans allowlist IP documentée.
- [ ] Une tentative d'appel API sans jeton d'authentification retourne `401 Unauthorized`.
