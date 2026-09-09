# Remédiation — Exposition du Kubernetes Dashboard

## Principe
Exposer le Dashboard uniquement en interne (`ClusterIP`), désactiver le skip-login/login non sécurisé, et scoper strictement le `ServiceAccount` associé plutôt que de le lier à `cluster-admin`.

## Service exposant le Dashboard
```yaml
# Avant — vulnérable
apiVersion: v1
kind: Service
metadata:
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
spec:
  type: LoadBalancer
  ports:
    - port: 443
      targetPort: 8443
  selector:
    k8s-app: kubernetes-dashboard

# Après — sécurisé : accès uniquement via kubectl proxy ou VPN interne
apiVersion: v1
kind: Service
metadata:
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
spec:
  type: ClusterIP
  ports:
    - port: 443
      targetPort: 8443
  selector:
    k8s-app: kubernetes-dashboard
```

## Flags de déploiement du Dashboard
```yaml
# Avant — vulnérable
containers:
  - name: kubernetes-dashboard
    args:
      - --enable-skip-login=true
      - --enable-insecure-login=true
      - --disable-csrf-protection

# Après — sécurisé
containers:
  - name: kubernetes-dashboard
    args:
      - --auto-generate-certificates
      - --namespace=kubernetes-dashboard
```

## ClusterRoleBinding du ServiceAccount du Dashboard
```yaml
# Avant — vulnérable
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: kubernetes-dashboard
subjects:
  - kind: ServiceAccount
    name: kubernetes-dashboard
    namespace: kubernetes-dashboard
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io

# Après — sécurisé : rôle scoped, jamais cluster-admin
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: kubernetes-dashboard-minimal
  namespace: kubernetes-dashboard
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    resourceNames: ["kubernetes-dashboard-key-holder", "kubernetes-dashboard-certs"]
    verbs: ["get", "update", "delete"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: kubernetes-dashboard-minimal
  namespace: kubernetes-dashboard
subjects:
  - kind: ServiceAccount
    name: kubernetes-dashboard
    namespace: kubernetes-dashboard
roleRef:
  kind: Role
  name: kubernetes-dashboard-minimal
  apiGroup: rbac.authorization.k8s.io
```

## Ingress sans authentification en amont
```yaml
# Avant — vulnérable : aucune authentification devant le Dashboard
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: dashboard-ingress
  namespace: kubernetes-dashboard
spec:
  rules:
    - host: dashboard.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: kubernetes-dashboard
                port:
                  number: 443

# Après — sécurisé : oauth2-proxy en amont + restriction réseau
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: dashboard-ingress
  namespace: kubernetes-dashboard
  annotations:
    nginx.ingress.kubernetes.io/auth-url: "https://oauth2-proxy.example.com/oauth2/auth"
    nginx.ingress.kubernetes.io/auth-signin: "https://oauth2-proxy.example.com/oauth2/start"
    nginx.ingress.kubernetes.io/whitelist-source-range: "10.0.0.0/8"
spec:
  rules:
    - host: dashboard.internal.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: kubernetes-dashboard
                port:
                  number: 443
```

## Checklist de vérification post-patch
- [ ] Le Service du Dashboard est de type `ClusterIP` (plus de `NodePort`/`LoadBalancer` public).
- [ ] `--enable-skip-login` et `--enable-insecure-login` sont absents ou à `false`.
- [ ] Le `ServiceAccount` du Dashboard n'est plus lié à `cluster-admin` : `kubectl get clusterrolebinding -o json | jq '.items[] | select(.subjects[]?.name=="kubernetes-dashboard")'` retourne un rôle scoped.
- [ ] `kubectl auth can-i --list --as=system:serviceaccount:kubernetes-dashboard:kubernetes-dashboard` ne montre que les droits strictement nécessaires.
- [ ] Une ressource `NetworkPolicy` limite le trafic entrant vers le namespace du Dashboard.
- [ ] L'accès externe au Dashboard nécessite une authentification forte (OIDC, oauth2-proxy) validée par un test manuel.
