# Remédiation — Exposition de Secrets Kubernetes

## Principe
Ne jamais committer de `Secret` avec des valeurs réelles, préférer un gestionnaire externe (Vault, SOPS, sealed-secrets, cloud KMS), activer le chiffrement au repos d'etcd, et injecter les secrets via `secretKeyRef` plutôt qu'en clair dans `env.value`.

## Secret commité en clair dans un dépôt Git
```yaml
# Avant — vulnérable (secret.yaml versionné avec des valeurs réelles)
apiVersion: v1
kind: Secret
metadata:
  name: db-credentials
type: Opaque
data:
  username: YWRtaW4=
  password: UyN1cGVyU2VjcmV0UGFzc3dvcmQh

# Après — sécurisé : sealed-secrets, chiffré, committable en toute sécurité
apiVersion: bitnami.com/v1alpha1
kind: SealedSecret
metadata:
  name: db-credentials
  namespace: production
spec:
  encryptedData:
    username: AgBy8hF...tronqué...
    password: AgCk2mQ...tronqué...
  template:
    metadata:
      name: db-credentials
    type: Opaque
```

## Variables d'environnement en clair
```yaml
# Avant — vulnérable
spec:
  containers:
    - name: app
      env:
        - name: DB_PASSWORD
          value: "S3cretPassword123"

# Après — sécurisé
spec:
  containers:
    - name: app
      env:
        - name: DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: db-credentials
              key: password
```

## ConfigMap utilisé à tort pour des données sensibles
```yaml
# Avant — vulnérable
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  api_key: "sk_live_EXAMPLE_NOT_A_REAL_KEY"

# Après — sécurisé : déplacer vers un Secret
apiVersion: v1
kind: Secret
metadata:
  name: app-secrets
type: Opaque
stringData:
  api_key: "sk_live_EXAMPLE_NOT_A_REAL_KEY"
```

## Chiffrement au repos d'etcd
```yaml
# Avant — vulnérable : aucune EncryptionConfiguration référencée par le kube-apiserver

# Après — sécurisé
apiVersion: apiserver.config.k8s.io/v1
kind: EncryptionConfiguration
resources:
  - resources:
      - secrets
    providers:
      - kms:
          name: myKmsProvider
          endpoint: unix:///var/run/kmsplugin/socket.sock
      - identity: {}
```
```yaml
# Flag correspondant sur le kube-apiserver
command:
  - kube-apiserver
  - --encryption-provider-config=/etc/kubernetes/enc/enc.yaml
```

## RBAC trop permissif sur la ressource `secrets`
```yaml
# Avant — vulnérable
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: broad-secret-reader
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "list", "watch"]
# ...lié via ClusterRoleBinding à un large groupe de ServiceAccounts

# Après — sécurisé : Role namespacé, scoped à un secret nommé
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: app-secret-reader
  namespace: production
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    resourceNames: ["db-credentials"]
    verbs: ["get"]
```

## Checklist de vérification post-patch
- [ ] Aucun manifest `Secret`/`SealedSecret` avec des valeurs réelles en clair n'est présent dans l'historique Git (rotation obligatoire si déjà commité).
- [ ] `git log -p -- **/secret*.yaml` ne montre plus de valeurs base64 réelles ajoutées après le correctif.
- [ ] Le kube-apiserver référence bien `--encryption-provider-config` et `kubectl get secrets -A -o json` déchiffré confirme le chiffrement au repos.
- [ ] Aucun `ConfigMap` ne contient de clé nommée `password`, `token`, `api_key` ou `secret`.
- [ ] Les manifests de Pod/Deployment utilisent `secretKeyRef`/`envFrom` et non `env.value` pour les données sensibles.
- [ ] `kubectl auth can-i list secrets --as=system:serviceaccount:<ns>:<sa>` confirme un accès restreint au strict nécessaire.
