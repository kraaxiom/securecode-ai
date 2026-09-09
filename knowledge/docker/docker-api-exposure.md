---
id: docker-api-exposure
category: docker
cwe: CWE-306
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Exposition de l'API Docker Engine (daemon TCP non authentifié)

## Description
Le daemon Docker peut être configuré pour écouter sur un socket TCP (`-H tcp://0.0.0.0:2375`) au lieu du socket Unix local. Sans authentification mutuelle TLS, quiconque peut atteindre ce port dispose d'un accès équivalent à root sur l'hôte, car l'API Docker permet de monter n'importe quel volume du système de fichiers hôte dans un conteneur.

## Où ça apparaît typiquement
- Configuration du daemon (`daemon.json`, unit systemd) exposant le port 2375 (non chiffré) plutôt que 2376 (TLS).
- Environnements CI/CD ou cloud où le socket Docker est exposé pour faciliter l'orchestration distante.
- Règles de pare-feu/groupe de sécurité cloud n'ayant pas restreint l'accès au port de l'API Docker.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Option `-H tcp://0.0.0.0:2375` ou équivalent dans la configuration du daemon Docker.
- Port 2375/2376 ouvert vers l'extérieur dans une configuration réseau/cloud (security group, firewall).
- Absence de configuration `tlsverify`, `tlscacert`, `tlscert`, `tlskey` sur un daemon exposé en TCP.

## Remédiation
- Ne jamais exposer l'API Docker sans authentification ; utiliser le socket Unix local par défaut.
- Si un accès distant est requis, activer TLS mutuel obligatoire (`tlsverify=true`) avec certificats clients.
- Restreindre l'accès réseau au port de l'API via pare-feu/groupe de sécurité, idéalement uniquement depuis un réseau privé.
- Voir `rules/remediation/docker-api-exposure.md` pour les diffs de configuration.

## Exemple avant/après
Voir `examples/config/docker-api-exposure/` (à créer selon le même schéma).

## Références
- Docker documentation: Protect the Docker daemon socket
- CIS Docker Benchmark
- CWE-306: Missing Authentication for Critical Function
