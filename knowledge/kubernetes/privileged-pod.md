---
id: privileged-pod
category: kubernetes
cwe: CWE-250
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Pod privilégié / échappement de conteneur

## Description
Un pod configuré avec des privilèges excessifs (`privileged: true`, `hostPID`, `hostNetwork`, `hostIPC`, capacités Linux étendues) rompt l'isolation normale entre le conteneur et le nœud hôte. Ce risque est de niveau infrastructure (manifests de Pod/Deployment/DaemonSet, `securityContext`) et ne dépend d'aucun langage applicatif. Un pod ainsi configuré peut accéder aux périphériques de l'hôte, voir les processus d'autres pods, ou monter le système de fichiers du nœud, ce qui facilite une élévation de privilèges jusqu'au nœud lui-même.

## Où ça apparaît typiquement
- `securityContext` au niveau pod ou conteneur dans les manifests Deployment/DaemonSet/StatefulSet.
- Charts Helm exposant des valeurs comme `privileged`, `hostNetwork`, `hostPID` sans garde-fou.
- DaemonSets d'agents systèmes (monitoring, CNI, logging) qui légitimement nécessitent parfois des privilèges élevés mais sans scope minimal.
- Montage de volumes hostPath sensibles (`/`, `/var/run/docker.sock`, `/etc`).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- `securityContext.privileged: true` dans la spec d'un conteneur.
- `hostPID: true`, `hostNetwork: true` ou `hostIPC: true` au niveau du pod.
- `securityContext.capabilities.add` incluant des capacités larges (`SYS_ADMIN`, `NET_ADMIN`, `ALL`).
- `allowPrivilegeEscalation` absent ou explicitement à `true`.
- Volume `hostPath` pointant vers des chemins sensibles du nœud (`/`, `/var/run/docker.sock`, `/proc`).
- Absence de `runAsNonRoot: true` combinée à une image tournant en root par défaut.

## Remédiation
- Bannir `privileged: true` sauf besoin technique strictement justifié et isolé (nœuds dédiés, admission controller).
- Définir `allowPrivilegeEscalation: false`, `runAsNonRoot: true`, et un `capabilities.drop: ["ALL"]` par défaut, n'ajoutant que le strict nécessaire.
- Utiliser des Pod Security Standards (`restricted` profile) ou un admission controller (OPA/Gatekeeper, Kyverno) pour bloquer ces configurations au niveau cluster.
- Voir `rules/remediation/privileged-pod.md` pour les manifests `securityContext` corrigés.

## Exemple avant/après
Voir `examples/k8s/privileged-pod/`.

## Références
- CIS Kubernetes Benchmark: section 5.2 "Pod Security Standards" — contrôles sur conteneurs privilégiés et hostPID/hostNetwork/hostIPC
- CWE-250: Execution with Unnecessary Privileges
