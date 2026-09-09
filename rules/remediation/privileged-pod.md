# Remédiation — Pod privilégié / échappement de conteneur

## Principe
Bannir `privileged: true`, désactiver `hostPID`/`hostNetwork`/`hostIPC` sauf besoin technique strictement justifié, forcer `runAsNonRoot`, `allowPrivilegeEscalation: false` et `capabilities.drop: ["ALL"]`, et proscrire le montage de volumes `hostPath` sensibles.

## Conteneur privilégié
```yaml
# Avant — vulnérable
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app
spec:
  template:
    spec:
      containers:
        - name: app
          image: myapp:1.0
          securityContext:
            privileged: true

# Après — sécurisé
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app
spec:
  template:
    spec:
      containers:
        - name: app
          image: myapp:1.0
          securityContext:
            privileged: false
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            runAsUser: 10001
            readOnlyRootFilesystem: true
            capabilities:
              drop: ["ALL"]
```

## hostPID / hostNetwork / hostIPC au niveau pod
```yaml
# Avant — vulnérable
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app
spec:
  template:
    spec:
      hostPID: true
      hostNetwork: true
      hostIPC: true
      containers:
        - name: app
          image: myapp:1.0

# Après — sécurisé
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app
spec:
  template:
    spec:
      hostPID: false
      hostNetwork: false
      hostIPC: false
      containers:
        - name: app
          image: myapp:1.0
```

## Capacités Linux étendues
```yaml
# Avant — vulnérable
securityContext:
  capabilities:
    add: ["SYS_ADMIN", "NET_ADMIN", "ALL"]

# Après — sécurisé : ne garder que le strict nécessaire, jamais SYS_ADMIN/ALL
securityContext:
  capabilities:
    drop: ["ALL"]
    add: ["NET_BIND_SERVICE"]
```

## Volume hostPath sensible
```yaml
# Avant — vulnérable
volumes:
  - name: docker-sock
    hostPath:
      path: /var/run/docker.sock
      type: Socket

# Après — sécurisé : supprimer ce volume ; si un accès au runtime est requis,
# utiliser un mécanisme dédié (ex: CRI socket exposé via un DaemonSet restreint
# et un admission controller), jamais un montage direct depuis un pod applicatif.
```

## Application via Pod Security Standards
```yaml
# Après — sécurisé : forcer le profil "restricted" au niveau namespace
apiVersion: v1
kind: Namespace
metadata:
  name: production
  labels:
    pod-security.kubernetes.io/enforce: restricted
    pod-security.kubernetes.io/audit: restricted
    pod-security.kubernetes.io/warn: restricted
```

## Checklist de vérification post-patch
- [ ] Aucun conteneur du manifest corrigé n'a `securityContext.privileged: true`.
- [ ] `hostPID`, `hostNetwork` et `hostIPC` sont absents ou explicitement à `false`.
- [ ] `allowPrivilegeEscalation: false`, `runAsNonRoot: true` et `capabilities.drop: ["ALL"]` sont présents sur chaque conteneur.
- [ ] Aucune capacité `SYS_ADMIN`, `NET_ADMIN` ou `ALL` n'est ajoutée sans justification documentée.
- [ ] Aucun volume `hostPath` ne pointe vers `/`, `/var/run/docker.sock`, `/proc` ou `/etc`.
- [ ] Le namespace applique le label `pod-security.kubernetes.io/enforce: restricted` (ou équivalent OPA/Gatekeeper/Kyverno) et un pod non conforme est bien rejeté à la création (`kubectl apply` test).
